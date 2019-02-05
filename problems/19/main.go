//https://projecteuler.net/problem=19
package main

func a() (count int) {
	day := 1 // https://en.wikipedia.org/wiki/January_1901#January_1,_1901_(Tuesday)
	for year := 1901; year <= 2000; year++ {
		for month := 1; month < 13; month++ {
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				day += 31
			case 4, 6, 9, 11:
				day += 30
			case 2:
				if year%4 == 0 {
					if year%100 == 0 {
						if year%400 == 0 {
							day += 29
						}
					} else {
						day += 29
					}
				}
			}
			if day%7 == 6 {
				count++
			}
		}
	}
	return
}

func main() {
	println("a:", a())
}

// Solution is 171
