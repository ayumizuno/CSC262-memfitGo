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
	/* one unit of memory allocation */
	name string
	size int
	offset int
}

func (block Block) isAdjacent(currBlock Block) bool {
	/*
	Returns true if block passed in is next to
	another block
	*/
	if currBlock.offset == block.offset + block.size ||
			block.offset == currBlock.offset + block.size {
		return true
	}
	return false
}

type Simulation struct {
	/* Properties of entire simulation */
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
	sort.Slice(sim.usedList, func(i, j int) bool {
		return sim.usedList[i].offset < sim.usedList[j].offset
	})
	//loop through and print each block
	for _, block := range sim.usedList {
		fmt.Println(
			"\t offset:", block.offset,
			"\t size:", block.size,
			"\t name:", block.name)
	}
	fmt.Println("\n")
}

func (sim *Simulation) getStats() (float64, float64) {
	/* Returns pct of free and used space */
	free := 0 //keeps track of total block size in free list
	used := 0 //keeps track of total block size in used list
	for _, block := range sim.freeList {
		free += block.size
	}
	for _, block := range sim.usedList {
		used += block.size
	}
	pctFree := (float64(free)/float64(sim.size)) * 100
	pctUsed := (float64(used)/float64(sim.size)) * 100
	return pctFree, pctUsed
}

func (sim *Simulation) alloc(name string, size int) {
	/* Splits an available block based on input size */
	block := &sim.freeList[0]
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
		err := "\nInvalid algorithm specified"
		fmt.Println(err)
		os.Exit(1)
	}
	if block != nil {
		sim.blockSplit(block, name, size)
	} else {
		sim.failed += 1
		fmt.Println("Failed allocation of ", name)
	}
}

func (sim *Simulation) find(size int) *Block {
	/*
	Return a block in the free list that is
	equal to or greater than the specified size.
	If no such block, return nil block
	 */
	for i, _ := range sim.freeList {
		if sim.freeList[i].size >= size {
			return &sim.freeList[i]
		}
	}
	return nil
}

func (sim *Simulation) findWithIndex(size int, start int, stop int) *Block {
	/*
	Return next available block in free list in
	range(start, stop). If no such block, return None
	 */
	for i, _ := range sim.freeList {
		if sim.freeList[i].offset >= start && sim.freeList[i].offset <= stop {
			if sim.freeList[i].size >= size {
				sim.lastOffset = sim.freeList[i].offset
				return &sim.freeList[i]
			}
		}
	}
	return nil
}

func (sim *Simulation) allocFirst(size int) *Block {
	/*
	Returns next available space based on first-fit alg
	 */
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].offset < sim.freeList[j].offset
	})
	blk := sim.find(size)
	return blk
}

func (sim *Simulation) allocBest(size int) *Block {
	/*
		Returns next available space based on best-fit alg
	*/
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].size < sim.freeList[j].size
	})
	blk := sim.find(size)
	return blk
}

func (sim *Simulation) allocWorst(size int) *Block {
	/*
		Returns next available space based on worst-fit alg
	*/
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].size > sim.freeList[j].size
	})
	blk := sim.find(size)
	return blk
}

func (sim *Simulation) allocNext(size int) *Block {
	/*
		Returns next available space based on next-fit alg
	*/
	fmt.Println(sim.lastOffset)
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].offset < sim.freeList[j].offset
	})
	block := sim.findWithIndex(size, sim.lastOffset, sim.size)
	if block !=nil {
		return block
	} else{
		block = sim.findWithIndex(size, 0, sim.lastOffset)
		return block
	}
}

func (sim *Simulation) allocRandom(size int) *Block {
	/*
		Returns random block from free list
	*/
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sim.freeList), func(i, j int){
		sim.freeList[i], sim.freeList[j] = sim.freeList[j], sim.freeList[i]
	})
	blk := sim.find(size)
	return blk
}

func (sim *Simulation) free(name string){
	/*
		Free block in used list and add it back to free list
	*/
	for i, _ := range sim.usedList {
		if sim.usedList[i].name == name {
			sim.freeList = append(sim.freeList, sim.usedList[i])
			sim.usedList = append(sim.usedList[:i], sim.usedList[i+1:]...) //'...' turns slice into arg
			return
		}
	}
}

func (sim *Simulation) compactFree(){
	/*
		Go through free list and compact adjacent blocks
	*/
	// sort free list by offset
	sort.Slice(sim.freeList, func(i, j int) bool {
		return sim.freeList[i].offset < sim.freeList[j].offset
	})
	copyList := make([]Block, 0)
	copyBlock := Block{name:"free", size:0, offset:0}

	for i, _ := range sim.freeList {
		if copyBlock.isAdjacent(sim.freeList[i]){
			copyBlock.size += sim.freeList[i].size
		} else {
			if copyBlock.size > 0 {
				copyList = append(copyList, copyBlock)
			}
			copyBlock = sim.freeList[i]
		}
	}
	if copyBlock.size > 0 {
		copyList = append(copyList, copyBlock)
	}
	sim.freeList = copyList
}

func (sim *Simulation) blockSplit(block *Block, newName string, size int){
	/*
		Construct one or two new blocks given a block and
		a specific size to allocate
	*/
	if size > block.size {
		err := "\nShouldn't be splitting this block"
		fmt.Println(err)
		os.Exit(1)
	} else if size == block.size {
		block.name = newName
		sim.usedList = append(sim.usedList, *block)
		blkIndex := 0
		//find block in free list and delete
		for i, _ := range sim.freeList {
			blkIndex = i
			if sim.freeList[i] == *block{
				break
			}
		}
		sim.freeList = append(sim.freeList[:blkIndex], sim.freeList[blkIndex+1:]...)
	} else {
		newBlock := Block{name:newName, size:size, offset:block.offset}
		block.offset += size
		block.size -= size
		sim.usedList = append(sim.usedList, newBlock)
	}
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

func printStats(pctFree float64, pctUsed float64, failed int) {
	/*
		Prints stats of simulation
	 */
	fmt.Println("Percent of used memory:", pctUsed, "%")
	fmt.Println("Percent of free memory:", pctFree, "%")
	fmt.Println("Number of failed allocations:", failed)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sim := start(scanner.Text()) //simulation object to run alg on
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) > 0 {
			line := strings.Fields(line)
			if line[0] == "alloc" {
				size, _ := strconv.Atoi(line[2])
				sim.alloc(line[1], size)
			} else if line[0] == "free" {
				sim.free(line[1])
				sim.compactFree()

			} else {
				err := "\nInvalid line in input file"
				fmt.Println(err)
				os.Exit(1)
			}
			sim.printLists()
		}
	}
	//print stats
	pctFree, pctUsed := sim.getStats()
	printStats(pctFree, pctUsed, sim.failed)
}