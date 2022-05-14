package clientmodels

type (
	RecordPersonalData struct {
		Username string
		Password string
		URL      string
		Comment  string
	}

	RecordBankCardData struct {
		CardNumber string
		Year       string
		CVV        string
		Comment    string
	}

	RecordTextData struct {
		Text    string
		Comment string
	}

	RecordBinaryData struct {
		Data    []byte
		Comment string
	}
)
