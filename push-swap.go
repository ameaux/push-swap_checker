package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pa(A []int, B []int) ([]int, []int) {
	if len(B) > 0 {
		A = append([]int{B[0]}, A...)
		B = B[1:]
	}
	return A, B
}

func pb(A []int, B []int) ([]int, []int) {
	if len(A) > 0 {
		B = append([]int{A[0]}, B...)
		A = A[1:]
	}
	return A, B
}

func sa(A []int) []int {
	if len(A) >= 2 {
		A[0], A[1] = A[1], A[0]
	}
	return A
}

func sb(B []int) []int {
	if len(B) >= 2 {
		B[0], B[1] = B[1], B[0]
	}
	return B
}

func ss(A []int, B []int) ([]int, []int) {
	A = sa(A)
	B = sb(B)
	return A, B
}

func ra(A []int) []int {
	if len(A) >= 2 {
		A = append(A[1:], A[0])
	}
	return A
}

func rb(B []int) []int {
	if len(B) >= 2 {
		B = append(B[1:], B[0])
	}
	return B
}

func rr(A []int, B []int) ([]int, []int) {
	A = ra(A)
	B = rb(B)
	return A, B
}

func rra(A []int) []int {
	if len(A) >= 2 {
		A = append([]int{A[len(A)-1]}, A[:len(A)-1]...)
	}
	return A
}

func rrb(B []int) []int {
	if len(B) >= 2 {
		B = append([]int{B[len(B)-1]}, B[:len(B)-1]...)
	}
	return B
}

func rrr(A []int, B []int) ([]int, []int) {
	A = rra(A)
	B = rrb(B)
	return A, B
}

func push_swap(A []int, B []int, gestes []string, n int) ([]int, []int, []string) {
	if n <= 1 {
		return A, B, gestes
	}
	h := n / 2

	A, B, gestes = push_swap(A, B, gestes, h)
	for i := 0; i < h; i++ {
		A, B = pb(A, B)
		gestes = append(gestes, "pb\n")
	}

	A, B, gestes = push_swap(A, B, gestes, n-h)
	for i := 0; i < n-h; i++ {
		A = ra(A)
		gestes = append(gestes, "ra\n")
	}

	i, j := n-h, h
	for i > 0 || j > 0 {
		if j == 0 || (i > 0 && A[len(A)-1] > B[0]) {
			A = rra(A)
			gestes = append(gestes, "rra\n")
			i--
		} else {
			A, B = pa(A, B)
			gestes = append(gestes, "pa\n")
			j--
		}
	}
	return A, B, gestes
}

func read_input() []int {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("Error: Invalid number of arguments")
		return nil
	}
	args := strings.Fields(os.Args[1])
	if len(args) == 0 {
		fmt.Println("Error: No input provided")
		return nil
	}
	A := make([]int, len(args))
	for i, v := range args {
		A[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error: A or B contains non-integer values")
			return nil
		}
	}
	return A
}

func main() {
	A := read_input()
	if A == nil {
		return
	}
	B := []int{}
	gestes := []string{}
	A, B, gestes = push_swap(A, B, gestes, len(A))
	fmt.Println(A, B, gestes, len(gestes))
}
