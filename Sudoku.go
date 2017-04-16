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
    "fmt"
    "io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("sudoku-sample.txt")
	if err != nil { 
		panic(err) 
	}
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")	
	if err != nil { 
		panic(err) 
	}

	var board [9][9]int
	
	
}