package split

import (
	"reflect"
	"testing"
)

// 测试存在分隔符时的split函数
func TestSplit(t *testing.T) {
	str := "a:b:c"
	got := Split(str, ":")
	want := []string{"a", "b", "c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("want: %v, got: %v\n", want, got)
	}
}

// 测试不存在分隔符时的split函数
func TestNonSplit(t *testing.T) {
	str := "a:b:c"
	got := Split(str, "*")
	want := []string{"a:b:c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("want: %v, got: %v\n", want, got)
	}
}
