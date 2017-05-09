package main

import "fmt"
import "os/exec"
import "io/ioutil"

func main() {
	fmt.Println("vim-go")
	dateCmd := exec.Command("date")
	fmt.Println(dateCmd)
	
	dateOut,err := dateCmd.Output()
	if err!=nil {
		panic(err)
		//fmt.Println(string(dateOut))
	} else {
		fmt.Println(string(dateOut))
		fmt.Println("dateOut")
	}

	grepCmd := exec.Command("grep","hello")
	grepIn,_ := grepCmd.StdinPipe()
	grepOut,_:= grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()
	
	fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
		panic(err)
	}
	
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
