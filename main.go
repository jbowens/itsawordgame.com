package main

import (
	"fmt"

	"github.com/jbowens/itsawordgame.com/app"
	"github.com/jbowens/itsawordgame.com/app/game"
)

func main() {

	b := game.NewBoard(5, 5)
	fmt.Println(b.String())

	var a app.App
	a.Start()
	a.Wait()
}
