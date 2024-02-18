package pkg

import (
	"fmt"
	"log"
	"net"

	"github.com/ayush2419/cache/internal/cache"
	"github.com/ayush2419/cache/internal/command"
)

type Handler struct {
	cache cache.ICache
}

func NewHandler(cache cache.ICache) *Handler {
	return &Handler{cache: cache}
}

func (h *Handler) HandleRequestMessage(conn net.Conn, message *command.Message) {
	switch message.Cmd {
	case command.CMDSet:
		err := h.cache.Set(message.Key, message.Value, message.TTL)
		if err != nil {
			log.Fatalf("Error in cache SET, message: %v", message)
			conn.Write([]byte(err.Error()))
			return
		}

		conn.Write([]byte("Saved cache successfully"))
		return

	case command.CMDGet:
		value, found, err := h.cache.Get(message.Key)
		if err != nil {
			log.Fatalf("Error in cache GET, message: %v", message)
			conn.Write([]byte(err.Error()))
			return
		}

		if !found {
			conn.Write([]byte(fmt.Sprintf("Cache empty for key: %s", string(message.Key))))
			return
		}

		conn.Write([]byte(fmt.Sprintf("Cache found: %s", string(value))))
		return

	case command.CMDDelete:
		done, err := h.cache.Delete(message.Key)
		if err != nil {
			log.Fatalf("Error in cache DELETE, message: %v", message)
			conn.Write([]byte(err.Error()))
			return
		}

		if !done {
			conn.Write([]byte(fmt.Sprintf("Cache empty for key: %s", string(message.Key))))
			return
		}

		conn.Write([]byte("Cache deleted successfully"))
		return
	}
}
