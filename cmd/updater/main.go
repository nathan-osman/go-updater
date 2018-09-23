package main

import (
	"context"

	"github.com/nathan-osman/go-updater/dialog"
)

func main() {

	// Create the dialog
	d := dialog.New()

	// Create a context for canceling the operation
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			d.Close()
		}
	}()

	// Run the event loop
	d.Exec(cancelFunc)
}
