package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	BeginningN := n

	var pfn []int

	for n%2 == 0 {
		pfn = append(pfn, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfn = append(pfn, i)
			n = n / i
		}
	}

	if n > 2 {
		pfn = append(pfn, n)
	}

	fmt.Printf("The prime factors of %d are: %v\n", BeginningN, pfn)
}
