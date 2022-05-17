package testhelpers

import (
	"bytes"
	"fmt"
	"strings"
)

type LogField struct {
	Type  LogType
	Value string
}

type AppLogInspector struct {
	Buffer *bytes.Buffer
}

func (ins AppLogInspector) HasError() bool {
	return ins.Has([]LogField{{Type: LogTypeLevel, Value: "error"}})
}

func (ins AppLogInspector) Has(fields []LogField) bool {
	logs := strings.Split(ins.Buffer.String(), "\n")

	for _, log := range logs {
		for _, field := range fields {
			if !strings.Contains(log, fmt.Sprintf("%s=%s", field.Type, field.Value)) {
				break
			}

			return true
		}
	}

	return false
}
