package maps

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sort"
	"testing"
)

func TestCount_deepEqual(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	log.Info("codeCount", codeCount)

	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func TestCount_if(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if len(codeCount) != 3 {
		t.Error("codeCount:", codeCount)
		t.Fatal("count should be 3 but:", len(codeCount))
	}
	if codeCount['가'] != 1 || codeCount['나'] != 2 || codeCount['다'] != 1 {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func ExampleCount_sort() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice //Int Slice로 정의함

	//codeCount에서 key 값만 extract
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	log.Info("keys", keys)
	sort.Sort(keys)
	log.Info("sorted keys", keys)

	for _, key := range keys {
		fmt.Printf("%c %d\n", key, codeCount[rune(key)])
		//fmt.Println(string(key), codeCount[rune(key)])
	}
	// Output:
	// 가 1
	// 나 2
	// 다 1
}
