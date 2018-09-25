package main

import (
	"github.com/nathan-osman/go-updater/dialog"
)

func main() {

	// Create the dialog
	d := dialog.New()

	// Close the dialog when canceled
	go func() {
		select {
		case <-d.Context().Done():
			d.Close()
		}
	}()

	// Run the event loop
	d.Exec()
}
