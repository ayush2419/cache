package cache

import (
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Set(key []byte, value []byte, ttl time.Duration) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	go func() {
		<-time.After(ttl)
		delete(c.data, string(key))
	}()

	c.data[string(key)] = value
	return
}

func (c *Cache) Get(key []byte) (value []byte, found bool, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	value, found = c.data[string(key)]
	if !found {
		return
	}

	return
}

func (c *Cache) Delete(key []byte) (done bool, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, found := c.data[string(key)]
	if found {
		done = true
		delete(c.data, string(key))
	}

	return
}
