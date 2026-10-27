package main

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

func push_swap(A []int, gestes []string) []int {
	if len(A) <= 1 {
		return A
	}
	if len(A) == 2 && A[0] > A[1] {
		A = sa(A)
		gestes = append(gestes, "sa\n")
		return append(push_swap(A[len(A)/2:], gestes), push_swap(A[:len(A)/2], gestes)...)
	}
	return append(push_swap(A[len(A)/2:], gestes), push_swap(A[:len(A)/2], gestes)...)

}
