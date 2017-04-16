/*
#Name: Zachary Romick
#Email: zachary.romick@vanderbilt.edu
#VUnet ID: romickz
#Class: CS3270
#Date: 4/16/17
#Honor Statement: I pledge on my honor that I have neither
#given nor received unauthorized aid on this assignment.

#Name: Arunabh Singh
#Email: arunabh.singh@vanderbilt.edu
#VUnet ID: singha4
#Class: CS3270
#Date: 4/16/17
#Honor Statement: I pledge on my honor that I have neither
#given nor received unauthorized aid on this assignment.
*/

package main

import (
    "bufio"
    "fmt"
    "strconv"
    "os"
)

var board = make([][]int, 9)

func readBoard(name string) () {
	f, err := os.Open(name)	
	check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	i := 0
	for scanner.Scan() {
		res, err := strconv.Atoi(scanner.Text())
		check(err)
	    board[i] = append(board[i], res)
	    if len(board[i]) == 9 {
	    	i = i + 1
	    }
	    if i > 8 {
	    	break
	    }
	}	

	fmt.Print(board)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	readBoard("sudoku-sample.txt")
}