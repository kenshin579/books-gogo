package eval

// Eval returns the evaluation result of the given expr.
import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

//todo : 이거 다시 스터디 하기
// The expression can have +, -, *, /, (, ) operators and
// decimal integers. Operators and operands should be
func Eval(expr string) int {

	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(higher string) {
		log.WithFields(log.Fields{
			"higher": higher,
			"ops":    ops,
		}).Info("reduce")
		for len(ops) > 0 {
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
			b, a := pop(), pop()

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
	for _, token := range strings.Split(expr, " ") {
		log.WithFields(log.Fields{
			"ops":   ops,
			"nums":  nums,
			"token": token,
		}).Info("")
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
