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
	CallSaveRecordPersonalDataHandler        = "RPCHandler.SaveRecordPersonalDataHandler"
	CallLoadRecordPersonalDataByKeyHandler   = "RPCHandler.LoadRecordPersonalDataByKeyHandler"
	CallRemoveRecordPersonalDataByKeyHandler = "RPCHandler.RemoveRecordPersonalDataByKeyHandler"

	CallSaveRecordBankCardHandler        = "RPCHandler.SaveRecordBankCardHandler"
	CallLoadRecordBankCardByKeyHandler   = "RPCHandler.LoadRecordBankCardByKeyHandler"
	CallRemoveRecordBankCardByKeyHandler = "RPCHandler.RemoveRecordBankCardByKeyHandler"

	CallSaveRecordTextDataHandler        = "RPCHandler.SaveRecordTextDataHandler"
	CallLoadRecordTextDataByKeyHandler   = "RPCHandler.LoadRecordTextDataByKeyHandler"
	CallRemoveRecordTextDataByKeyHandler = "RPCHandler.RemoveRecordTextDataByKeyHandler"

	CallSaveRecordBinaryDataHandler        = "RPCHandler.SaveRecordBinaryDataHandler"
	CallLoadRecordBinaryDataByKeyHandler   = "RPCHandler.LoadRecordBinaryDataByKeyHandler"
	CallRemoveRecordBinaryDataByKeyHandler = "RPCHandler.RemoveRecordBinaryDataByKeyHandler"

	CallRegisterHandler = "RPCHandler.RegisterHandler"
	CallSignInHandler   = "RPCHandler.SignInHandler"
)
