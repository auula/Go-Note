package str

import (
	"bytes"
	"fmt"
	"strings"
)

const S = "今晚打老虎"

func Merge() string {
	var s string
	s+="今晚"+"打"+"老虎"+"\n"
	s+="今晚"+"打"+"老虎"+"\n"
	s+="今晚"+"打"+"老虎"+"\n"
	return s
}

func MergeByFmt() string {
	return fmt.Sprint("今晚","打","老虎","\n","今晚","打","老虎","\n","今晚","打","老虎","\n")

}

func MergeByJoin() string {
	s:=[]string{"今晚","打","老虎","\n","今晚","打","老虎","\n","今晚","打","老虎","\n"}
	return strings.Join(s,"")
}


func MergeByBuffer() string {
	var s bytes.Buffer
	s.WriteString("今晚")
	s.WriteString("打")
	s.WriteString("老虎")
	s.WriteString("\n")
	s.WriteString("今晚")
	s.WriteString("打")
	s.WriteString("老虎")
	s.WriteString("\n")
	s.WriteString("今晚")
	s.WriteString("打")
	s.WriteString("老虎")
	s.WriteString("\n")
	return s.String()
}

func MergeByBuilder() string {
	var s strings.Builder
	s.WriteString("今晚")
	s.WriteString("打")
	s.WriteString("老虎")
	s.WriteString("\n")
	s.WriteString("今晚")
	s.WriteString("打")
	s.WriteString("老虎")
	s.WriteString("\n")
	s.WriteString("今晚")
	s.WriteString("打")
	s.WriteString("老虎")
	s.WriteString("\n")
	return s.String()
}

func initStrings(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = S
	}
	return s
}

func MergeByBuilders(s []string) string {
	var b strings.Builder
	l := len(s)
	for i := 0; i < l; i++ {
		b.WriteString(s[i])
	}
	return b.String()
}

// StringBuilder 优化后的代码
func StringBuilder(p []string,caps int) string {
	var b strings.Builder
	l:=len(p)
	b.Grow(caps)
	for i:=0;i<l;i++{
		b.WriteString(p[i])
	}
	return b.String()
}

