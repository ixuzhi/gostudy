package main

import "fmt"
import "os"
import "os/exec"
import "syscall"

func main() {
	fmt.Println("vim-go")
	binary ,lookerr := exec.LookPath("ls")
	if lookerr!=nil {
		panic(lookerr)
	} else {
		fmt.Println(binary)
	}


	args := []string{"ls","-a","-l","-h"}
	env := os.Environ()
	fmt.Println(env)	
	execerr := syscall.Exec(binary,args,env)
	if execerr!=nil {
		panic(execerr)
	}
}
