package main

import (
	"fmt"
	"testing"

	"github.com/godemo/common"
)

func TestReverse(t *testing.T) {
	fmt.Println(common.ReverseRunes("!oG ,olleH"))
	fmt.Println(common.ReverseRunes("dddaaa"))
}
