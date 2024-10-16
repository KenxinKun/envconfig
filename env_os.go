//go:build appengine || go1.5
// +build appengine go1.5

package envconfig

import "os"

var lookupEnv = os.LookupEnv

func SetLookupEnv(f func(key string) (string, bool)) {
	lookupEnv = f
}
