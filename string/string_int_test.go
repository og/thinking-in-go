package tig_string

import (
	gtest "github.com/og/x/test"
	"strconv"
	"testing"
)

var _= `
字符串数字直接是不能直接转换,并且使用 string(int) 时编译期不会报错，所以千万要注意不要使用 string(int)。
`
func TestIntString(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	age := 18
	as.Equal(strconv.FormatInt(int64(age), 10), "18")
	as.Equal(string(age), "\x12")
}
