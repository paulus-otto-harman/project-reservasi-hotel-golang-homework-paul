package view

import (
	"context"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
)

type RetrieveRoom struct {
}

func (screen *RetrieveRoom) Render(ctx context.Context) int {
	util.H1("Daftar Kamar")
	gola.Wait("Tekan sembarang tombol untuk kembali ke menu Kelola Data Kamar")
	return 0
}
