package view

import (
	"context"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"homework/util"
)

type Login struct {
	Username string
}

func (login *Login) Render(ctx context.Context) int {
	util.H1("Login")

	for err := errors.New(""); err != nil; {
		login.Username, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Username atau [0] untuk keluar")))))
	}

	if login.Username == "0" {
		return 0
	}

	password, _ := gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Password"))))

	if login.Username == "x" && password == "x" {
		menuUtama := Home{}
		Render(&menuUtama, ctx)
		if !menuUtama.isLogout {
			gola.Wait("Sesi telah berakhir. Tekan Enter untuk login kembali")
		}
	}
	return -1
}
