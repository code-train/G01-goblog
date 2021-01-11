package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

type Flashes map[string]interface{}

var flashKey = "_flashes"

func init() {
	gob.Register(Flashes{})
}

func Info(message string) {
	addFlush("info", message)
}

func Warning(message string) {
	addFlush("info", message)
}

func Success(message string) {
	addFlush("info", message)
}

func Danger(message string) {
	addFlush("info", message)
}

func All() Flashes {
	val := session.Get(flashKey)
	flashMessage, ok := val.(Flashes)

	if !ok {
		return nil
	}
	session.Forget(flashKey)
	return flashMessage
}

func addFlush(key string, message string)  {
	flashes := Flashes{}

	flashes[key] = message

	session.Put(flashKey, flashes)
	session.Save()
}