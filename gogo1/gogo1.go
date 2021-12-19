package gogo1

import gogotest4 "github.com/hirokuma/gogo-test4"

var (
	_val int
)

func SetValue(val int) {
	_val = val
	gogotest4.SetValue(val)
}

func GetValue() (int, int) {
	return _val, gogotest4.GetValue()
}
