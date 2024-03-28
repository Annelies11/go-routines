package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestGoRoutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Hehehehe")
	time.Sleep(1 * time.Second)
}
