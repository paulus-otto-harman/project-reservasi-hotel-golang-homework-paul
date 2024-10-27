package view

import (
	"context"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
)

type RetrieveReservation struct {
}

func (screen *RetrieveReservation) Render(ctx context.Context) int {
	util.H1("Daftar Reservasi")
	gola.Wait("Tekan sembarang tombol untuk kembali ke menu Kelola Reservasi")
	return 0
}
