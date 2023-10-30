package main

func main() {

}

func generateMatrix(n int) [][]int {
	var startX, startY int
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	count := 1
	offSet := 1
	center := n / 2
	var loop = n / 2
	for loop > 0 {
		i, j := startX, startY
		// 行不变，列在变
		for j = startY; j < n-offSet; j++ {
			res[startX][j] = count
			count++
		}

		// 列不变，行在变
		for i = startX; i < n-offSet; i++ {
			res[i][j] = count
			count++
		}

		// 行不变，列再减
		for ; j > startY; j-- {
			res[i][j] = count
			count++
		}

		// 列不变，行再减
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}

		startX++
		startY++
		offSet++
		loop--
	}
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
