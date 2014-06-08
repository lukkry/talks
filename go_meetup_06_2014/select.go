select {
case v1 := <-c1
	fmt.Printf("Received from c1")
case c2 <- true
	fmt.Printf("sent to c2")
default:
	fmt.Printf("no one was ready to communicate")
}
