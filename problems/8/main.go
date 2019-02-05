package main

import (
	"fmt"
	"strconv"
	"testing"
)

var dd []int

func init() {
	s := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	for i := range s {
		d, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			panic(err)
		}
		dd = append(dd, d)
	}
	if len(dd) != len(s) {
		panic("len(dd) != len(s)")
	}
}

func a(n int) (str []int, result int) {
	sizeOfDigits := len(dd)
	if n > sizeOfDigits {
		return
	}
	best := 0
	bestStr := []int{}
	for i, j := 0, n; i < sizeOfDigits && j < sizeOfDigits; {
		str := dd[i:j]
		local := 1
		for i := range str {
			local *= str[i]
		}
		if local > best {
			best = local
			bestStr = str
		}
		i++
		j++
	}
	return bestStr, best
}

func main() {
	{
		if _, y := a(4); y != 5832 {
			panic("_, y := a(4); y != 5832")
		}
	}
	x, y := a(10) // [8 8 3 9 9 8 7 9 7 9] 493807104
	fmt.Println("a:", x, y)
	a := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a(10)
		}
	})
	println("a: ", a.String(), a.MemString())
}
