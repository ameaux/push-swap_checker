package main

import "fmt"

var A []int
var B []int
var gestes []string

func pa() {
	if len(B) > 0 {
		A = append([]int{B[0]}, A...)
		B = B[1:]
	}

}

func pb() {
	if len(A) > 0 {
		B = append([]int{A[0]}, B...)
		A = A[1:]
	}

}

func sa() {
	if len(A) >= 2 {
		A[0], A[1] = A[1], A[0]
	}
}

func sb() {
	if len(B) >= 2 {
		B[0], B[1] = B[1], B[0]
	}
}

func ss() {
	sa()
	sb()
}

func ra() {
	if len(A) >= 2 {
		A = append(A[1:], A[0])
	}
}

func rb() {
	if len(B) >= 2 {
		B = append(B[1:], B[0])
	}
}

func rr() {
	ra()
	rb()
}

func rra() {
	if len(A) >= 2 {
		A = append([]int{A[len(A)-1]}, A[:len(A)-1]...)
	}
}

func rrb() {
	if len(B) >= 2 {
		B = append([]int{B[len(B)-1]}, B[:len(B)-1]...)
	}
}

func rrr() {
	rra()
	rrb()
}

func push_swap(A []int, B []int, gestes []string) ([]int, []int, []string) {
	if len(A) <= 1 {
		return A, B, gestes
	}
	if len(A) == 2 && A[0] > A[1] {
		sa()
		gestes = append(gestes, "sa")
	}
	return A, B, gestes
}

func main() {
	A := []int{2, 1}
	var B []int
	gestes := []string{}
	A, B, gestes = push_swap(A, B, gestes)
	fmt.Println(A, B, gestes)
}
