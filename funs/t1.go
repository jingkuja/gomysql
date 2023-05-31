package funs

var Mb = "123456"

func Max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func Getnum(x int) func() int {

	return func() int {
		return x
	}

}
