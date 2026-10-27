package main

func main(A []int) ([]int, int) {
	gestes := []string{}
	return push_swap(A, gestes), len(gestes)
}
