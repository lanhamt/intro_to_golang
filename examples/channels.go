package main
import "fmt"

func main() {
	ch := make(chan int)    // create a channel
	go worker(ch)           // launch concurrent worker function - more later
	for {                   // loop continuously
		message, is_active := <-ch  // get messages from worker
		if is_active { 
			fmt.Printf("Got a message %d\n", message) 
		} else {
			fmt.Println("All done, bye...")
			return
		}
	}
}

func worker(ch chan<- int) {
	defer close(ch)    		// never forget to close
	for i := 0; i < 5; i++ {
		ch <- i  // send messages through channel
	}
}
