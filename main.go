package main

import (
	"fmt"

	"github.com/hirokuma/gogo-test3/gogo1"
	"github.com/hirokuma/gogo-test3/gogo2"
)

func main() {
	gogo1.SetValue(10)
	ret1, ret2 := gogo1.GetValue()
	fmt.Printf("gogo1 %d, %d\n", ret1, ret2)
	gogo2.SetValue(10)
	ret1, ret2 = gogo2.GetValue()
	fmt.Printf("gogo2 %d, %d\n", ret1, ret2)
	ret1, ret2 = gogo1.GetValue()
	fmt.Printf("gogo1 %d, %d\n", ret1, ret2)
}
