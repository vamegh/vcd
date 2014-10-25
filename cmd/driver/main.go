package main

import (
	"github.com/houst0n/vcd"
)

func main() {
	v, _ := vcd.NewClient()
	v.ShowToken()
}
