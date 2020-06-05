package split

import (
	"testing"
)

func BenchmarkSplit(b *testing.B) {
	b.Log("这是一个基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
