package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	"github.com/spf13/cast"
)

// An Item is something we manage in a priority queue.
type ScoredCave struct {
	Cave   Cave // The value of the item; arbitrary.
	fScore int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*ScoredCave

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].fScore < pq[j].fScore
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ScoredCave)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(cave *ScoredCave, value Cave, score int) {
	cave.Cave = value
	cave.fScore = score
	heap.Fix(pq, cave.index)
}

type Coordinate struct {
	x, y int
}

type Cave struct {
	Coord      Coordinate
	RiskRating int
}

func GetCaveNeighbors(xOrigin, yOrigin, gridHeight, gridWidth int) []Coordinate {
	var rv []Coordinate
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 { //Exclude origin
				continue
			}
			if (x == -1 && y == 1) || (x == -1 && y == -1) || (x == 1 && y == 1) || (x == 1 && y == -1) { //Exclude diagonals
				continue
			}
			if (xOrigin+x >= 0 && xOrigin+x < gridWidth) && (yOrigin+y >= 0 && yOrigin+y < gridHeight) {
				rv = append(rv, Coordinate{x: xOrigin + x, y: yOrigin + y})
			}
		}
	}
	return rv
}

func PrintGrid(grid [][]Cave) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(grid[y][x].RiskRating)
			if x != len(grid[y])-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	caveMap, err := readLines("input-test-large.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	Part1(caveMap)
}

func readLines(path string) ([][]Cave, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//Real values
	//gridHeight := 100
	//gridWidth := 100
	//Test values
	gridHeight := 50
	gridWidth := 50

	caveMap := make([][]Cave, gridHeight)
	for index := range caveMap {
		caveMap[index] = make([]Cave, gridWidth)
	}

	//Populate cave map
	y := 0
	for scanner.Scan() {
		mapLine := scanner.Text()
		for x := 0; x < len(mapLine); x++ {
			caveMap[y][x] = Cave{RiskRating: cast.ToInt(string(mapLine[x])), Coord: Coordinate{x: x, y: y}}
		}
		y++
	}

	return caveMap, err
}

func ReconstructPath(cameFrom map[Coordinate]Cave, currentPosition Cave) int {
	totalPath := []Cave{currentPosition}
	totalRisk := 0
	current := currentPosition
	for {
		if prevPosition, exists := cameFrom[current.Coord]; exists {
			totalPath = append([]Cave{current}, totalPath...)
			totalRisk += current.RiskRating
			current = prevPosition
		} else {
			break
		}
	}
	return totalRisk
}

func GetIdealStepsToGoal(start Cave, goal Cave) int {
	return (goal.Coord.y - start.Coord.y) + (goal.Coord.x - start.Coord.x)
}

func AStar(caveMap [][]Cave) int {
	start := caveMap[0][0]
	goal := caveMap[len(caveMap)-1][len(caveMap[0])-1]
	openSet := make(PriorityQueue, 1)
	openSetMap := make(map[Coordinate]int)
	initialfScore := GetIdealStepsToGoal(start, goal)
	openSet[0] = &ScoredCave{Cave: start, index: 0, fScore: initialfScore}
	openSetMap[start.Coord] = initialfScore
	heap.Init(&openSet)

	cameFrom := make(map[Coordinate]Cave)
	gScore := make(map[Cave]int)
	gScore[start] = 0
	fScore := make(map[Cave]int)
	fScore[start] = initialfScore

	for {
		if len(openSet) == 0 {
			break
		}
		current := heap.Pop(&openSet).(*ScoredCave).Cave
		if current == goal {
			return ReconstructPath(cameFrom, current)
		}

		neighbors := GetCaveNeighbors(current.Coord.x, current.Coord.y, len(caveMap), len(caveMap[0]))
		for _, nCoord := range neighbors {
			n := caveMap[nCoord.y][nCoord.x]
			if _, gCurrentExists := gScore[current]; !gCurrentExists {
				gScore[current] = 9999999999
			}
			if _, gNeighborExists := gScore[n]; !gNeighborExists {
				gScore[n] = 9999999999
			}

			tentativegScore := gScore[current] + n.RiskRating
			if tentativegScore < gScore[n] {
				cameFrom[n.Coord] = current
				neighborfScore := tentativegScore + GetIdealStepsToGoal(n, goal)
				gScore[n] = tentativegScore
				fScore[n] = neighborfScore
				if _, exists := openSetMap[nCoord]; !exists {
					scoredNeighbor := &ScoredCave{
						Cave:   n,
						fScore: neighborfScore,
					}
					heap.Push(&openSet, scoredNeighbor)
					openSetMap[nCoord] = neighborfScore
				}
			}
		}
	}

	fmt.Println("Could not find optimal path")
	return -1
}

func Part1(caveMap [][]Cave) {
	PrintGrid(caveMap)
	lowestRiskRating := AStar(caveMap)
	fmt.Println(lowestRiskRating)
}
