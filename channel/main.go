package main

import (
	"fmt"
)

func selectRand() {
	chanCap := 5
	c := make(chan int, chanCap)

	for i := 0; i < chanCap; i++ {
		select {
		case c <- 1:
		case c <- 2:
		case c <- 3:
		case c <- 4:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("value %v\n", <-c)
	}
}

const (
	RepoGit = iota
	RepoSvn
)

func main() {
	fmt.Println("", RepoGit, " ", RepoSvn)
	selectRand()
	fmt.Println("chen")
}
