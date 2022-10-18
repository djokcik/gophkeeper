package testhelpers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

type Application struct {
	t *testing.T

	Cmd        *exec.Cmd
	AppLog     *AppLogInspector
	AppRequest *AppRequest

	exit chan error
}

func CreateApplication(ctx context.Context, t *testing.T) *Application {
	cmd := exec.CommandContext(ctx, "/app/bin/app")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Env = os.Environ()

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	return &Application{
		Cmd:    cmd,
		t:      t,
		AppLog: &AppLogInspector{Buffer: &out},
		exit:   make(chan error),
	}
}

func (a *Application) Run() {
	a.t.Logf("Start app. Envs: %+v", a.Cmd.Env)
	err := a.Cmd.Run()
	fmt.Println(err)

	a.exit <- err
}

func (a *Application) InitClient() {
	a.AppRequest = NewAppRequest(a.t)
}

func (a *Application) AddEnv(key string, value string) {
	a.Cmd.Env = append(a.Cmd.Env, fmt.Sprintf("%s=%s", key, value))
}

func (a *Application) ClearUsers() {
	err := os.RemoveAll(fmt.Sprintf("%s", E2EConfig.StorePath))
	require.Equal(a.t, err, nil)
	err = os.MkdirAll(E2EConfig.StorePath, os.ModeDir)
	require.Equal(a.t, err, nil)
}

func (a *Application) Close() {
	a.Cmd.Process.Signal(syscall.SIGINT)
	err := <-a.exit

	a.t.Cleanup(func() {
		a.t.Logf("Получен STDOUT лог процесса:\n\n%s", a.AppLog.Buffer.String())
		require.Equal(a.t, err, nil, fmt.Sprintf("Запуск приложения завершился с ошибкой: %v", err))
		require.Equal(
			a.t,
			a.AppLog.HasError(),
			false,
			"В логах приложения есть ошибки. Найден level=error.",
		)
	})
}
