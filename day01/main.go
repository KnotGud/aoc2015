package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputPath = "./input"

func main() {
	f, err := os.OpenFile(inputPath, os.O_RDONLY, 0644)
	if err != nil {
		log.Panic(err)
	}

	s := bufio.NewScanner(f)
	c := make(chan string, 10)
	go func() {
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()

	part2(c)
}

func part1(in chan string) {
	for str := range in {
		floor := 0
		for _, c := range str {
			switch c {
			case '(':
				floor++
			case ')':
				floor--
			}
		}
		fmt.Println("Floors:", floor)
	}
}

func part2(in chan string) {
	for str := range in {
		floor := 0
		basementIndex := 0
		for i, c := range str {
			switch c {
			case '(':
				floor++
			case ')':
				floor--
				if (floor == -1) && (basementIndex == 0) {
					basementIndex = i
				}
			}
		}
		fmt.Println("Floors:", floor)
		fmt.Println("First basement instruction:", basementIndex+1) // account for index starting at 0, instructions start at 1
	}
}
