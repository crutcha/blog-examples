package main

import "fmt"

type PathTable struct {
	dest map[string]int
	prev map[string]int
	cost int
}

func (p *PathTable) String() {
	fmt.Println("not implemented yet")
}

func (g *Graph) DijkstraPaths(src string) PathTable {
	fmt.Println("yep")
	return PathTable{}
}
