package clientmodels

var (
	SaveMethod   RecordActionType = "save"
	RemoveMethod RecordActionType = "remove"
)

type (
	RecordActionType string

	RecordFileLine struct {
		Method     string           `json:"method"`
		ActionType RecordActionType `json:"action_type"`
		Data       string           `json:"data"`
		Key        string           `json:"key"`
	}
)