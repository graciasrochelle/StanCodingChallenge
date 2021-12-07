package main

import (
	"StanCodingChallenge/app/app"
)

func main() {
	app := &app.App{}
	app.Run(":3000")
}
