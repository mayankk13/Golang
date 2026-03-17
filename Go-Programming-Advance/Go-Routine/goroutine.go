// Goroutines are just functions that leave the main thread and run in the backgroubd and come back to join
// the main thread once the functions are finished t oreturn any value
// Goroutine does not stop program flow and are non blocking

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("hello from Goroutine!")
}

func main() {
	fmt.Println("begining program.")
	go sayHello()
	fmt.Println("After sayHello function")
	time.Sleep(2 * time.Second)
}

/*
when the go runtime comes across go say hello. It sees that there is a function with a preceding go.
It doesn't care about anything. It immediately extracts this function out of the main thread. So within a
fraction of milliseconds this is extracted away.

It's the go runtime that handles go routines.

And the execution moves on to the next line.
And the next line is immediately printed which says after say hello function.
Then it moves on to the next line which is blank.
So it moves to line which says time dot sleep and the duration.

So now the program execution blocks for two seconds.
And within these two seconds for one second say hello waits and then prints out hello from go routine.

So after this this function is done.
So then it comes to the main thread and then this print statement gets executed.

*/
