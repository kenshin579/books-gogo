package multiset

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

// 새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() MultiSet {
	return make(MultiSet)
}

// Insert 함수는 집합에 val을 추가한다.
func Insert(m MultiSet, val string) {
	count := m[val]
	m[val] = count + 1
}

func InsertFunc(m MultiSet) func(val string) {
	return func(val string) {
		Insert(m, val)
	}
}

//인자 고정 - Insert 함수의 첫 인자인 m을 고정한 함수
func BindMap(f SetOp, m MultiSet) func(val string) {
	return func(val string) {
		f(m, val)
	}
}

// Erase 함수는 집합에서 val을 제거한다. 집합에 val이 없는
// 경우에는 아무 일도 일어나지 않는다.
func Erase(m MultiSet, val string) {
	count := m[val]
	if count == 1 {
		delete(m, val)
	} else {
		m[val] = count - 1
	}
}

// Count 함수는 집합에 val이 들어 있는 횟수를 구한다.
func Count(m MultiSet, val string) int {
	return m[val]
}

// String 함수는 집합에 들어 있는 원소들을 { } 안에 빈 칸으로
// 구분하여 넣은 문자열을 반환한다.
func String(m MultiSet) string {
	var result string
	result = "{ "

	for k, v := range m {
		for i := 0; i < v; i++ {
			result += k + " "
		}
	}
	result += "}"
	return result
}
