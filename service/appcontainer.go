package service

import (
	"context"
	"homework/view"
	"sync"
	"time"
)

func sessionTimeout(timeout int) time.Duration {
	if timeout == 0 {
		return time.Until(time.Now().AddDate(10, 0, 0))
	}
	return time.Duration(timeout) * time.Second
}

func AppContainer(wg *sync.WaitGroup, timeout int) {
	defer wg.Done()
	sessionLifetime := sessionTimeout(timeout)
	for {
		loginScreen := view.Login{}
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), sessionLifetime)
			defer cancel()

			view.Render(&loginScreen, ctx)
		}()
		if loginScreen.Username == "0" {
			return
		}
	}
}
