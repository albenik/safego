package safego

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

type RecoverFunc func(r interface{}, stack string)

func Go(fn func(), rfn RecoverFunc) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				rfn(r, string(debug.Stack()))
			}
		}()
		fn()
	}()
}

func Gol(label string, fn func(), log logrus.FieldLogger) {
	go func() {
		log.Debugf("go(%s): begin", label)
		defer func() {
			if r := recover(); r == nil {
				log.Debugf("go(%s): end", label)
			} else {
				log.Errorf("go(%s) error: %v\n%v", label, r, debug.Stack())
			}
		}()
		fn()
	}()
}
