package tig_string

import (
	"errors"
	gtest "github.com/og/x/test"
	"testing"
)
var _= `
go 字符串的编码是 Unicode，不了解的可以去搜索 Unicode 编码了解字符串编码。
或者在 https://tool.oschina.net/encode 查看一些英文和中文的文字编码。
可以想象成字符串中每个字符都对应一个数字。
在 go 中 byte 等于 uint8，rune 等于 int32
这决定了像中文，蒙古文这些非英文的字符在 []byte 和 []rune 的存储方式不一样。
可以简单理解成 rune 代表一个字， byte 不一定代表一个字。
查看实例帮助理解：
`
func TestByteRuneString(t *testing.T) {
	as := gtest.NewAS(t)
	{
		bList := []byte("ab")
		as.Equal(len(bList), 2)
		as.Equal(bList[0], uint8(97))
		as.Equal(bList[1], uint8(98))
		s := string(bList)
		as.Equal(len(s), 2)
		as.Equal(s, "ab")
		runeList := []rune(s)
		as.Equal(len(runeList), 2)
		as.Equal(runeList[0], int32(97))
		as.Equal(runeList[1], int32(98))
	}
	{
		bList := []byte("a我")
		as.Equal(len(bList), 4)
		as.Equal(bList[0], uint8(97))
		as.Equal(bList[1], uint8(230))
		as.Equal(bList[2], uint8(136))
		as.Equal(bList[3], uint8(145))
		s := string(bList)
		as.Equal(len(s), 4)
		as.Equal(s, "a我")
		runeList := []rune(s)
		as.Equal(len(runeList), 2)
		as.Equal(runeList[0], int32(97))
		as.Equal(runeList[1], int32(25105))
	}
	// 使用 for range 遍历字符串是"安全"的
	{
		message := "a我"
		var count int
		for i, s := range message {
			// 此处 s 是 rune(int32)
			count++
			switch i {
			case 0:
				as.Equal(s, int32(97))
			case 1:
				as.Equal(s, int32(25105))
			default:
				panic(errors.New("can not happen"))
			}
		}
		as.Equal(count, 2)
	}
	// 使用 for len i++ 遍历字符串是不"安全"的,除非你的特意要遍历 []byte
	{
		message := "a我"
		var count int

		for i:=0;i<len(message);i++ {
			s := message[i] // 此处的 s 是 byte (uint8)
			count++
			switch i {
			case 0:
				as.Equal(s, uint8(97))
			case 1:
				as.Equal(s, uint8(230))
			case 2:
				as.Equal(s, uint8(136))
			case 3:
				as.Equal(s, uint8(145))
			default:
				panic(errors.New("can not happen"))
			}
		}
		as.Equal(count, 4)
	}
	// 忘记 byte 不是一个字是会写出错误的逻辑
	{
		message := "a我b"
		var newMessage []byte
		newMessage = append(newMessage, message[0])
		newMessage = append(newMessage, message[1])
		as.Equal(string(newMessage), "a\xe6")
	}
	// 遇到需要通过下标取字符串值时，应该转换为 []rune。因为你不能确定这个字符串永远不出现非英文
	{
		message := []rune("a我b")
		var newMessage []rune
		newMessage = append(newMessage, message[0])
		newMessage = append(newMessage, message[1])
		as.Equal(string(newMessage), "a我")
	}

}