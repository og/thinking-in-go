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
var userList = []User{
	{Name: "abc", Email: "abc@x.com"},
	{Name: "123", Email: "123@x.com"},
	{Name: "xyz", Email: "xyz@x.com"},
}
func CheckName(name string, has chan<- bool) {
	log.Print(`CheckName("` + name + `")`)
	for _, user := range userList {
		time.Sleep(time.Millisecond*500) // 模拟查询每次遍历需要0.5s延迟
		log.Print([]string{"CheckName:", user.Name, name})
		if user.Name == name {
			has <- true
			return
		}
	}
	has <- false
}
func CheckEmail(email string, has chan<-bool) {
	log.Print(`CheckEmail("` + email + `")`)
	for _, user := range userList {
		time.Sleep(time.Millisecond*500) // 模拟查询数据IO的时间差异
		log.Print([]string{"CheckEmail:", user.Email, email})
		if user.Email == email {
			has <- true
			return
		}
	}
	has <- false
}
/*
	因为email和name都是并发查询的特性，所以最少要查询2次有结果。（0.5s + 0.5s）
	如果是队列查询，则最少查询1次才会有结果。(0.5s)
	需要在查询效率和查询速度中取舍，决定使用队列查询还是并发查询。
	（实际工作中）
*/
func CheckNameAndEmail(user User) (has bool) {
	hasCh := make(chan bool, 2)
	go CheckName(user.Name, hasCh)
	go CheckEmail(user.Email, hasCh)
	for i:=0;i<cap(hasCh);i++{
		if <-hasCh {
			return true
		}
	}
	return false
}
// 耗时 0.5s
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
// 耗时 1s
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

// 耗时 1.5s
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
// 耗时 1.5s
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


