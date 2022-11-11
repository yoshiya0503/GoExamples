package functions

import "fmt"

type Function func()

func (f Function) Apply() {
	f()
}

func add(a, b int) int {
	return a + b
}

func RunExample() {
	var sum Function = func() { fmt.Println(add(10, 20)) }
	sum()
	sum.Apply()
	Function(func() { fmt.Println("call function") }).Apply()
}
