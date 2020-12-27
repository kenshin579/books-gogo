package multiset

// 새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int {
	return make(map[string]int)
}

// Insert 함수는 집합에 val을 추가한다.
func Insert(m map[string]int, val string) {
	count := m[val]
	m[val] = count + 1
}

// Erase 함수는 집합에서 val을 제거한다. 집합에 val이 없는
// 경우에는 아무 일도 일어나지 않는다.
func Erase(m map[string]int, val string) {
	count := m[val]
	if count == 1 {
		delete(m, val)
	} else {
		m[val] = count - 1
	}
}

// Count 함수는 집합에 val이 들어 있는 횟수를 구한다.
func Count(m map[string]int, val string) int {
	return m[val]
}

// String 함수는 집합에 들어 있는 원소들을 { } 안에 빈 칸으로
// 구분하여 넣은 문자열을 반환한다.
func String(m map[string]int) string {
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
