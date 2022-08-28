package piscine

func QuadC(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 || i == 0 && j == x-1 {
				print("A")
			} else if i == y-1 && j == 0 || i == y-1 && j == x-1 {
				print("C")
			} else if i == 0 || j == 0 || j == x-1 || i == y-1 {
				print("B")
			} else {
				print(" ")
			}
		}
		print("\n")
	}
}