// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"sync"
)

/*
What happens if you remove the go-command from the Seek call in the main function?
If the go-command is removed the program will execute sequentially and not parallel. A new thread wont be created.

What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
You would pass the value of the wait group instead of the pointer to the wait group. This would mean that the wait group would be copied and the wait group in the main function would not be updated.

What happens if you remove the buffer on the channel match?
It would result in the channel blocking. The channel would block until a receiver is ready to receive the message. This would result in a deadlock.

What happens if you remove the default-case from the case-statement in the main function?
If the default case is removed the program will block until a message is received. This would result inthe program hanging and not completing.

*/

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}
