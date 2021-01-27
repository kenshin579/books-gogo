package ref

import (
	"errors"
	"fmt"
)

func ExampleFieldNames() {
	s := struct {
		id   int
		Name string
		Age  int
	}{}
	fmt.Println(FieldNames(s))
	// Output: [id Name Age] <nil>
}

func ExampleAppendNilError() {
	f := func() {
		fmt.Println("called")
	}
	f2, err := AppendNilError(f, errors.New("test error"))
	fmt.Println("AppendNilError.err:", err)

	//f2는 interface{} 이여서 형 단언을 해서 실행해야 함
	fmt.Println(f2.(func() error)())
	// Output:
	// AppendNilError.err: <nil>
	// called
	// test error
}
