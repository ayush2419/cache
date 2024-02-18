package cache

import "time"

type ICache interface {
	Set(key []byte, value []byte, ttl time.Duration) (err error)
	Get(key []byte) (value []byte, found bool, err error)
	Delete(key []byte) (done bool, err error)
}
