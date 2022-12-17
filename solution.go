package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	worldmap := emoji.Sprint("Hello :world_map:")
	fmt.Println(worldmap)
}
