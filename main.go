package main

import (
	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("vstart", func(w *unison.Window) {
		w.Content().AddChild(vstart.New().Layout())
	})
}
