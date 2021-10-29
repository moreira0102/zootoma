package main

import (
	"fmt"
	"tomabase/test/experimentation/path"
)

func main() {
	// p := new(path.Path)
	// t1 := new(path.Path)
	// t2 := new(path.Path)

	// t1.Way = "/"
	// t2.Way = "/home/teste"

	// p.Child = append(p.Child, t2)
	// // t := new(path.Path)

	// p.Way = "/home"

	// fmt.Println(p.Child[0].Way)

	mapper := path.NewPathMap()
	p := path.NewPath("/home/lucas/data.dat", mapper)

	fmt.Println((*(*mapper).Pathmap[p.Way]).Way)
	p.ShowPath()
}