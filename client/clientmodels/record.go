// Package clientmodels provides client models
package clientmodels

type (
	// RecordPersonalData struct for type record is personal data
	RecordPersonalData struct {
		Username string
		Password string
		URL      string
		Comment  string
	}

	// RecordBankCardData struct for type record is bank card data
	RecordBankCardData struct {
		CardNumber string
		Year       string
		CVV        string
		Comment    string
	}

	// RecordTextData struct for type record is text data
	RecordTextData struct {
		Text    string
		Comment string
	}

	// RecordBinaryData struct for type record is binary data
	RecordBinaryData struct {
		Data    []byte
		Comment string
	}
)
