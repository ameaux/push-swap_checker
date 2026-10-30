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

func petit(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func grand(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inv(k, n int) int {
	if k == 0 {
		return 0
	}
	return n - k
}

func plan(i, j, nA, nB int) (int, int) {
	cout, mode := grand(i, j), 0
	if c := grand(inv(i, nA), inv(j, nB)); c < cout {
		cout, mode = c, 1
	}
	if c := i + inv(j, nB); c < cout {
		cout, mode = c, 2
	}
	if c := inv(i, nA) + j; c < cout {
		cout, mode = c, 3
	}
	return cout, mode
}

func jouer(A, B []int, gestes []string, i, j, mode int) ([]int, []int, []string) {
	nA, nB := len(A), len(B)
	var nrr, nrrr, nra, nrb, nrra, nrrb int
	switch mode {
	case 0:
		nrr = petit(i, j)
		nra, nrb = i-nrr, j-nrr
	case 1:
		ia, jb := inv(i, nA), inv(j, nB)
		nrrr = petit(ia, jb)
		nrra, nrrb = ia-nrrr, jb-nrrr
	case 2:
		nra, nrrb = i, inv(j, nB)
	case 3:
		nrra, nrb = inv(i, nA), j
	}
	for k := 0; k < nrr; k++ {
		A, B = rr(A, B)
		gestes = append(gestes, "rr\n")
	}
	for k := 0; k < nrrr; k++ {
		A, B = rrr(A, B)
		gestes = append(gestes, "rrr\n")
	}
	for k := 0; k < nra; k++ {
		A = ra(A)
		gestes = append(gestes, "ra\n")
	}
	for k := 0; k < nrb; k++ {
		B = rb(B)
		gestes = append(gestes, "rb\n")
	}
	for k := 0; k < nrra; k++ {
		A = rra(A)
		gestes = append(gestes, "rra\n")
	}
	for k := 0; k < nrrb; k++ {
		B = rrb(B)
		gestes = append(gestes, "rrb\n")
	}
	return A, B, gestes
}

func cibleB(B []int, x int) int {
	j := -1
	for k := range B {
		if B[k] < x && (j == -1 || B[k] > B[j]) {
			j = k
		}
	}
	if j == -1 {
		for k := range B {
			if j == -1 || B[k] > B[j] {
				j = k
			}
		}
	}
	return j
}

func cibleA(A []int, x int) int {
	i := -1
	for k := range A {
		if A[k] > x && (i == -1 || A[k] < A[i]) {
			i = k
		}
	}
	if i == -1 {
		for k := range A {
			if i == -1 || A[k] < A[i] {
				i = k
			}
		}
	}
	return i
}

func trier(A []int, B []int, gestes []string) ([]int, []int, []string) {
	// rangs : 0 = la plus petite carte
	rangs := make([]int, len(A))
	for i := range A {
		for j := range A {
			if A[j] < A[i] {
				rangs[i]++
			}
		}
	}
	A = rangs

	for k := 0; k < 2; k++ {
		if len(A) > 3 {
			A, B = pb(A, B)
			gestes = append(gestes, "pb\n")
		}
	}
	for len(A) > 3 {
		meilleur, bi, bj, bm := -1, 0, 0, 0
		for i := range A {
			j := cibleB(B, A[i])
			c, m := plan(i, j, len(A), len(B))
			if meilleur == -1 || c < meilleur {
				meilleur, bi, bj, bm = c, i, j, m
			}
		}
		A, B, gestes = jouer(A, B, gestes, bi, bj, bm)
		A, B = pb(A, B)
		gestes = append(gestes, "pb\n")
	}

	if len(A) == 3 {
		m := 0
		for i := range A {
			if A[i] > A[m] {
				m = i
			}
		}
		if m == 0 {
			A = ra(A)
			gestes = append(gestes, "ra\n")
		} else if m == 1 {
			A = rra(A)
			gestes = append(gestes, "rra\n")
		}
	}
	if len(A) >= 2 && A[0] > A[1] {
		A = sa(A)
		gestes = append(gestes, "sa\n")
	}

	for len(B) > 0 {
		meilleur, bi, bj, bm := -1, 0, 0, 0
		for j := range B {
			i := cibleA(A, B[j])
			c, m := plan(i, j, len(A), len(B))
			if meilleur == -1 || c < meilleur {
				meilleur, bi, bj, bm = c, i, j, m
			}
		}
		A, B, gestes = jouer(A, B, gestes, bi, bj, bm)
		A, B = pa(A, B)
		gestes = append(gestes, "pa\n")
	}

	p := 0
	for i := range A {
		if A[i] == 0 {
			p = i
		}
	}
	if p <= len(A)-p {
		for k := 0; k < p; k++ {
			A = ra(A)
			gestes = append(gestes, "ra\n")
		}
	} else {
		for k := 0; k < len(A)-p; k++ {
			A = rra(A)
			gestes = append(gestes, "rra\n")
		}
	}
	return A, B, gestes
}

func estRangee(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func main() {
	A := read_input()
	if A == nil {
		return
	}
	if estRangee(A) {
		return
	}
	gestes := []string{}
	A, _, gestes = trier(A, []int{}, gestes)
	for _, g := range gestes {
		fmt.Print(g)
	}
	fmt.Print(len(gestes), " ")
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
		for j := 0; j < i; j++ {
			if A[i] == A[j] {
				fmt.Println("Error: A or B contains duplicate values")
				return nil
			}
		}
	}
	return A
}
