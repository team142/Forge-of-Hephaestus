package model

import (
	"fmt"
	"sync"
)

const MaxHopDistance = 6

func ShortestPath(currentPosition, finishPosition NamedLocation, highway []NamedLocation) []NamedLocation {
	result := make([]NamedLocation, 0)
	t := &traversalRecord{
		ShortestCount: currentPosition.Distance2Dimensions(finishPosition) * 3,
		ShortestPath:  make([]NamedLocation, 0),
	}
	traverse(
		true,
		t,
		[]NamedLocation{{Name: "start-location", X: currentPosition.X, Y: currentPosition.Y, Z: currentPosition.Z}},
		0,
		highway,
		finishPosition,
	)
	result = t.ShortestPath
	return result
}

type traversalRecord struct {
	sync.Mutex
	ShortestCount int
	ShortestPath  []NamedLocation
}

func (t *traversalRecord) IsMyCostOverBudget(cost int) bool {
	t.Lock()
	defer t.Unlock()
	return t.ShortestCount <= cost
}

func (t *traversalRecord) Submit(cost int, path []NamedLocation) {
	for _, p := range path {
		fmt.Print(p.Name, " -> ")
	}
	fmt.Print("... ", cost, "\n")
	t.Lock()
	defer t.Unlock()
	if cost < t.ShortestCount {
		t.ShortestPath = path
		t.ShortestCount = cost
	}
	return
}

func traverse(startLocation bool, record *traversalRecord, path []NamedLocation, pathCost int, highway []NamedLocation, finishLocation NamedLocation) {
	//TODO: remove used locations from HIGHWAY
	currentHighway := make([]NamedLocation, 0)
	for _, n := range highway {
		found := false
		for _, p := range path {
			if n.Name == p.Name {
				found = true
			}
		}
		if !found {
			currentHighway = append(currentHighway, n)
		}
	}
	if len(currentHighway) == len(highway) && !startLocation {
		panic("didnt remove anything from the highway! Exiting")
	}

	//TODO: get distance from LAST to FINISH
	distanceLastToFinish := path[len(path)-1].Distance2Dimensions(finishLocation)
	fmt.Println("distanceLastToFinish:", distanceLastToFinish)

	//TODO: get distance from each HIGHWAY to FINISH
	distancesHWToFinish := make(map[string]int)
	for _, n := range currentHighway {
		distancesHWToFinish[n.Name] = n.Distance2Dimensions(finishLocation)
		fmt.Println("tile ", n.Name, " distance to finish: ", distancesHWToFinish[n.Name])
	}

	//TODO: IF LAST is closer than HIGHWAY... go straight to FINISH. Finish. Submit exit.
	isLastClosestToFinish := true
	for _, d := range distancesHWToFinish {
		if d < distanceLastToFinish {
			isLastClosestToFinish = false
			break
		}
	}
	fmt.Println("isLastClosestToFinish: ", isLastClosestToFinish)
	if isLastClosestToFinish {
		currentPath := path
		currentCost := pathCost + path[len(path)-1].Distance2Dimensions(finishLocation)
		currentPath = append(currentPath, finishLocation)
		record.Submit(currentCost, currentPath)
		fmt.Println("---------------------------------------------")
		return
	}

	//TODO: if nodes empty, go straight to FINISH. Finish. Submit exit.
	if len(currentHighway) == 0 {
		currentPath := path
		currentCost := pathCost + path[len(path)-1].Distance2Dimensions(finishLocation)
		currentPath = append(currentPath, finishLocation)
		record.Submit(currentCost, currentPath)
		fmt.Println("---------------------------------------------")
		return
	}

	//TODO: if first - find closest one
	if startLocation {
		shortestOne := record.ShortestCount
		var shortestNamedLocation NamedLocation
		for _, next := range currentHighway {
			d := path[len(path)-1].Distance2Dimensions(next)
			fmt.Println("start to ", next.Name, " is ", d)
			if d < shortestOne {
				fmt.Println("... and its the shortest now")
				shortestNamedLocation = next
				shortestOne = d
			}
		}
		if shortestNamedLocation.Name == "" {
			panic("no shortest path to NamedLocation")
		}
		//TODO: traverse from here
		fmt.Println(">>>>>>>>>>>>>>>>>> ADDING ", shortestNamedLocation.Name)
		currentPath := path
		currentCost := pathCost + path[len(path)-1].Distance2Dimensions(shortestNamedLocation)
		currentPath = append(currentPath, shortestNamedLocation)
		traverse(false, record, currentPath, currentCost, currentHighway, finishLocation)
		return
	}

	for _, next := range currentHighway {
		//TODO: copy weight
		currentCost := pathCost

		//TODO: copy path
		currentPath := path

		//TODO: add weight to path
		thisCost := path[len(path)-1].Distance2Dimensions(next)
		currentCost = currentCost + thisCost

		//TODO: REMOVE IF COST > 6
		if thisCost > MaxHopDistance {
			// We don't consider jumps bigger than this
			fmt.Println("!!! no longer considering route as > than MaxHopDistance", next.Name)
			continue
		}
		//TODO:
		if record.IsMyCostOverBudget(currentCost) {
			fmt.Println("!!! no longer considering route as total cost over budget:", path[len(path)-1].Name, "->", next.Name)
			continue
		}
		// REMOVE IF OVER OTHER COSTS
		//

		//TODO: add next to path
		currentPath = append(currentPath, next)

		//TODO: call traverse(......)
		fmt.Println(">>> going down path ", next.Name)
		traverse(false, record, currentPath, currentCost, currentHighway, finishLocation)

	}
}
