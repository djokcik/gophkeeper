package clientmodels

var (
	SaveMethod   RecordActionType = "save"
	RemoveMethod RecordActionType = "remove"
)

type (
	// StoreActions are actions with local storage
	StoreActions []RecordFileLine

	RecordActionType string

	// RecordFileLine is struct for store did offline
	RecordFileLine struct {
		Method     string           `json:"method"`
		ActionType RecordActionType `json:"action_type"`
		Data       string           `json:"data"`
		Key        string           `json:"key"`
	}
)
