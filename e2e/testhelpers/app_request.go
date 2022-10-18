package testhelpers

import (
	"crypto/tls"
	"github.com/djokcik/gophkeeper/e2e/testmodels"
	"github.com/stretchr/testify/require"
	"net/rpc"
	"testing"
)

type AppRequest struct {
	t   *testing.T
	api *rpc.Client
}

func NewAppRequest(t *testing.T) *AppRequest {
	conf := LoadClientCertificate()

	conn, err := tls.Dial("tcp", E2EConfig.Address, conf)
	if err != nil {
		t.Fatalf("error dial tls: %s. Address: %s", err.Error(), E2EConfig.Address)
	}

	return &AppRequest{
		t:   t,
		api: rpc.NewClient(conn),
	}
}

func (r AppRequest) Register() string {
	registerDto := testmodels.RegisterDto{Login: E2EUser.Username, Password: E2EUser.Password}

	var token string
	err := r.api.Call(CallRegisterHandler, registerDto, &token)

	require.Equal(r.t, err, nil, "Не удалось зарегистрировать пользователя")

	return token
}

func (r AppRequest) Login() string {
	loginDto := testmodels.RegisterDto{Login: E2EUser.Username, Password: E2EUser.Password}

	var token string
	err := r.api.Call(CallSignInHandler, loginDto, &token)

	require.Equal(r.t, err, nil, "Не удалось авторизоваться")

	return token
}

func (r AppRequest) SaveData(token, key string, data string, recordType RecordType) {
	var err error

	switch recordType {
	case PersonalDataRecord:
		var reply struct{}
		err = r.api.Call(
			CallSaveRecordPersonalDataHandler,
			testmodels.SaveRecordRequestDto{Token: token, Key: key, Data: data},
			&reply,
		)
	case TextDataRecord:
		var reply struct{}
		err = r.api.Call(
			CallSaveRecordTextDataHandler,
			testmodels.SaveRecordRequestDto{Token: token, Key: key, Data: data},
			&reply,
		)
	case BinaryDataRecord:
		var reply struct{}
		err = r.api.Call(
			CallSaveRecordBinaryDataHandler,
			testmodels.SaveRecordRequestDto{Token: token, Key: key, Data: data},
			&reply,
		)
	case BankCardRecord:
		var reply struct{}
		err = r.api.Call(
			CallSaveRecordBankCardHandler,
			testmodels.SaveRecordRequestDto{Token: token, Key: key, Data: data},
			&reply,
		)
	default:
		require.Fail(r.t, "Передан неизвестный recordType")
	}

	require.Equal(r.t, err, nil)
}

func (r AppRequest) LoadData(token, key string, recordType RecordType) string {
	var err error
	var reply string

	switch recordType {
	case PersonalDataRecord:
		err = r.api.Call(
			CallLoadRecordPersonalDataByKeyHandler,
			testmodels.LoadRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	case TextDataRecord:
		err = r.api.Call(
			CallLoadRecordTextDataByKeyHandler,
			testmodels.LoadRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	case BinaryDataRecord:
		err = r.api.Call(
			CallLoadRecordBinaryDataByKeyHandler,
			testmodels.LoadRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	case BankCardRecord:
		err = r.api.Call(
			CallLoadRecordBankCardByKeyHandler,
			testmodels.LoadRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	default:
		require.FailNow(r.t, "Передан неизвестный recordType")
	}

	if err == nil {
		return reply
	}

	require.Equal(r.t, err, nil)

	return ""
}

func (r AppRequest) RemoveData(token, key string, recordType RecordType) {
	var err error
	var reply struct{}

	switch recordType {
	case PersonalDataRecord:
		err = r.api.Call(
			CallRemoveRecordPersonalDataByKeyHandler,
			testmodels.RemoveRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	case TextDataRecord:
		err = r.api.Call(
			CallRemoveRecordTextDataByKeyHandler,
			testmodels.RemoveRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	case BinaryDataRecord:
		err = r.api.Call(
			CallRemoveRecordBinaryDataByKeyHandler,
			testmodels.RemoveRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	case BankCardRecord:
		err = r.api.Call(
			CallRemoveRecordBankCardByKeyHandler,
			testmodels.RemoveRecordRequestDto{Token: token, Key: key},
			&reply,
		)
	default:
		require.FailNow(r.t, "Передан неизвестный recordType")
	}

	if err != nil {
		require.Equal(r.t, err, nil)
	}
}
