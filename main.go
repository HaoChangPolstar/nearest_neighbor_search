package main

import (
	"log"
)

func main() {

	var costMatrix = [][]int{
		{0, 2, 5, 7, 6}, // a b c d e
		{2, 0, 8, 3, 7}, // b
		{5, 8, 0, 1, 8}, // c
		{7, 3, 1, 0, 9}, // d
		{5, 6, 7, 8, 0}} // e

	points := []int{0, 1, 2, 3, 4}

	visitedPoint := map[int]bool{}

	bestPath := []int{0}

	curIndex := 0
	visitedPoint[curIndex] = true
	for i := 0; len(bestPath) != len(points); i++ {
		lowestCostNextPoint := 0
		lowestCost := 999
		for j := 0; j < len(points); j++ {
			if costMatrix[curIndex][j] < lowestCost && !visitedPoint[j] {
				lowestCostNextPoint = j
				lowestCost = costMatrix[i][j]
			}
			log.Println(i, j, costMatrix[i][j], lowestCost, lowestCostNextPoint)
		}
		bestPath = append(bestPath, lowestCostNextPoint)
		visitedPoint[lowestCostNextPoint] = true
		curIndex = lowestCostNextPoint
	}
	log.Println(bestPath)
}
