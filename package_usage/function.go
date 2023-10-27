package packageusage

// 1.  buatlah sebuah fungsi getFactorial(num int) yang akan mengembalikan hasil dari faktorial dari num (num!)

func getFactorial(num int) int {
	result := 1
	if num == 0 {
		return result
	}

	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
