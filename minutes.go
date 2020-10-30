package main

import (
	"fmt"
	"time"
)

func main() {
	m, _ := time.ParseDuration("1h40m")
	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
}
