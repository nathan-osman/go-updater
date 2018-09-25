package dialog

import (
	"context"
)

// Dialog describes the interface for a dialog box that displays status and progress information onscreen during the update process.
// Types that implement this interface should show a dialog using the platform's native API methods.
type Dialog interface {
	Context() context.Context
	Exec()
	SetStatus(string)
	SetProgress(int)
	Close()
}
