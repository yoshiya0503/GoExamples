package main

import (
	"fmt"

	"github.com/yoshiya0503/GoExamples/pkg/concurrency"
	"github.com/yoshiya0503/GoExamples/pkg/functions"
	"github.com/yoshiya0503/GoExamples/pkg/generics"
	"github.com/yoshiya0503/GoExamples/pkg/stacktrace"
)

func main() {
	fmt.Println("------------functions------------------")
	functions.RunExample()
	fmt.Println("------------collection------------------")
	generics.RunLibExample()
	fmt.Println("------------generics------------------")
	generics.RunExample()
	fmt.Println("------------server------------------")
	functions.RunServer()
	fmt.Println("------------stacktrace------------------")
	stacktrace.RunExample()
	fmt.Println("------------concurrency generator------------------")
	// concurrency.RunGenerator()
	fmt.Println("------------concurrency pipeline------------------")
	concurrency.RunPipeline()
	// concurrency.RunForSelect()
}
