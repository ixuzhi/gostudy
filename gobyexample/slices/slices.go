package main

import "fmt"

func main() {
	s:=make([]string,3)
	fmt.Println("empty slice:",s)

	fmt.Println("slice len:",len(s))
	s[0]="0"
	s[1]="1"
	s[2]="2"
	fmt.Println("slice:",s);
	
	fmt.Println("slice len:",len(s))

	s = append(s,"3")
	s=append(s,"4")
	fmt.Println("slice:",s);

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("slice c:",c)

	l := s[2:]
	fmt.Println("slice:",l)
	
	k := s[:4]
	fmt.Println("slice:",k)

	t := []string{"5","6","7"}
	fmt.Println("slice t:",t);

	twoD:=make([][]int,3)
	for i:=0;i<3;i++ {
		interLen := i+1
		twoD[i] = make([]int,interLen)
		for j:=0;j<interLen;j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("twoD:",twoD);
}
