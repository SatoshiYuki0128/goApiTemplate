package controller

import (
	"goApiFramework/model"
	"goApiFramework/service/common"
	"goApiFramework/service/hello"
	"net/http"
)

// Hello controllerCode:H1
func Hello(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if common.GetAndValidateRequest[model.HelloReq](r, &flowData, "H1") {
		hello.Hello(&flowData, "H1")
	}
}
