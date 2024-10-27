package view

import (
	"context"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
)

type DeleteRoom struct {
}

func (screen *DeleteRoom) Render(ctx context.Context) int {
	util.H1("Hapus Data Kamar")

	cont, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", "Enter atau [0] untuk kembali"))))
	if cont == "0" {
		return 0
	}
	return -1
}
