package view

import (
	"context"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
)

type CreateRoom struct {
}

func (screen *CreateRoom) Render(ctx context.Context) int {
	util.H1("Tambah Data Kamar")

	cont, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", "Enter atau [0] untuk kembali"))))
	if cont == "0" {
		return 0
	}
	return -1
}
