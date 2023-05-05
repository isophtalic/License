package helpers

import (
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

/*
	JsonStruct is used to select specific attribute of some struct
*/
type JsonResponse struct {
	obj    interface{}
	result map[string]interface{}
}

/*
	Initiate new JsonStruct
*/
func JSON(obj interface{}) *JsonResponse {
	if reflect.TypeOf(obj).Kind() != reflect.Pointer {
		panic("JSON: input must be pointer")
	}

	if reflect.TypeOf(obj).Elem().Kind() != reflect.Struct {
		panic("JSON: value input must be struct")
	}

	return &JsonResponse{
		obj: obj,
	}
}

/*
	Get results after using Include or Exclude attributes
*/
func (j *JsonResponse) Get() map[string]interface{} {
	return j.result
}

/*
	Include attributes
*/
func (j *JsonResponse) Include(attrs ...string) *JsonResponse {
	for _, attr := range attrs {
		if strings.TrimSpace(attr) == "" {
			panic("JSON: Invalid attribute: empty string")
		}
		j.result[attr] = reflect.ValueOf(j.obj).Elem().FieldByName(attr).Interface()
	}
	return j
}

/*
	Exclude attributes
*/
func (j *JsonResponse) Exclude(attrs ...string) *JsonResponse {
	fields := reflect.TypeOf(j.obj).Elem()
	for i := 0; i < fields.NumField(); i++ {
		currentAttr := fields.Field(i).Name
		if !slices.Contains(attrs, currentAttr) {
			j.result[currentAttr] = reflect.ValueOf(j.obj).Elem().FieldByName(currentAttr).Interface()
		}
	}
	return j
}
