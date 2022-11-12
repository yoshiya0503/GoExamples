package stacktrace

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

type I interface {
	A()
}

type S struct {
	I
}

type U func()

func (u U) A() {
	println("this is u")
}

func (s S) Call() {
	fmt.Println(s.I)
	s.A()
}

func Trace1(value int) (int, error) {
	if value == 0 {
		return 0, xerrors.Errorf("Trace1: %w", errors.New("input 0 value"))
	}
	return value, nil
}

func Trace2(value int) (int, error) {
	value, err := Trace1(value)
	if err != nil {
		return 0, xerrors.Errorf("Trace2: %w", err)
	}
	return value, nil
}

func Trace3(value int) (int, error) {
	value, err := Trace2(value)
	if err != nil {
		return 0, xerrors.Errorf("Trace3: %w", err)
	}
	return value, nil
}

func RunExample() {
	_, err := Trace3(0)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// a := S{I: U(func() {})}
	// a := S{}
	// a.Call()
}
