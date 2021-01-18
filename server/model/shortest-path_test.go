package model

import (
	"fmt"
	"testing"
)

//ShortestPath(currentPosition, finishPosition NamedLocation, highway []NamedLocation) []NamedLocation {
func TestShortestPath2(t *testing.T) {
	currentPosition := NamedLocation{
		Name: "start",
		X:    0,
		Y:    0,
		Z:    0,
	}
	finishPosition := NamedLocation{
		Name: "finish",
		X:    3,
		Y:    0,
		Z:    4,
	}
	highway := []NamedLocation{
		{
			Name: "hw-1",
			X:    0,
			Y:    0,
			Z:    1,
		},
		{
			Name: "hw-2",
			X:    0,
			Y:    0,
			Z:    4,
		},
		{
			Name: "hw-3",
			X:    1,
			Y:    0,
			Z:    0,
		},
	}
	result := ShortestPath(currentPosition, finishPosition, highway)
	t.Log(fmt.Sprint("Path: ", len(result)))
	for _, next := range result {
		t.Log(fmt.Sprint(next.Name, "... ", next.X, "... ", next.Z))
	}

}
