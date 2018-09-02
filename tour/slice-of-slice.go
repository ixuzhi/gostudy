package main

import(
	"fmt"
	"strings"
)
//切片可包含任何类型，甚至包括其它的切片
func main(){
	board := [][]string {
		[]string{"-","-",""},
		[]string{"-","-",""},
		[]string{"-","-",""},
	}
	board[0][0] = "x"
	board[2][2] = "o"
	board[1][2] = "x"

	for i := 0;i <len(board);i++ {
		fmt.Printf("%s\n",strings.Join(board[i],""))
	}
}