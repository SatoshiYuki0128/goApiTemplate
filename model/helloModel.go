package model

type HelloReq struct {
	Number int `json:"number"`
}

type HelloRes struct {
	Str string `json:"str"`
}
