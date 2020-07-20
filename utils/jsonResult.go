package utils

import (
	"math"
)

type result struct {
	Code string
	Message string
	Data interface{}
	Page int
	PageSize int
	Count int
	CountPage int
	Extra interface{}
}

const (
	//CodeSuccess 成功
	CodeSuccess = "000";
	//CodeFail 失败
	CodeFail = "001";
	//CodeNoLogin 未登录
	CodeNoLogin = "002"
	//MsgSuccess 操作成功
	MsgSuccess = "操作成功"
	//MsgFail 操作失败
	MsgFail = "操作失败"
	//MsgNoLogin 请先登录
	MsgNoLogin = "请先登录"
	//MsgBuzy 系统繁忙
	MsgBuzy = "系统繁忙"
)

//NewResult 返回新jsonResult
func NewResult() *result {
	return new(result)
}

//Success 成功放回
func Success(message string) *result {
	rs := NewResult()
	rs.Code = CodeSuccess
	rs.Message = message
	return rs
}

//Error 失败
func Error(message string) *result {
	rs := NewResult()
	rs.Code = CodeFail
	rs.Message = message
	return rs
}

//Item 普通返回
func Item(data interface{}) *result {
	rs := NewResult()
	rs.Code = CodeSuccess
	rs.Message = MsgSuccess
	rs.Data = data
	return rs
}

//Paginate  分页返回($data, $count, $page, $pageSize, $extra = [])
func Paginate(data interface{}, count int, page int, pageSzie int, extra interface{}) *result {
	rs := NewResult()
	rs.Code = CodeSuccess
	rs.Message = MsgSuccess
	rs.Data = data
	rs.Page = page
	rs.PageSize = pageSzie
	rs.Extra = extra
	rs.Count = count
	rs.CountPage = int(math.Ceil(float64(count) / float64(pageSzie)))
	return rs
}

