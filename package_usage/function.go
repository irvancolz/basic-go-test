package packageusage

// 1.  buatlah sebuah fungsi getFactorial(num int) yang akan mengembalikan hasil dari faktorial dari num (num!)

func getFactorial(num int) int {
	if num == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
