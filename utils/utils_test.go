package utils

import (
	"testing"
)

func TestStructToValues_BasicFields(t *testing.T) {
	type Req struct {
		AppId   string `json:"app_id"`
		Method  string `json:"method"`
		Charset string `json:"charset,omitempty"`
		Skip    string `json:"-"`
		Empty   string `json:"empty,omitempty"`
	}

	got := StructToValues(&Req{
		AppId:   "2021000000000001",
		Method:  "alipay.trade.create",
		Charset: "UTF-8",
		Skip:    "should-not-appear",
		Empty:   "",
	})

	if got.Get("app_id") != "2021000000000001" {
		t.Errorf("app_id = %q, want %q", got.Get("app_id"), "2021000000000001")
	}
	if got.Get("method") != "alipay.trade.create" {
		t.Errorf("method = %q, want alipay.trade.create", got.Get("method"))
	}
	if got.Get("charset") != "UTF-8" {
		t.Errorf("charset = %q, want UTF-8", got.Get("charset"))
	}
	if _, ok := got["-"]; ok {
		t.Errorf(`json:"-" field leaked into output: %v`, got)
	}
	if got.Get("empty") != "" {
		t.Errorf("empty string should not be encoded, got %q", got.Get("empty"))
	}
}

func TestStructToValues_NilAndNonStruct(t *testing.T) {
	if v := StructToValues(nil); len(v) != 0 {
		t.Errorf("nil input should yield empty values, got %v", v)
	}

	var ptr *struct{ X string }
	if v := StructToValues(ptr); len(v) != 0 {
		t.Errorf("nil pointer should yield empty values, got %v", v)
	}

	if v := StructToValues(42); len(v) != 0 {
		t.Errorf("non-struct input should yield empty values, got %v", v)
	}
}

func TestStructToValues_NestedAndSlice(t *testing.T) {
	type Goods struct {
		Id    string `json:"goods_id"`
		Name  string `json:"goods_name"`
		Count int64  `json:"quantity"`
	}
	type Req struct {
		OutTradeNo string   `json:"out_trade_no"`
		GoodsVal   []Goods  `json:"goods_detail_val,omitempty"`
		Tags       []string `json:"tags,omitempty"`
	}

	got := StructToValues(&Req{
		OutTradeNo: "T001",
		GoodsVal: []Goods{
			{Id: "g-1", Name: "ipad", Count: 2},
			{Id: "g-2", Name: "apple", Count: 1},
		},
		Tags: []string{"a", "b"},
	})

	if got.Get("out_trade_no") != "T001" {
		t.Errorf("out_trade_no = %q", got.Get("out_trade_no"))
	}
	// 值类型结构体切片：递归展开为索引化 key
	if got.Get("goods_detail_val[0][goods_id]") != "g-1" {
		t.Errorf("goods_detail_val[0][goods_id] = %q", got.Get("goods_detail_val[0][goods_id]"))
	}
	if got.Get("goods_detail_val[1][quantity]") != "1" {
		t.Errorf("goods_detail_val[1][quantity] = %q", got.Get("goods_detail_val[1][quantity]"))
	}
	// 基本类型切片：以 JSON 字符串形式整体编码
	if got.Get("tags") != `["a","b"]` {
		t.Errorf(`tags = %q, want ["a","b"]`, got.Get("tags"))
	}
}

func TestStructToValues_PointerSliceEncodedAsJSON(t *testing.T) {
	// 注意：指针切片的元素 Kind 是 reflect.Ptr（不是 Struct），
	// 因此按基本类型切片路径以 JSON 整体编码。
	// 这是当前实现的行为；Alipay 业务参数通常放在 biz_content（字符串）中，
	// 所以 PublicRequestParameters 展平时不会触及该分支。
	type Goods struct {
		Id string `json:"goods_id"`
	}
	type Req struct {
		Items []*Goods `json:"items,omitempty"`
	}
	got := StructToValues(&Req{Items: []*Goods{{Id: "g1"}}})
	if got.Get("items") != `[{"goods_id":"g1"}]` {
		t.Errorf(`items = %q, want [{"goods_id":"g1"}]`, got.Get("items"))
	}
}

func TestStructToValues_PointerStructNestedNil(t *testing.T) {
	type Inner struct {
		X string `json:"x"`
	}
	type Req struct {
		A string `json:"a"`
		B *Inner `json:"b,omitempty"`
	}

	got := StructToValues(&Req{A: "hi", B: nil})
	if got.Get("a") != "hi" {
		t.Errorf("a = %q", got.Get("a"))
	}
	if _, ok := got["b"]; ok {
		t.Errorf("nil pointer struct should not be encoded, got %v", got)
	}
	if _, ok := got["b[x]"]; ok {
		t.Errorf("nil pointer struct inner fields should not be encoded, got %v", got)
	}
}
