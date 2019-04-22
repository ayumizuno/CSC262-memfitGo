package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	name string
	size int
	offset int
}

func (block Block) isAdjacent(currBlock Block) bool {
	/* Returns true if block is next to another block */
	if currBlock.offset == block.offset + block.size ||
			block.offset == currBlock.offset + block.size {
		return true
	}
	return false
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
	fmt.Println("Free List")
	//sort free list by offset
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].offset < sim.freeList[j].offset
	})
	for _, block := range sim.freeList {
		fmt.Println(
			"\t offset: ", block.offset,
			"\t size:", block.size)

	}
	fmt.Println("Used List")
	//sort used list by offset
	//loop through and print each block
	for _, block := range sim.usedList {
		fmt.Println(
			"\t offset:", block.offset,
			"\t size:", block.size,
			"\t name:", block.name)
	}
}

func (sim Simulation) getStats() (float32, float32) {
	/*
	Returns pct of free and used space
	in pool
	 */
	free := 0
	used := 0
	// add block size to free for all blocks in free list
	for _, block := range sim.freeList {
		free += block.size
	}
	//add block size to used for all blocks in used list
	for _, block := range sim.usedList {
		used += block.size
	}
	pctFree := float32((free/sim.size) * 100)
	pctUsed := float32((used/sim.size) * 100)
	return pctFree, pctUsed
}

func (sim Simulation) alloc(name string, size int) {
	/* Splits an available block based on input size */
	var block *Block
	if sim.algorithm == "first" {
		block = sim.allocFirst(size)
	} else if sim.algorithm == "best" {
		block = sim.allocBest(size)
	} else if sim.algorithm == "worst" {
		block = sim.allocWorst(size)
	} else if sim.algorithm == "next" {
		block = sim.allocNext(size)
	} else if sim.algorithm == "random" {
		block = sim.allocRandom(size)
	} else {
		fmt.Println("raise some error")
	}
	if block != nil {
		sim.blockSplit(block, name, size)
	}
}

func (sim Simulation) find(size int) *Block {
	/*
	Return a block in the free list that is
	equal to or greater than the specified size.
	If no such block, return nil block
	 */
	var nilBlock Block
	for _, block := range sim.freeList {
		if block.size >= size {
			return &block
		}
	}
	return &nilBlock
}

func (sim Simulation) findWithIndex(size int, start int, stop int) *
	Block {
	/*
	Return next available block in free list in
	range(start, stop). If no such block, return None
	 */
	var nilBlock Block
	for _, block := range sim.freeList {
		if block.offset >= start && block.offset <= stop {
			if block.size >= size {
				sim.lastOffset = block.offset
				return &block
			}
		}
	}
	return &nilBlock
}

func (sim Simulation) allocFirst(size int) *Block {
	/*
	Returns next available space based on first-fit alg
	 */
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].offset < sim.freeList[j].offset
	})
	return sim.find(size)
}

func (sim Simulation) allocBest(size int) *Block {
	/*
		Returns next available space based on best-fit alg
	*/
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].size < sim.freeList[j].size
	})
	return sim.find(size)
}

func (sim Simulation) allocWorst(size int) *Block {
	/*
		Returns next available space based on worst-fit alg
	*/
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].size > sim.freeList[j].size
	})
	return sim.find(size)
}

func (sim Simulation) allocNext(size int) *Block {
	/*
		Returns next available space based on next-fit alg
	*/
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].offset < sim.freeList[j].offset
	})
	block := sim.findWithIndex(size, sim.lastOffset, sim.size)
	if block != nil {
		return block
	} else{
		block = sim.findWithIndex(size, 0, sim.lastOffset)
		return block
	}
}

func (sim Simulation) allocRandom(size int) *Block {
	/*
		Returns random block from free list
	*/
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sim.freeList), func(i, j int){
		sim.freeList[i], sim.freeList[j] = sim.freeList[j], sim.freeList[i]
	})
	return sim.find(size)
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

func (sim Simulation) blockSplit(block *Block, newName string, size int){
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

