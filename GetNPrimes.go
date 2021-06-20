package gtoolkit

import (
	"math"
)

// return all prime numbers between 2 and n
func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}

        isNotPrime := make([]bool, n+1)
        prime := []int{}
        pivot := int(math.Ceil(math.Sqrt(float64(n))))

        for i:=2; i<n+1; i++ {
                if isNotPrime[i] == false {
                        prime = append(prime, i)
                        if i<=pivot {
                                for j:=i*i; j<=n; j+=i {
                                        isNotPrime[j] = true
                                }
                        }
                }
        }

        return prime
}
