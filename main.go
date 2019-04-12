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

func (block Block) isAdjacent() bool {
	/* Returns true if block is next to another block */
}

type Simulation struct {
	algorithm string
	size int
	freeList []Block
	usedList []Block
	lastOffset int //default value 0
	failed int //default value 0
}

func (sim Simulation) printLists() {
	/* prints free list and used list */
}

func (sim Simulation) getStats() (float32, float32) {
	/*
	Returns pct of free and used space
	in pool
	 */
}

func (sim Simulation) alloc(name string, size int) {
	/* Splits an available block based on input size */
}

func (sim Simulation) find(size int) Block {
	/*
	Return a block in the free list that is
	equal to or greater than the specified size.
	If no such block, return None
	 */
}

func (sim Simulation) findWithIndex(size int, start int, stop int) Block {
	/*
	Return next available block in free list in
	range(start, stop). If no such block, return None
	 */
}

func (sim Simulation) allocFirst(size int) Block {
	/*
	Returns next available space based on first-fit alg
	 */
}

func (sim Simulation) allocBest(size int) Block {
	/*
		Returns next available space based on best-fit alg
	*/
}

func (sim Simulation) allocWorst(size int) Block {
	/*
		Returns next available space based on worst-fit alg
	*/
}

func (sim Simulation) allocNext(size int) Block {
	/*
		Returns next available space based on next-fit alg
	*/
}

func (sim Simulation) allocRandom(size int) Block {
	/*
		Returns random block from free list
	*/
}

func (sim Simulation) free(){
	/*
		Free block in used list and add it back to free list
	*/
}

func (sim Simulation) compactFree(){
	/*
		Go through free list and compact adjacent blocks
	*/
}

func (sim Simulation) blockSplit(block Block, newName string, size int){
	/*
		Construct one or two new blocks given a block and
		a specific size to allocate
	*/
}

func start(line string) Simulation {
	/* Returns a new simulation object */
	poolLine := strings.Fields(line)
	size, _ := strconv.Atoi(poolLine[2])
	sim := Simulation{algorithm: poolLine[1], size: size}
	free := make([]Block, 1)
	free[0] = Block{"pool", size, 0}
	sim.freeList = free
	return sim
}

func printStats(pctFree float32, pctUsed float32, failed int) {
	/*
		Prints stats of simulation
	 */
}
func main() {
	fmt.Println("hello")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//sim := start(scanner.Text()) //simulation object to run alg on
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
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

