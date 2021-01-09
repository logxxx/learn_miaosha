package common

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

//根据结构提中sql标签，映射到数据结构体中，并且转换类型
func DataToStructByTagSql(data map[string]string, obj interface{}) {
	objValue := reflect.ValueOf(obj).Elem()
	for i:=0; i<objValue.NumField(); i++ {
		//获取sql对应值
		value := data[objValue.Type().Field(i).Tag.Get("sql")]
		//获取对应字段类型
		name := objValue.Type().Field(i).Name
		//获取对应字段类型
		structFieldType := objValue.Field(i).Type()
		//获取变量类型，也可以直接写string
		val := reflect.ValueOf(value)
		var err error
		if structFieldType != val.Type() {
			//类型转换
			val, err = TypeConversion(value, structFieldType.Name()) //类型转换
			if err != nil {

			}
		}
		//设置类型值
		objValue.FieldByName(name).Set(val)
	}
}

func TypeConversion(value string, ntype string) (reflect.Value, error) {
	switch ntype {
	case "string":
		return reflect.ValueOf(value), nil
	case "time.Time":
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		if err != nil {
			return reflect.ValueOf(t), err
		}
		return reflect.ValueOf(t), nil
	case "int":
		i, err := strconv.Atoi(value)
		if err != nil {
			return reflect.ValueOf(i), err
		}
		return reflect.ValueOf(i), nil
	case "int64":
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return reflect.ValueOf(i), err
		}
		return reflect.ValueOf(i), nil
	default:
		return reflect.ValueOf(nil), errors.New("无法识别的类型")
	}
}