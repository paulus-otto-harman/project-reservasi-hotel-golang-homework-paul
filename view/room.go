package view

import (
	"context"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
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
		gola.Wait("Tambah Data Kamar")
	case 2:
		gola.Wait("Hapus Data Kamar")
	case 3:
		gola.Wait("Ubah Data Kamar")
	case 4:
		gola.Wait("Tambah Data Kamar")
	}
	return -1
}
