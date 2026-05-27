package notifier

import (
	"github.com/crazy-max/diun/v4/internal/model"
)

// Handler is a notifier interface
type Handler interface {
	Name() string
	Send(entry model.NotifEntry) error
}

// Closer is an optional interface for notifiers that hold connections
type Closer interface {
	Close()
}

// Notifier represents an active notifier object
type Notifier struct {
	Handler
}
