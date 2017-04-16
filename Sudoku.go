/*
#Name: Zachary Romick
#Email: zachary.romick@vanderbilt.edu
#VUnet ID: romickz
#Honor Statement: I pledge on my honor that I have neither
#given nor received unauthorized aid on this assignment.

#Name: Arunabh Singh
#Email: arunabh.singh@vanderbilt.edu
#VUnet ID: singha4
#Honor Statement: I pledge on my honor that I have neither
#given nor received unauthorized aid on this assignment.

#Class: CS3270
#Date: 4/16/17
*/

/*
Description: This is a Go program that tries to solve a given Sudoku board.
*/


package main

import (
    "bufio"
    "fmt"
    //"io"
    "strconv"
    //"strings"
    "os"
)

// global variables
sudokuBoard := make([]int ([]int, 9), 9)

// main()
// Does all the work, calls all necessary functions to solve
func main() {
	f, err := os.Open("sudoku-sample.txt")	
	checkError(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	
	i := 0
	for scanner.Scan() {
		res, err := strconv.Atoi(scanner.Text())
		checkError(err)
	    sudokuBoard[i] = append(sudokuBoard[i], res)
	    if len(sudokuBoard[i]) == 3 || len(sudokuBoard[i]) == 6 {
	    	fmt.Print("------+-------+------\n")
	    }
	    if len(sudokuBoard[i]) == 9 {
	    	i = i + 1
	    }
	    if i > 8 {
	    	break
	    }
	}	
	fmt.Print(sudokuBoard)
}

// check(error)
// ensures that no errors occur in main
func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

// printSudokuBoard will print the globally accessible board
func printBoard(board [9][9]int) {

}

// checkCol(col,num)
// checks whether a number already exists within a column
func checkCol(col int, num int) bool {
	for i := 0; i != 9; i++ {
		if sudokuBoard[i][col] == num {
			return false
		}
	}
	return true
}

// checkRow(row,num)
// checks whether a number already exists within a row
func checkRow(row int, num int) bool {
	//var disjointRow []int = sudokuBoard[row] - [num]
	//return (disjointRow.length == 9)
	return true;
}

// checkSubgrid(row,col,num)
// checks whether a number already exists within a subgrid
func checkSubgrid(row int, col int, num int) bool {
	var gR int = gridCorner(row) //gridRow
	var gC int = gridCorner(col) //gridCol
	//var grid [9]int = sudokuBoard[gR][gC..(gC+2)] + sudokuBoard[gR+1][gC..(gC+2)] + sudokuBoard[gR+2][gC..(gC+2)]
	//var disjointRow []int = grid[0..8] - [num]
	//return (disjointRow.length == 9)
	return true;
}
	

// gridCorner(roworcol)
// returns the leftmost corner of a grid
// acts as a helper for checkSubgrid
func gridCorner(roworcol int) int {
	var returnVal int = (((roworcol/3)+1)*3)-3
	return returnVal
}

// check(row,col,num)
// performs all checks to see if a number can be placed
func check(row int, col int, num int) bool {
	return (checkCol(col,num) && checkRow(row,num) && checkSubgrid(row,col,num))
}

// solve(row,col)
// attempts to solve a board using recursion
// returns true if solvable, else false
func solve(row int, col int) bool {
	/*if (row == 9) {
		return true
	}
	else if (col == 9) {
		return solve(row+1,0)
	}
	else if (sudokuBoard[row][col] > 0) {
		return solve(row,col+1)
	}
	else {
		for i := 1; i != 10; i++ { 
			if (check(row,col,i)){
				sudokuBoard[row][col] = i
				if (solve(row,col+1)){
					return true
				}
			}
		}
		sudokuBoard[row][col] = 0
		return false
	}		*/
	return true
}