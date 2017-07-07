package safego

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
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
		log.Debugf("go(%s) begin", label)
		defer func() {
			if r := recover(); r == nil {
				log.Debugf("go(%s) end", label)
			} else {
				log.Errorf("go(%s) panic %v\n%v", label, r, string(debug.Stack()))
			}
		}()
		fn()
	}()
}

func Goe(label string, fn func() error, log logrus.FieldLogger) {
	go func() {
		log.Debugf("go(%s) begin", label)
		defer func() {
			if r := recover(); r == nil {
				log.Debugf("go(%s) end", label)
			} else {
				log.Errorf("go(%s) panic %v\n%v", label, r, string(debug.Stack()))
			}
		}()
		if err := fn(); err != nil {
			log.WithError(err).Errorf("go(%s) func error", label)
		} else {
			log.Debugf("go(%s) func success", label)
		}
	}()
}
