package main

import (
	"fmt"
	"time"
)

func signedOverflow() {
	stop := time.After(5 * time.Second)
	for i := 9223372036854775807; ; i++ {
		fmt.Printf("\r%d", i)
		select {
		case <-stop:
			return
		default:
		}
	}
}

func unsignedOverflow() {
	stop := time.After(5 * time.Second)
	for i := uint(18446744073709551615); ; i++ {
		fmt.Printf("\r%d", i)
		select {
		case <-stop:
			return
		default:
		}
	}

}

func main() {
	unsignedOverflow()
}
