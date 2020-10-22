package main

func Burbuja(s []int64) {
	t := len(s) - 1
	for i := 0; i < t; i++ {
		for j := i + 1; j <= t; j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func main() {
	Burbuja([]int64{5, 2, 9, 1, 0})
}
