package gtoolkit

func GetLowercasePrimeDict() map[int]int {
        m := make(map[int]int)
        primes := []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}

        for i:=0; i<26; i++ {
                m[i+97] = primes[i]
        } 

        return m
}
