package main

import "fmt"

//二维数组查找
func Find(matrix [4][4]int, columns int, rows int, number int) bool {
	found := false
	if len(matrix) > 0 && columns > 0 && rows > 0 {
		row := 0
		column := columns - 1
		for row < rows && column >= 0 {
			if matrix[row][column] == number {
				found = true
				break
			} else if matrix[row][column] > number {
				column--
			} else {
				row++
			}
		}
	}
	return found
}

func main() {
	matrix := [4][4]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(Find(matrix, 4, 4, 20))
}
