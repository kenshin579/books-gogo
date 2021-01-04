package eval

// Eval returns the evaluation result of the given expr.
import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func Eval(expr string) int {
	log.Printf("expr: %v", expr)
	var ops []string //연산자 저장하는 곳
	var nums []int   //숫자 및 연산결과 저장하는 곳

	//마지막 값을 pop해서 반환함
	popNum := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	//nums에 있는 값을 계산하는 곳임 (a + b -> c)
	reduce := func(higher string) {
		log.Printf("  higher:%v ops:%v\n", higher, ops)
		for len(ops) > 0 { //연산자가 있는 경우에만 진행함
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				// 목록에 없는 연산자이므로 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := popNum(), popNum()
			log.Printf("b:%v a:%v", b, a)

			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}

	//여기서부터 시작됨
	for _, token := range strings.Split(expr, " ") {
		log.Printf("ops:%v nums:%v token:%v", ops, nums, token)

		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			// 덧셈과 뺄셈 이상의 우선순위를 가진 사칙연산 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			// 곱셈과 나눗셈 이상의 우선순위를 가진 것은 둘 뿐
			reduce("*/")
			ops = append(ops, token)
		case ")":
			// 닫는 괄호는 여는 괄호까지 계산을 하고 제거
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce("+-*/")
	return nums[0]
}
