package main

func receive(over chan<- bool) {
	//<-over
	over <- true
}

func main() {
	ch := make(chan bool)
	go receive(ch)
	<-ch
}
