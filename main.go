package main

import (
	"vstart"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("vstart", func(w *unison.Window) {
		vstart.New().Layout(w.Content())
	})
}
