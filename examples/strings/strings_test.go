package strings

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Example_range() {
	for i, r := range "가나다" {
		fmt.Println(i, r)
	}
	fmt.Println(len("가나다"))
	// Output:
	//0 44032
	//3 45208
	//6 45796
	//9
}

func Example_range2() {
	for i, r := range "abc" {
		fmt.Println(i, r)
	}
	fmt.Println(len("abc"))
	// Output:
	//0 97
	//1 98
	//2 99
	//3
}

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	s := []byte("가나다")
	s[2]++
	fmt.Println(string(s))
	// Output:
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_strCat2() {
	s := "abc"
	s1 := fmt.Sprint(s, "def")
	s2 := fmt.Sprintf("%sdef", s)
	s3 := strings.Join([]string{s, "def"}, "")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// Output:
	// abcdef
	// abcdef
	// abcdef
}

func Example_strConv() {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	Use(err, k, f, s)
	i, err = strconv.Atoi("350")
	fmt.Println(i)
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k)
	f, err = strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
	s = strconv.Itoa(340)
	fmt.Println(s)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)

	var num int
	fmt.Sscanf("57", "%d", &num) //string -> int
	fmt.Println(num)

	//todo : 아래는 안되는 듯함
	//var str string
	//str = fmt.Sprint(3.14) //"3.14"
	//fmt.Println(str)
	//str = fmt.Sprintf("%x", 13402077)
	//fmt.Println(str)

	// Output:
	//350
	//13402077
	//13402077
	//3.14
	//340
	//cc7fdd
	//57
}

// To make sure not to know the result in compile time.
var s4 = time.Now().Format("20060102")

func BenchmarkSprintf4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s", s1, s2, s3, s4)
	}
}

func BenchmarkPlus4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = s1 + s2 + s3 + s4
	}
}

func BenchmarkSprint4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(s1, s2, s3, s4)
	}
}

func BenchmarkJoin4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s1, s2, s3, s4}, "")
	}
}

func BenchmarkBytes(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString(s1)
		buf.WriteString(s2)
		buf.WriteString(s3)
		buf.WriteString(s4)
		_ = buf.String()
	}
}

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
