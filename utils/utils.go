package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// StructToValues 将结构体转换为 url.Values
func StructToValues(s interface{}) url.Values {
	values := url.Values{}
	if s == nil {
		return values
	}

	val := reflect.ValueOf(s)
	// 处理指针类型
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return values
		}
		val = val.Elem()
	}

	// 确保是结构体
	if val.Kind() != reflect.Struct {
		return values
	}

	structToValues(val, "", values)
	return values
}

// structToValues 递归处理结构体字段
func structToValues(val reflect.Value, prefix string, values url.Values) {
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 跳过未导出的字段
		if !fieldType.IsExported() {
			continue
		}

		// 获取字段标签
		tag := fieldType.Tag.Get("json") // 默认使用 json 标签
		if tag == "-" {
			continue
		}

		// 处理标签名称
		name := fieldType.Name
		if tag != "" {
			if idx := strings.Index(tag, ","); idx != -1 {
				name = tag[:idx]
			} else {
				name = tag
			}
		}

		// 构建完整的键名
		fullName := name
		if prefix != "" {
			fullName = prefix + "[" + name + "]"
		}

		// 处理字段值
		addFieldToValues(field, fullName, values)
	}
}

// addFieldToValues 将字段添加到 url.Values
func addFieldToValues(field reflect.Value, key string, values url.Values) {
	// 处理接口类型
	if field.Kind() == reflect.Interface {
		if !field.IsNil() {
			field = field.Elem()
		}
	}

	// 处理指针类型
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			return
		}
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.String:
		if str := field.String(); str != "" {
			values.Add(key, str)
		}

	case reflect.Bool:
		values.Add(key, fmt.Sprintf("%t", field.Bool()))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.Add(key, fmt.Sprintf("%d", field.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		values.Add(key, fmt.Sprintf("%d", field.Uint()))

	case reflect.Float32, reflect.Float64:
		values.Add(key, fmt.Sprintf("%.2f", field.Float()))

	case reflect.Slice, reflect.Array:
		if field.Len() > 0 {
			// 处理基本类型的切片
			if field.Type().Elem().Kind() != reflect.Struct {
				bytes, err := json.Marshal(field.Interface())
				if err == nil {
					values.Add(key, string(bytes))
				}
			} else {
				// 处理结构体切片
				for i := 0; i < field.Len(); i++ {
					structToValues(field.Index(i), fmt.Sprintf("%s[%d]", key, i), values)
				}
			}
		}

	case reflect.Map:
		if field.Len() > 0 {
			bytes, err := json.Marshal(field.Interface())
			if err == nil {
				values.Add(key, string(bytes))
			}
		}

	case reflect.Struct:
		// 特殊处理时间类型
		if t, ok := field.Interface().(time.Time); ok {
			values.Add(key, t.Format(time.RFC3339))
		} else {
			// 递归处理嵌套结构体
			structToValues(field, key, values)
		}
	}
}
