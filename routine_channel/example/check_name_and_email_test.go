package tig_routine_channel_example_test

import (
	gtest "github.com/og/x/test"
	"log"
	"testing"
	"time"
)

var _= `
# 并发查询邮箱和姓名是否存在
`

type User struct {
	Name string
	Email string
}
var nameList = []string{"abc", "123", "xyz"}
var emailList = []string{"abc@x.com", "123@x.com", "xyz@x.com"}

func CheckName(queryName string, has chan<- bool) {
	log.Print(`CheckName("` + queryName + `")`)
	for _, name := range nameList {
		time.Sleep(time.Millisecond*500) // 模拟查询数据IO的时间差异
		log.Print([]string{"CheckName:", name, queryName})
		if name == queryName {
			has <- true
			return
		}
	}
	has <- false
}
func CheckEmail(queryEmail string, has chan<-bool) {
	log.Print(`CheckEmail("` + queryEmail + `")`)
	for _, email := range emailList {
		time.Sleep(time.Millisecond*500) // 模拟查询数据IO的时间差异
		log.Print([]string{"CheckEmail:", email, queryEmail})
		if email == queryEmail {
			has <- true
			return
		}
	}
	has <- false
}
func CheckNameAndEmail(user User) (has bool) {
	hasCh := make(chan bool, 2)
	go CheckName(user.Name, hasCh)
	go CheckEmail(user.Email, hasCh)
	// 根据容量（2）获取2次 hasCh 的值
	for i:=0;i<cap(hasCh);i++{
		if <-hasCh {
			return true
		}
	}
	return false
}
func TestCheckNameAndEmail_1(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(
		CheckNameAndEmail(User{
			Name:  "abc",
			Email: "ooo",
		}),
		true,
	)
}
func TestCheckNameAndEmail_2(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(
		CheckNameAndEmail(User{
			Name:  "123",
			Email: "ooo",
		}),
		true,
	)
}

func TestCheckNameAndEmail_3(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(
		CheckNameAndEmail(User{
			Name:  "xyz",
			Email: "ooo",
		}),
		true,
	)
}
func TestCheckNameAndEmail_4(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(
		CheckNameAndEmail(User{
			Name:  "ooo",
			Email: "ooo",
		}),
		false,
	)
}


