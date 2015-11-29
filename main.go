package main

import "github.com/jbowens/itsawordgame.com/app"

func main() {
	var a app.App
	a.Start()
	a.Wait()
}
