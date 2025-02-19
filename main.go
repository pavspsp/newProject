package main

import (
	"fmt"
	pkg "newProject/src" //
)

func main() {
	sum := pkg.Add(2, 3)
	reversed := pkg.Reverse("hello")
	fmt.Println(sum, reversed) // 5 olleh
}
