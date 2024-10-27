package view

import (
	"context"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
	view "homework/view/reservation"
)

type Reservation struct {
}

func (screen *Reservation) Render(ctx context.Context) int {
	util.H1("Kelola Reservasi")
	fmt.Println("[1] Buat Reservasi")
	fmt.Println("[2] Hapus Data Reservasi")
	fmt.Println("[3] Ubah Data Reservasi")
	fmt.Println("[4] Daftar Reservasi")

	fmt.Println()
	fmt.Println("[0] Kembali")

	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[4] atau [0] untuk Kembali"))))
	switch menuItem.(int) {
	case 0:
		return 0
	case 1:
		Render(&view.CreateReservation{}, ctx)
	case 2:
		Render(&view.DeleteReservation{}, ctx)
	case 3:
		Render(&view.UpdateReservation{}, ctx)
	case 4:
		Render(&view.RetrieveReservation{}, ctx)
	}
	return -1
}
