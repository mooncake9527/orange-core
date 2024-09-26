package ebus

import "github.com/mooncake9527/x/eventbus"

var EventBus = eventbus.New()

const (
	EventApplicationStarted = "application:started"
	EventApplicationQuit    = "application:quit"
	EventCoreInit           = "application:core:init"
)
