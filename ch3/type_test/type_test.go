package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	//不支持隐式类型转换，即使是自定义的也不行
	var a int32 = 1
	var b int64
	// b = a(wrong)
	b = int64(a)
	var c MyInt
	// c = b(wrong)
	c = MyInt(b)
	t.Log(a, b, c)
}

//指针类型
//1.不支持指针运算
//2.string是值类型，其默认的初始化值为空字符串，而不是nil
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1(wrong)
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	//if s == ""{
	//
	//}
}
