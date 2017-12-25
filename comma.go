package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(commaChanged("123456789"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaChanged(str string) string {
	if len(str) <= 3 {
		return str
	}
	var buf bytes.Buffer
	var tem []byte
	s := []byte(str)
	comm := bytes.IndexByte(s, '.')
	if comm != -1 {
		tem = s[comm:] //小数部分
		s = s[:comm]   //整数部分
	}
	if strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+") {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	for n := len(s); n > 0; {
		m := n % 3
		if m > 0 {
			buf.Write(s[:m])
			buf.WriteString(",")
			s = s[m:n]
			n -= m
		}
		if n > 3 {
			fmt.Fprintf(&buf, "%s,", s[:3])
			s = s[3:n]
			n = n - 3
		} else {
			fmt.Fprintf(&buf, "%s", s[:3])
			s = s[3:n]
			n = n - 3
		}
	}
	buf.Write(tem)
	return buf.String()
}
