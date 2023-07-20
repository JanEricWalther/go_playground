package graph

type GraphList [][]GNode

type GraphMatrix [][]int

type GNode struct {
	to     int
	weight int
}
