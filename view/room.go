package view

import (
	"context"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
	view "homework/view/room"
)

type Room struct {
}

func (screen *Room) Render(ctx context.Context) int {
	util.H1("Kelola Data Kamar")
	fmt.Println("[1] Tambah Data Kamar")
	fmt.Println("[2] Hapus Data Kamar")
	fmt.Println("[3] Ubah Data Kamar")
	fmt.Println("[4] Daftar Kamar")

	fmt.Println()
	fmt.Println("[0] Kembali")

	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[4] atau [0] untuk Kembali"))))
	switch menuItem.(int) {
	case 0:
		return 0
	case 1:
		Render(&view.CreateRoom{}, ctx)
	case 2:
		Render(&view.DeleteRoom{}, ctx)
	case 3:
		Render(&view.UpdateRoom{}, ctx)
	case 4:
		Render(&view.RetrieveRoom{}, ctx)
	}
	return -1
}
