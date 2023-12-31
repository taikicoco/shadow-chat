package resolver

import (
	"server/graphql/generated/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MessageID map[int64][]chan<- *model.Message
	Mutex     sync.Mutex
}
