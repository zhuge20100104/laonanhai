package setup

import (
	"reflect"
	"testing"
)

func setUpTestCase(t *testing.T) func(t *testing.T) {
	t.Logf("这里是整个Case的Setup代码\n")

	return func(t *testing.T) {
		t.Logf("这里是整个Case的Teardown代码\n")
	}
}

func setUpSubTest(t *testing.T) func(t *testing.T) {
	t.Logf("这里是某个子Case的Setup代码\n")
	return func(t *testing.T) {
		t.Logf("这里是某个子Case的Teardown代码\n")
	}
}

// TestSplit Split方法的测试函数
func TestSplit(t *testing.T) {
	type Test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]Test{
		"Simple":       Test{"a:b:c", ":", []string{"a", "b", "c"}},
		"Wrong Sep":    Test{"a:b:c", ",", []string{"a:b:c"}},
		"Multiple Sep": Test{"abcd", "bc", []string{"a", "d"}},
		"Leading Sep":  Test{"沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	tearDown := setUpTestCase(t)
	defer tearDown(t)

	for name, tCase := range tests {
		t.Run(name, func(t *testing.T) {
			subTearDown := setUpSubTest(t)
			defer subTearDown(t)
			got := Split(tCase.input, tCase.sep)
			if !reflect.DeepEqual(got, tCase.want) {
				t.Errorf("测试用例[%v]执行失败, expected: %v, actual: %v\n", name,
					tCase.want, got)
			}
		})
	}

}
