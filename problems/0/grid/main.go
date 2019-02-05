package main

func a(grid [][]int) (x, y, xp, yp, best int) {
	for sx := 0; sx < len(grid[0]); sx++ { // start x
		for sy := 0; sy < len(grid); sy++ { // start y
			for ex := sx; ex < len(grid[0]); ex++ { // end x
				for ey := sy; ey < len(grid); ey++ { // end y
					result := 0
					ints := []int{}
					for mx := sx; mx <= ex; mx++ { // moving x
						for my := sy; my <= ey; my++ { // moving y
							n := grid[my][mx]      // case
							result += n            // add
							ints = append(ints, n) // count
						}
					}
					if best < len(ints) {
						best = len(ints)
						x, y, xp, yp = sx, sy, ex, ey
					}
				}
			}
		}
	}
	return
}

func main() {
	grid := [][]int{
		[]int{1, 2, 3},
		[]int{4, 0, -4},
		[]int{-3, -2, -1},
		[]int{7, 7, 7},
	}
	a(grid)
}
