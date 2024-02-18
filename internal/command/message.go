package command

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

func ParseMessage(rawMsg string) (message *Message, err error) {
	parts := strings.Split(rawMsg, " ")

	cmd := Command(parts[0])
	key := []byte(parts[1])

	if !cmd.isValid() {
		err = fmt.Errorf("invalid command type")
		return
	}

	message = &Message{
		Cmd: cmd,
		Key: key,
	}

	switch cmd {
	case CMDSet:
		if len(parts) < 4 {
			err = fmt.Errorf("invalid command for SET")
			return
		}

		ttl, _ := strconv.Atoi(parts[3])

		message.Value = []byte(parts[2])
		message.TTL = time.Duration(time.Duration(rand.Intn(ttl)) * time.Second)
		return
	case CMDGet, CMDDelete:
		return
	}
	return
}
