package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Merge() string {
	var s string
	s+="php"+"是"+"最好的语言"+"\n"
	s+="php"+"是"+"最好的语言"+"\n"
	s+="php"+"是"+"最好的语言"+"\n"
	return s
}

func MergeByFmt() string {
	return fmt.Sprint("php","是","最好的语言","\n","php","是","最好的语言","\n","php","是","最好的语言","\n")

}

func MergeByJoin() string {
	s:=[]string{"php","是","最好的语言","\n","php","是","最好的语言","\n","php","是","最好的语言","\n"}
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