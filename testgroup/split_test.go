package split

import (
	"reflect"
	"testing"
)

type TestT struct {
	Str  string
	Sep  string
	Want []string
}

func TestMultiSplit(t *testing.T) {
	// 使用map来构建一个测试集
	var tests = map[string]TestT{
		"Normal": TestT{"a:b:c", ":", []string{"a", "b", "c"}},
		"None":   TestT{"a:b:c", "*", []string{"a:b:c"}},
		"Multi":  TestT{"abcfbcabcd", "bc", []string{"a", "f", "a", "d"}},
		"Number": TestT{"1231", "1", []string{"", "23", ""}},
	}

	for name, item := range tests {
		ret := Split(item.Str, item.Sep)
		if ok := reflect.DeepEqual(item.Want, ret); !ok {
			t.Errorf("测试用例: [%v], 期望值 %v, 实际值: %v\n", name, item.Want, ret)
		}
	}
}

// 使用子测试完成测试套
func TestMultiSplit2(t *testing.T) {
	// 使用map来构建一个测试集
	var tests = map[string]TestT{
		"Normal": TestT{"a:b:c", ":", []string{"a", "b", "c"}},
		"None":   TestT{"a:b:c", "*", []string{"a:b:c"}},
		"Multi":  TestT{"abcfbcabcd", "bc", []string{"a", "f", "a", "d"}},
		"Number": TestT{"1231", "1", []string{"", "23", ""}},
	}

	for name, item := range tests {
		t.Run(name, func(tIn *testing.T) {
			ret := Split(item.Str, item.Sep)
			if ok := reflect.DeepEqual(item.Want, ret); !ok {
				tIn.Errorf("测试用例: [%v], 期望值 %v, 实际值: %v\n", name, item.Want, ret)
			}
		})

	}
}
