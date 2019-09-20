package main

import (
	"github.com/zhongzc/lexerGen/codegen"
	"github.com/zhongzc/lexerGen/ruledef"
	"os"
)

func main() {
	file, err := os.Open("tmp.lx")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = ruledef.BuildAll(file, "lexertmp", codegen.NewGoGen())
	if err != nil {
		panic(err)
	}
}
