package main

//
//import (
//	"fmt"
//	"github.com/team142/Forge-of-Hephaestus/server/model"
//)
//
//func main() {
//	fmt.Println("Starting...")
//	currentPosition := model.NamedLocation{
//		Name: "start",
//		X:    -100,
//		Y:    0,
//		Z:    0,
//	}
//	finishPosition := model.NamedLocation{
//		Name: "finish",
//		X:    3,
//		Y:    0,
//		Z:    4,
//	}
//	highway := []model.NamedLocation{
//		{
//			Name: "hw-1",
//			X:    0,
//			Y:    0,
//			Z:    1,
//		},
//		{
//			Name: "hw-2",
//			X:    0,
//			Y:    0,
//			Z:    4,
//		},
//		{
//			Name: "hw-3",
//			X:    3,
//			Y:    0,
//			Z:    0,
//		},
//		{
//			Name: "hw-4",
//			X:    300,
//			Y:    0,
//			Z:    0,
//		},
//		{
//			Name: "hw-5",
//			X:    301,
//			Y:    0,
//			Z:    0,
//		},
//		{
//			Name: "hw-6",
//			X:    2,
//			Y:    0,
//			Z:    4,
//		},
//		{
//			Name: "hw-7",
//			X:    2,
//			Y:    0,
//			Z:    0,
//		},
//	}
//	result := model.ShortestPath(currentPosition, finishPosition, highway)
//	fmt.Println(fmt.Sprint("Path: ", len(result)))
//	for _, next := range result {
//		fmt.Println(fmt.Sprint(next.Name, "... ", next.X, "... ", next.Z))
//	}
//
//}
