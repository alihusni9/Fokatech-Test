package main

import "fmt"

func main() {
	dataShiftArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dataShift := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	//rotateArray(3,3, dataShift)
	for i:=0; i<3; i++ {
		for j:=0; j<len(dataShiftArray); j++ {
			if j < 3 {
				dataShift[i][j]=dataShiftArray[j]
				dataShiftArray = dataShiftArray[:j]
			}
		}
	}
	rotateArray(3,3, dataShift)
	fmt.Println("dataShift : ",dataShift)
}

//func ShiftArray(array []int, x int) []int {
//	var resultList []int
//	dataShift := [][]int{
//		{1,2,3},
//		{4,5,6},
//		{7,8,9},
//	}
//	for i:=0; i<3; i++ {
//		for j:=0; j<len(array); j++ {
//			if j < 3 {
//				dataShift[i][j]=array[j]
//				array = array[:j]
//			}
//		}
//	}
//	rotateArray(3,3, dataShift)
//
//	return array
//}

func rotateArray(m int, n int, mat[][]int) [][]int {
	row := 0
	col := 0
	var prev, curr int
	for row < m && col < n {
		if row+1==m || col+1==n {
			break;
		}
		prev = mat[row+1][col]
		for i:=col; i<n; i++ {
			curr = mat[row][i]
			mat[row][i]=prev
			prev = curr
		}
		row++
		for i:=row; i<m; i++ {
			curr = mat[i][n-1]
			mat[i][n-1]=prev
			prev = curr
		}
		n--
		if col < n {
			for i:=m-1; i>=row; i-- {
				curr = mat[i][col]
				mat[i][col] = prev
				prev = curr
			}
		}
		col++
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			fmt.Println("array : ",mat[i][j])
		}
	}
	return nil
}