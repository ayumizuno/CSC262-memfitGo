package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	name string
	size int
	offset int
}

//is_adjacent method for block struct

type Simulation struct {
	algorithm string
	size int
	freeList []Block
	usedList []Block
	lastOffset int //default value 0
	failed int //default value 0
}

func start(line string) Simulation {
	//Returns a new simulation object
	poolLine := strings.Fields(line)
	size, _ := strconv.Atoi(poolLine[2])
	sim := Simulation{algorithm: poolLine[1], size: size}
	free := make([]Block, 1)
	free[0] = Block{"pool", size, 0}
	sim.freeList = free
	return sim
}

func main() {
	fmt.Println("hello")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sim := start(scanner.Text())
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fmt.Println(line)
			line := strings.Fields(line)
			if line[0] == "alloc" {
				//call alloc
			} else if line[0] == "free" {
				//call free
				//compact free list
			} else {
				//raise error
			}
			//print both free list and used list
		}
	}
	//print stats
}

