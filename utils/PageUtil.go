package utils

import (
	"gin-base/conf"
)

var pageSize uint

func init() {
	pageSize = *conf.PageSize
}

type PageResult struct {
	List      interface{} `json:"list"`
	Count     uint        `json:"count"`
	TotalPage uint        `json:"totalPage"`
	PageSize  uint        `json:"pageSize"`
}

func PageRes(list interface{}, count uint) *PageResult {
	return &PageResult{list, count, TotalPage(count), pageSize}
}

func TotalPage(count uint) uint {
	var totalPage uint
	if count%*conf.PageSize == 0 {
		totalPage = count / *conf.PageSize
	} else {
		totalPage = count / *conf.PageSize + 1
	}
	return totalPage
}

func PageToStart(page, size uint) uint {
	var start uint = 0
	if page > 0 {
		start = (page - 1) * size
	}
	return start
}
