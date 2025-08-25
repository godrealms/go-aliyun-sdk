package alipay

import (
	"fmt"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
	"net/url"
	"reflect"
	"strconv"
)

func (c *Client) Notify(values url.Values) (*types.Notify, error) {
	notify := &types.Notify{}
	err := BindURLValues(values, notify)
	if err != nil {
		return nil, err
	}
	// 验证签名
	verifyService := community.NewAlipayVerifyService(c.AlipayPublicKey)
	ok, err := verifyService.VerifyNotifySign(values)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("signature error")
	}

	return notify, nil
}

func BindURLValues(values url.Values, dst interface{}) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 获取 form tag
		tag := fieldType.Tag.Get("form")
		if tag == "" {
			continue
		}

		// 从 url.Values 中获取值
		value := values.Get(tag)
		if value == "" {
			continue
		}

		// 根据字段类型设置值
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int, reflect.Int64:
			if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
				field.SetInt(intVal)
			}
		case reflect.Float64:
			if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
				field.SetFloat(floatVal)
			}
		case reflect.Invalid:
		case reflect.Bool:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
		case reflect.Uintptr:
		case reflect.Float32:
		case reflect.Complex64:
		case reflect.Complex128:
		case reflect.Array:
		case reflect.Chan:
		case reflect.Func:
		case reflect.Interface:
		case reflect.Map:
		case reflect.Pointer:
		case reflect.Slice:
		case reflect.Struct:
		case reflect.UnsafePointer:
		}
	}

	return nil
}
