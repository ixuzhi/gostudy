package main

import(
	"fmt"
	"io"
	"strings"
)
/*
io.Reader 接口有一个 Read 方法:
func (T) Read(b []byte) (n int, err error)
*/
func main(){
	r := strings.NewReader("hello reader!")
	b := make([]byte,8)
	for {
		n,err := r.Read(b)
		fmt.Printf("%v,%v",n,err)
		fmt.Printf("b[:n]=%q\n",b[:n])
		if err == io.EOF{
			break
		}

	}
}