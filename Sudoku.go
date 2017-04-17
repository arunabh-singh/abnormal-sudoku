/*
 * Name: Zachary Romick
 * Email: zachary.romick@vanderbilt.edu
 * VUnet ID: romickz
 * Honor Statement: I pledge on my honor that I have neither
 * given nor received unauthorized aid on this assignment.
 
 * Name: Arunabh Singh
 * Email: arunabh.singh@vanderbilt.edu
 * VUnet ID: singha4
 * Honor Statement: I pledge on my honor that I have neither
 * given nor received unauthorized aid on this assignment.

 * Class: CS3270
 * Date: 4/16/17
 */

/*
 * Description: This is a Go program that tries to solve a given Sudoku board.
 */

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"    
    "time"
)

// global variables
var sudokuBoard = make([][]int, 9)	// represents board

// main()
// Does all the work, calls all necessary functions for solving
func main() {
	readBoard("sudoku-sample.txt")
	fmt.Print("\n Here is the given puzzle:\n ")
	printBoard()
	
	fmt.Print("\n\n Here is the solution:\n ")

	start := time.Now()
	solveable := solve(0, 0)

	if !solveable {
		fmt.Print("No solution.\n")
	} else {
		execution := time.Since(start)
		printBoard()

		fmt.Print("\nThis puzzle was solved in ", execution, "\n")
	}
}

// check(error)
// ensures that no errors occur in reading a file
// @param e 	error code to compare
func checkError(e error) {
	if e != nil {
        panic(e)
    }
}

// readBoard(string)
// reads a file into a sudoku board slice
// @param name 	name of the file to read from
func readBoard(name string) () {
	f, err := os.Open(name)	
	checkError(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	i := 0
	for scanner.Scan() {
		res, err := strconv.Atoi(scanner.Text())
		checkError(err)
	    sudokuBoard[i] = append(sudokuBoard[i], res)
	    if len(sudokuBoard[i]) == 9 {
	    	i = i + 1
	    }
	    if i > 8 {
	    	break
	    }
	}	
}

// printSudokuBoard()
// will print the globally accessible board
func printBoard() {
	for i := 0; i != 9; i++ {
		if i % 3 == 0&& i != 0 {
			fmt.Print("------+-------+------\n ")
		}
		for j := 0; j != 9; j++ {
			if j % 3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Print(sudokuBoard[i][j], " ")
		}
		fmt.Print("\n ")
	}
}

// checkCol(col,num)
// checks whether a number already exists within a column
// @param col 	column to check
// @param num 	number to check
// @return 	whether a number is not in the column
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
// @param row 	row to check
// @param num 	number to check
// @return 	whether a number is not in the row
func checkRow(row int, num int) bool {
	for i := 0; i != 9; i++ {
		if sudokuBoard[row][i] == num {
			return false
		}
	}
	return true
}

// checkSubgrid(row,col,num)
// checks whether a number already exists within a subgrid
// @param row 	row to check
// @param col 	column to check
// @param num 	number to check
// @return 	whether a number is not in the subgrid
func checkSubgrid(row int, col int, num int) bool {
	var gR int = gridCorner(row) //gridRow
	var gC int = gridCorner(col) //gridCol
	var gridIsGood = true
	if sudokuBoard[gR][gC] == num || sudokuBoard[gR][gC+1] == num || sudokuBoard[gR][gC+2] == num {
		gridIsGood = false
	}
	if sudokuBoard[gR+1][gC] == num	|| sudokuBoard[gR+1][gC+1] == num || sudokuBoard[gR+1][gC+2] == num {
		gridIsGood = false
	}
	if sudokuBoard[gR+2][gC] == num || sudokuBoard[gR+2][gC+1] == num || sudokuBoard[gR+2][gC+2] == num {
		gridIsGood = false
	}
	return gridIsGood
}
	

// gridCorner(roworcol)
// returns the leftmost corner of a grid; acts as a helper for checkSubgrid
// @param roworcol 	position to check
// @return 	a unique value assignd to this particular subgrid 
func gridCorner(roworcol int) int {
	var returnVal int = (((roworcol/3)+1)*3)-3
	return returnVal
}

// check(row,col,num)
// performs all checks to see if a number can be placed
// @param row 	row to check
// @param col 	column to check
// @param num 	number to check
// @return 	whether a number is placeable at the described position
func check(row int, col int, num int) bool {
	return (checkCol(col,num) && checkRow(row,num) && checkSubgrid(row,col,num))
}

// solve(row,col)
// attempts to solve a board using recursion
// @param row 	row to begin check
// @param col 	column to begin check
// @return 	whether the puzzle can be solved with given constraints
func solve(row int, col int) bool {
	if row == 9 {
		return true
	} else if col == 9 {
		return solve(row+1,0)
	} else if sudokuBoard[row][col] > 0 {
		return solve(row,col+1)
	} else {
		for i := 1; i != 10; i++ { 
			if check(row,col,i) {
				sudokuBoard[row][col] = i
				if solve(row,col+1){
					return true
				}
			}
		}
		sudokuBoard[row][col] = 0
		return false
	}
}