package polymorph

import (
	"fmt"
	"reflect"
)

func ExampleTotalArea() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
	}))
	// Output: 50
}

func ExampleRectangleCircum() {
	r := RectangleCircum{Rectangle{3, 4}}
	fmt.Println(r.Area())
	fmt.Println(r.Circum())
	// Output:
	// 12
	// 14
}

func ExampleWrongRectangle() {
	r := WrongRectangle{Rectangle{3, 4}}
	fmt.Println(r.Area())
	// Output: 24
}

func ExampleTotalArea_moreTypes() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
		RectangleCircum{Rectangle{8, 9}},
		WrongRectangle{Rectangle{1, 2}},
	}))
	// Output: 126
}

//내장된 구조체가 있는지는 구조체에서 내장된 구조체의 이름으로 필드를 찾은 다음에 anonymous 필드를 찾아보면 된다
func Example_ReflectTypeOf로_어떤_자료형이_주어진_인터페이스를_구현하고_있는지_여부_확인() {
	//impl := reflect.TypeOf(RectangleCircum{}).Implements(
	//	reflect.TypeOf((*Shape)(nil)).Elem(),
	//	)

	//RectangleCircum{} 빈 객체를 만들지 않고 하는 방법
	impl := reflect.TypeOf((*RectangleCircum)(nil)).Elem().Implements(reflect.TypeOf((*Shape)(nil)).Elem())
	fmt.Println(impl)

	field, ok := reflect.TypeOf(RectangleCircum{}).FieldByName("Rectangle")
	emb := ok && field.Anonymous && field.Type == reflect.TypeOf(Rectangle{})
	fmt.Println(emb)
	//Output:
	//true
	//true
}
