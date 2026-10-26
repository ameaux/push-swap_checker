package main

func pa() {
	if len(B) > 0{
		A = append([]int{B[0]}, A...)
		B = B[1:]
	}

}

func pb() {
	if len(A) > 0{
		B = append([]int{A[0]}, B...)
		A = A[1:]
	}

}

func sa() {
	if len(A) >= 2{
		A[0], A[1] = A[1], A[0]
	}
}