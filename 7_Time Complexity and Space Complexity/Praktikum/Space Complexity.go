func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		temp := pow(x, n/2)
		return temp * temp
	}

	temp := pow(x, (n-1)/2)
	return x * temp * temp
}

func main() {
	fmt.Printf(pow(2, 3))  // 8
	fmt.Printf(pow(5, 3))  // 125
	fmt.Printf(pow(10, 2)) // 100
	fmt.Printf(pow(2, 5))  // 32
	fmt.Printf(pow(7, 3))  // 343
}
