package str

import (

	"testing"
)

func TestMerge(t *testing.T) {
	Merge()
}

func BenchmarkMerge(b *testing.B) {
	for i:=0;i<b.N;i++{
		Merge()
	}
}

func BenchmarkMergeByFmt(b *testing.B) {
	for i:=0;i<b.N;i++{
		MergeByFmt()
	}
}

func BenchmarkMergeByJoin(b *testing.B) {
	for i:=0;i<b.N;i++{
		MergeByJoin()
	}
}

func BenchmarkMergeByBuffer(b *testing.B) {
	for i:=0;i<b.N;i++{
		MergeByBuffer()
	}
}

func BenchmarkMergeByBuilder(b *testing.B) {
	for i:=0;i<b.N;i++{
		MergeByBuilder()
	}
}

func BenchmarkMergeByBuilders(b *testing.B) {
	p:= initStrings(1000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		MergeByBuilders(p)
	}
}

func BenchmarkStringBuilder1000(b *testing.B) {
	p:= initStrings(1000)
	caps:=1000*len(S)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		StringBuilder(p,caps)
	}
}