// Create a new channel
ch := make(chan bool)

// Send v to channel ch.
// The "arrow" indicates the direction of data flow.
ch <- true

// Receive from ch, and assign value to v.
v := <-ch
