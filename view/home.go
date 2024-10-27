package view

import (
	"context"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
)

type Home struct {
	isLogout bool
}

func (home *Home) Render(ctx context.Context) int {
	util.H1("Menu Utama")
	fmt.Println("[1] Kelola Kamar")
	fmt.Println("[2] Reservasi")

	fmt.Println()
	fmt.Println("[0] Logout")

	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[5] atau [0] untuk Logout"))))
	switch menuItem.(int) {
	case 0:
		home.isLogout = true
		return 0
	case 1:
		Render(&Room{}, ctx)
	case 2:
		Render(&Room{}, ctx)
	}
	return -1
}
