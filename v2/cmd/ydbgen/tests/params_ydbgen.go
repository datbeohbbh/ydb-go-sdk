// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
	"github.com/yandex-cloud/ydb-go-sdk/v2/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func (p *Params) QueryParameters() *table.QueryParameters {
	var v0 ydb.Value
	{
		vp0 := ydb.OptionalValue(ydb.UTF8Value(p.Name))
		v0 = vp0
	}
	var v1 ydb.Value
	{
		vp0 := ydb.Uint32Value(ydbConvI16ToU32(p.Int16ToUint32))
		v1 = vp0
	}
	var v2 ydb.Value
	{
		vp0 := ydb.Int64Value(int64(p.IntToInt64))
		v2 = vp0
	}
	return table.NewQueryParameters(
		table.ValueParam("$name", v0),
		table.ValueParam("$int16_to_uint32", v1),
		table.ValueParam("$int_to_int64", v2),
	)
}

func ydbConvI16ToU32(x int16) uint32 {
	if x < 0 {
		panic("ydbgen: convassert: conversion of negative int16 to uint32")
	}
	return uint32(x)
}
