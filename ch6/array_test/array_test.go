package array_test

import "testing"

// c := [2][2]int[{1,2},{3,4}} 多维数组初始化

func TestArrayInit(t *testing.T) {
	var arr [3]int             //声明并初始化为默认零值
	arr1 := [4]int{1, 2, 3, 4} //声明同时初始化
	arr2 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestMultiArray(t *testing.T) {
	arr4 := [2][2]int{{1, 2}, {3, 4}} //多维数组初始化
	for i := 0; i < len(arr4[0]); i++ {
		for j := 0; j < len(arr4[1]); j++ {
			t.Log(arr4[i][j])
		}

	}
	t.Log(arr4)
}

func TestArraySection(t *testing.T) {
	arr5 := [...]int{1, 2, 3, 4, 5}
	arr5_sec1 := arr5[:3]
	arr5_sec2 := arr5[3:]
	//arr5_sec3 := arr5[-1:]  wrong
	t.Log(arr5_sec1)
	t.Log(arr5_sec2)
}
