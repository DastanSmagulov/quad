package piscine

func QuadB(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 || i == y-1 && j == x-1 && x-1 != 0 && y-1 != 0 {
				print("/")
			} else if i == y-1 && j == 0 || i == 0 && j == x-1 {
				print("\\")
			} else if j == 0 || j == x-1 || i == 0 || i == y-1 {
				print("*")
			} else {
				print(" ")
			}
		}
		print("\n")
	}
}