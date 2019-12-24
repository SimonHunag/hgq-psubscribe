package psubscribe

import "sync"

var factory *HandleFactory
var handleMap map[string]Handle

var (
	once sync.Once
)

func getHandleFactory() *HandleFactory {

	once.Do(func() {
		factory = new(HandleFactory)
		handleMap = make(map[string]Handle)
		handleMap["default"] = new(DefaultHandle)
	})

	return factory
}

func (factory *HandleFactory) buildHandle(handles map[string]Handle) {

	if handleMap == nil {
		once.Do(func() {
			handleMap = make(map[string]Handle)
			handleMap["default"] = new(DefaultHandle)
		})
	}

	for k, v := range handles {
		handleMap[k] = v
	}
}

type HandleFactory struct {
}

func (factory *HandleFactory) getHandle(channel string) Handle {
	var handle Handle

	handle = handleMap[channel]
	if handle == nil {
		handle = handleMap["default"]
	}

	return handle
}
