package model

import "github.com/team142/Forge-of-Hephaestus/server/util"

type NamedLocation struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Z    int    `json:"z"`
}

func (first NamedLocation) Distance2Dimensions(other NamedLocation) int {
	return util.Distance2Dimensions(first.X, first.Z, other.X, other.Z)
}
