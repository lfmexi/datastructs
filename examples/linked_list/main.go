package main

import (
	"fmt"

	"github.com/lfmexi/datastructs/sets"
)

func main() {
	list := sets.NewLinkedList()

	list.Add("Hi")
	list.Add("There")

	list.Add("There2")
	list.Add("There3")

	fmt.Println(list.Size)

	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))
	fmt.Println(list.Get(3))

	list.Delete(2)
	fmt.Println(list.Size)

	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))

}
