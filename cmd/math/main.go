package main

import (
	"github.com/surdeus/gox/src/gx"
	"fmt"
)

func main() {
	p := gx.Polygon{
		gx.Vertex{0, 0},
		gx.Vertex{0, 50},
		gx.Vertex{25, 50},
		gx.Vertex{30, 40},
		gx.Vertex{40, 30},
	}
	
	fmt.Println("barycenter:", p.Barycenter())
	fmt.Println("edges:", p.Edges())
}



