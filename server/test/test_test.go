package test

import (
	"fmt"
	"server/library/command"
	"testing"
)

func TestBase62(t *testing.T) {
	mobile := 15821793512
	str := command.Encode62(mobile)
	fmt.Println(str) //output: HGkdFw
	num := command.Decode62(str)
	fmt.Println(num) //output: 15821793512
}
