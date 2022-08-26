package main

import (
	"log"
	"math"
)

func twoOptSwap(route []int, v1, v2 int) []int {
	newRoute := make([]int, len(route))
	reverseCounter := v2
	for i := 0; i < len(route); i++ {
		if i < v1 || i > v2 {
			newRoute[i] = route[i]
		} else {
			newRoute[i] = route[reverseCounter]
			reverseCounter = reverseCounter - 1
		}
	}
	return newRoute
}

func calculateTotalDistance(order []int, costMatrix [][]int) int {
	totalCost := 0
	for i := 1; i < len(order); i++ {
		totalCost = totalCost + costMatrix[order[i-1]][order[i]]
	}
	return totalCost
}

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
	totalCost := 0
	visitedPoint[curIndex] = true

	for i := 0; len(bestPath) != len(points); i++ {
		lowestCostNextPoint := 0
		lowestCost := math.MaxInt
		for j := 0; j < len(points); j++ {
			if costMatrix[curIndex][j] < lowestCost && !visitedPoint[j] && i != j {
				lowestCostNextPoint = j
				lowestCost = costMatrix[curIndex][j]
			}
		}
		bestPath = append(bestPath, lowestCostNextPoint)
		visitedPoint[lowestCostNextPoint] = true
		curIndex = lowestCostNextPoint
		totalCost = totalCost + lowestCost
	}

	log.Println(bestPath)
	log.Println(totalCost)
	log.Println(calculateTotalDistance(bestPath, costMatrix))

	//====================
	initalOrder := bestPath
	initalOrderCost := calculateTotalDistance(bestPath, costMatrix)
	twoOptCost := initalOrderCost
	twoOptOrder := make([]int, len(initalOrder))
	copy(twoOptOrder, initalOrder)
	optimized := true

	for optimized {
		optimized = false
		for i := 1; i < len(twoOptOrder)-2; i++ {
			for j := i + 1; j < len(twoOptOrder)-1; j++ {
				newRoute := twoOptSwap(twoOptOrder, i, j)
				newCost := calculateTotalDistance(newRoute, costMatrix)
				if newCost < twoOptCost {
					twoOptCost = newCost
					twoOptOrder = newRoute
					log.Println("optimized!")
					log.Println(twoOptCost)
					log.Println(twoOptOrder)
					optimized = true
					break
				}
			}
			if optimized {
				break
			}
		}
	}
	log.Println(twoOptCost)
	log.Println(twoOptOrder)
}
