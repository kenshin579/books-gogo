package multiset

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// ReadFrom calls f for each line from r.
func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Example_MultiSet() {
	m := NewMultiSet()
	r := bytes.NewBufferString("1\n4\n5\n3\n2\n")
	ReadFrom(r, func(line string) {
		Insert(m, line)
	})

	//fmt.Println(String(m))
	fmt.Println(Count(m, "3"))

	//Output:
	//1
}

func Example_MultiSet_함수_추상화() {
	m := NewMultiSet()
	r := bytes.NewBufferString("1\n4\n5\n3\n2\n")
	ReadFrom(r, InsertFunc(m))

	//fmt.Println(String(m))
	fmt.Println(Count(m, "3"))

	//Output:
	//1
}

func Example_MultiSet_함수_추상화_even_more() {
	m := NewMultiSet()
	r := bytes.NewBufferString("1\n4\n5\n3\n2\n")
	ReadFrom(r, BindMap(Insert, m))

	//fmt.Println(String(m))
	fmt.Println(Count(m, "3"))

	//Output:
	//1
}
