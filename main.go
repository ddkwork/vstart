package main

import (
	"vstart"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("vstart", func(w *unison.Window) {
		w.Content().AddChild(vstart.New().Layout())
	})
}
