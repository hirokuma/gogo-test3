package gogo2

import gogotest4 "github.com/hirokuma/gogo-test4"

var (
	_val int
)

func SetValue(val int) {
	_val = val * 2
	gogotest4.SetValue(val * 2)
}

func GetValue() (int, int) {
	return _val, gogotest4.GetValue()
}
