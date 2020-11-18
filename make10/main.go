package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	methods := []string{"+", "-", "*", "/"}
	for _, num := range list() {
		for _, op1 := range methods {
			for _, op2 := range methods {
				for _, op3 := range methods {
					if calc(calc(calc(num[0], num[1], op1), num[2], op2), num[3], op3) == 10 {
						print(num, op1, op2, op3)
					}
				}
			}
		}
	}
}

func print(num []int, op1, op2, op3 string) {
	fmt.Printf("%d %s %d %s %d %s %d\n", num[0], op1, num[1], op2, num[2], op3, num[3])
}

func calc(i, j int, m string) int {
	switch m {
	case "+":
		return i + j
	case "-":
		return i - j
	case "*":
		return i * j
	case "/":
		if i == 0 || j == 0 {
			return 0
		}
		return i / j
	default:
		panic("invalid")
	}
}

func list() [][]int {
	l := [][]int{}
	for i := 0; i <= 9999; i++ {
		num := fmt.Sprintf("%04d", i)
		nums := strings.SplitN(num, "", -1)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		c, _ := strconv.Atoi(nums[2])
		d, _ := strconv.Atoi(nums[3])
		l = append(l, []int{a, b, c, d})
	}
	return l
}
