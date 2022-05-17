package testhelpers

import "github.com/djokcik/gophkeeper/e2e/testmodels"

type LogType string

var (
	LogTypeLevel LogType = "level"
)

type RecordType string

var (
	PersonalDataRecord RecordType = "PersonalDataRecord"
	TextDataRecord     RecordType = "TextDataRecord"
	BinaryDataRecord   RecordType = "BinaryDataRecord"
	BankCardRecord     RecordType = "BankCardRecord"
)

var E2EUser = testmodels.ClientUser{Username: "my_name", Password: "my_password"}

const (
	CallSaveRecordPersonalDataHandler        = "RpcHandler.SaveRecordPersonalDataHandler"
	CallLoadRecordPersonalDataByKeyHandler   = "RpcHandler.LoadRecordPersonalDataByKeyHandler"
	CallRemoveRecordPersonalDataByKeyHandler = "RpcHandler.RemoveRecordPersonalDataByKeyHandler"

	CallSaveRecordBankCardHandler        = "RpcHandler.SaveRecordBankCardHandler"
	CallLoadRecordBankCardByKeyHandler   = "RpcHandler.LoadRecordBankCardByKeyHandler"
	CallRemoveRecordBankCardByKeyHandler = "RpcHandler.RemoveRecordBankCardByKeyHandler"

	CallSaveRecordTextDataHandler        = "RpcHandler.SaveRecordTextDataHandler"
	CallLoadRecordTextDataByKeyHandler   = "RpcHandler.LoadRecordTextDataByKeyHandler"
	CallRemoveRecordTextDataByKeyHandler = "RpcHandler.RemoveRecordTextDataByKeyHandler"

	CallSaveRecordBinaryDataHandler        = "RpcHandler.SaveRecordBinaryDataHandler"
	CallLoadRecordBinaryDataByKeyHandler   = "RpcHandler.LoadRecordBinaryDataByKeyHandler"
	CallRemoveRecordBinaryDataByKeyHandler = "RpcHandler.RemoveRecordBinaryDataByKeyHandler"

	CallRegisterHandler = "RpcHandler.RegisterHandler"
	CallSignInHandler   = "RpcHandler.SignInHandler"
)
