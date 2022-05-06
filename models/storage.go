package models

type StorageData struct {
	User User `json:"user"`

	PersonalData map[string]string `json:"personal_data"`
	BinaryData   map[string]string `json:"binary_data"`
	TextData     map[string]string `json:"text_data"`
	BankCardData map[string]string `json:"bankCard_data"`
}
