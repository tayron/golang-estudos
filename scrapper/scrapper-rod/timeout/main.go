package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage()

	err := rod.Try(func() {
		page.Timeout(0 * time.Second).MustNavigate("https://www.hospeda.app")
		fmt.Println(page.MustHTML())
	})
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("timeout error")
	} else if err != nil {
		fmt.Println("other types of error")
	}
}
