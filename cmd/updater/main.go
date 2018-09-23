package main

import (
	"github.com/nathan-osman/go-updater/dialog"
)

func main() {

	// Create the dialog
	d := dialog.New()

	// Create a channel for cancel events from the button
	cancelCh := make(chan bool)
	go func() {
		select {
		case <-cancelCh:
			d.Close()
		}
	}()

	// Run the event loop
	d.Exec(cancelCh)
}
