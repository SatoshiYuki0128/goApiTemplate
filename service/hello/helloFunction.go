package helloService

import (
	"goApiFramework/model"
	"math"
	"strconv"
)

func printHello(flowData *model.FlowData, controllerCode string) {
	n := flowData.Request.(model.HelloReq).Number

	var result []string
	for i := 0; i < int(math.Pow(2, 2*float64(n))); i++ {

		bin := strconv.FormatInt(int64(i), 2)

		for len(bin) < 2*n {
			bin = "0" + bin
		}

		str := ""
		for j := range bin {
			if bin[j] == '0' {
				str += "("
			} else {
				str += ")"
			}
		}

		if validate(str) {
			result = append(result, str)
		}

	}

	flowData.Response = result
}

func validate(str string) bool {
	var stack []string

	for i := range str {
		if str[i] == '(' {
			stack = append(stack, "(")
		} else {
			if len(stack) == 0 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
