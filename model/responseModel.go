package model

type FlowData struct {
	CtrlError
	Request  interface{}            `json:"request"`
	Response interface{}            `json:"response"`
	Data     map[string]interface{} `json:"data"`
}

type CtrlError struct {
	ControlCode string `json:"controlCode"`
	ServError
}

type ServError struct {
	ServerCode   string `json:"serverCode"`
	FunctionCode string `json:"functionCode"`
	Err          error  `json:"err"`
	Msg          string `json:"msg"`
}
