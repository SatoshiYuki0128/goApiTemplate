package common

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goApiFramework/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func initFlowData(flowData *model.FlowData, controllerCode, serviceCode string) {
	if flowData.Data == nil {
		flowData.Data = map[string]interface{}{}
	}
}

func getRequestBody(r *http.Request, flowData *model.FlowData, controllerCode, serviceCode string) (data []byte, isOK bool) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		flowData.CtrlError = model.CtrlError{
			ControlCode: controllerCode,
			ServError: model.ServError{
				ServerCode:   serviceCode,
				FunctionCode: "00",
				Msg:          "GetRequestBody Error",
				Err:          err}}
	}
	isOK = true
	return
}

func checkModel[T any](data []byte, r *http.Request, flowData *model.FlowData, controllerCode, serviceCode string) {

	var request T
	requestJson2Obj(data, flowData, &request, controllerCode, serviceCode)

	setRequestHeaderIntoRequest(r.Header, &request)

	setQueryStringIntoRequest(r.URL.Query(), &request)

	setUrlParamIntoRequest(mux.Vars(r), &request)

	//validateRequest(flowData, request, controllerCode, "HH2")

	flowData.Request = request

	//flowData.Data["RequestIP"] = getRequestIP(r)

	return
}

//func validateRequest(flowData *model.FlowData, request interface{}, controllerCode, serviceCode string) {
//	zhTranslator := zh.New()
//	enTranslator := en.New()
//	uni := ut.New(zhTranslator, zhTranslator, enTranslator)
//	curLocales := "zh"                             // 设置当前语言类型
//	translator, _ := uni.GetTranslator(curLocales) // 获取对应语言的转换器
//	validate := validator.New()
//	validate.RegisterValidation("m1gt0", isIdFieldMinusOneOrGreatThanZero)
//
//	switch curLocales {
//	case "zh":
//		// 内置tag注册 中文翻译器
//		_ = zhtrans.RegisterDefaultTranslations(validate, translator)
//
//		// 自定义tag注册 中文翻译器
//		_ = validate.RegisterTranslation("m1gt0", translator, func(ut ut.Translator) error {
//			if err := ut.Add("m1gt0", "{0}必須等於-1或大於0", false); err != nil {
//				return err
//			}
//
//			return nil
//		}, func(ut ut.Translator, fe validator.FieldError) string {
//			t, err := ut.T(fe.Tag(), fe.Field())
//			if err != nil {
//				fmt.Printf("警告: 翻译字段错误: %#v", fe)
//				return fe.(error).Error()
//			}
//
//			return t
//		})
//	case "en":
//		// 内置tag注册 英文翻译器
//		_ = entrans.RegisterDefaultTranslations(validate, translator)
//	}
//
//	err := validate.Struct(request)
//	if err != nil {
//		validateErrorMessage := ""
//		errs := err.(validator.ValidationErrors)
//		for _, e := range errs {
//			validateErrorMessage += e.Translate(translator) + ";"
//		}
//		flowData.CtrlError = model.CtrlError{
//			ControlCode: controllerCode,
//			ServError: model.ServError{
//				ServerCode:   serviceCode,
//				FunctionCode: "HH7",
//				Msg:          "Validate Error:" + validateErrorMessage,
//				Err:          err}}
//	}
//
//	return
//}

func setUrlParamIntoRequest(urlParams map[string]string, request interface{}) {
	for urlParam := range urlParams {
		dynamicSetModel(request, strings.ToLower(urlParam), urlParams[urlParam])
	}
}

func setQueryStringIntoRequest(querys url.Values, request interface{}) {
	for query := range querys {
		dynamicSetModel(request, strings.ToLower(query), querys.Get(query))
	}
}

func requestJson2Obj(data []byte, flowData *model.FlowData, request interface{}, controllerCode, serviceCode string) {
	err := json.Unmarshal(data, &request)
	if err != nil {
		flowData.CtrlError = model.CtrlError{
			ControlCode: controllerCode,
			ServError: model.ServError{
				ServerCode:   serviceCode,
				FunctionCode: "HH3",
				Msg:          "Json Unmarshal Error",
				Err:          err}}
	}

	return
}

func setRequestHeaderIntoRequest(headers http.Header, request interface{}) {
	for header := range headers {
		dynamicSetModel(request, strings.ToLower(header), headers.Get(header))
	}
}

func dynamicSetModel(targetModel interface{}, fieldName string, fieldValue string) {
	reflectModelValue := reflect.ValueOf(targetModel).Elem()
	reflectModelType := reflect.ValueOf(targetModel).Type().Elem()
	nestReflectSetModel(reflectModelType, reflectModelValue, fieldName, fieldValue)
}

func nestReflectSetModel(reflectType reflect.Type, reflectValue reflect.Value, fieldName string, fieldValue string) {
	for i := 0; i < reflectType.NumField(); i++ {
		rf := reflectType.Field(i)
		if rf.Type.Kind() == reflect.Struct {
			nestReflectSetModel(rf.Type, reflectValue.Field(i), fieldName, fieldValue)
			continue
		}
		v, ok := rf.Tag.Lookup("json")
		if ok && strings.Split(strings.ToLower(v), ",")[0] == fieldName {
			reflectField := reflectValue.Field(i)
			switch reflectField.Kind() {
			case reflect.Bool:
				boolData, _ := strconv.ParseBool(fieldValue)
				reflectField.SetBool(boolData)
				break
			case reflect.String:
				reflectField.SetString(fieldValue)
				break
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				intData, _ := strconv.Atoi(fieldValue)
				reflectField.SetInt(int64(intData))
				break
			case reflect.Float32, reflect.Float64:
				floatData, _ := strconv.ParseFloat(fieldValue, 64)
				reflectField.SetFloat(floatData)
				break
			default:
				reflectField.SetBytes([]byte(fieldValue))
				break
			}
		} else {
			reflectField := reflectValue.FieldByName(fieldName)
			if reflectField.IsValid() {
				switch reflectField.Kind() {
				case reflect.Bool:
					boolData, _ := strconv.ParseBool(fieldValue)
					reflectField.SetBool(boolData)
					break
				case reflect.String:
					reflectField.SetString(fieldValue)
					break
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					intData, _ := strconv.Atoi(fieldValue)
					reflectField.SetInt(int64(intData))
					break
				case reflect.Float32, reflect.Float64:
					floatData, _ := strconv.ParseFloat(fieldValue, 64)
					reflectField.SetFloat(floatData)
					break
				default:
					reflectField.SetBytes([]byte(fieldValue))
					break
				}
			}

		}
	}

}
