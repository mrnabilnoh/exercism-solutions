package diffsquares

func SquareOfSum(n int) int {
	// Iteration 1: slow loop
	// var sum int
	// for i := n; i > 0; i-- {
	// 	sum += i
	// }
	// return sum * sum

	// Iteration 2: math formula: Sum of natural numbers
	// reference: https://byjus.com/maths/sum-of-n-terms/
	sum := n * (n + 1) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	// Iteration 1: slow loop
	// var sum int
	// for i := n; i > 0; i-- {
	// 	sum += (i * i)
	// }
	// return sum

	// Iteration 2: math formula: Sum of square of ‘n’ natural numbers
	// reference: https://byjus.com/maths/sum-of-n-terms/
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
