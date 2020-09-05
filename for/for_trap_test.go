package tig_for_trap

import (
	"log"
	"testing"
	"time"
)

var _ = `
for 循环中的 key, value 在遇到 routine 等"延迟"场景时如果直接使用 key value
可能会导致获取的是 slice 最后一项的 key value
`

func TestNormalFor (t *testing.T) {
	names := []string{"nimo","og","github"}
	for index, name := range names {
		log.Print(index, name)
	}
	// 0 nimo
	// 1 og
	// 2 github
}

func TestForRoutine(t *testing.T) {
	names := []string{"nimo","og","github"}
	for index, name := range names {
		go func() {
			log.Print(index, name)
		}()
	}
	// 2 github
	// 2 github
	// 2 github
	// log.Print(name) 时 name 的值始终是最后一项 "github"
	// 可以理解为启动 routine 导致了"延迟"，"延迟"后获取的 name 时

	time.Sleep(time.Second) //等待routine执行完成
}



func TestForRoutineCopyValue(t *testing.T) {
	names := []string{"nimo","og","github"}
	for index, name := range names {
		// 通过传参复制值解决取值错误（但这样还是可能在内部使用 index, name 并且编译通过）
		go func(i int, n string) {
			log.Print(i, n)
		}(index, name)
	}
	// 1 og
	// 2 github
	// 0 nimo
	// 乱序是因为 routine 的执行是就是乱序的，这是正常的

	time.Sleep(time.Second) //等待routine执行完成
}

func LogName(index int, name string) {
	log.Print(index, name)
}
func TestForRoutineUseOtherFunc(t *testing.T) {
	names := []string{"nimo","og","github"}
	for index, name := range names {
		// 使用不能直接获取 name 变量的函数来解决匿名函数中可以直接获取name的问题
		go LogName(index, name)
	}
	// 1 og
	// 2 github
	// 0 nimo
	// 乱序是因为 routine 的执行是就是乱序的，这是正常的

	time.Sleep(time.Second) //等待routine执行完成
}