package problem0034

// func countPrime(number int) int {
// 	var count int
// 	for i := 1; i < number; i++ {
// 		if isPrime(i) {
// 			count++
// 		}
// 	}
// 	return count
// }

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrime(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	// Loop's ending condition is i * i < n
	// to avoid expensive function of sqrt
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}
	}

	var count int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
