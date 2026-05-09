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
		case reflect.Bool:
			if boolVal, err := strconv.ParseBool(value); err == nil {
				field.SetBool(boolVal)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
				field.SetInt(intVal)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if uintVal, err := strconv.ParseUint(value, 10, 64); err == nil {
				field.SetUint(uintVal)
			}
		case reflect.Float32, reflect.Float64:
			if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
				field.SetFloat(floatVal)
			}
		}
	}

	return nil
}
