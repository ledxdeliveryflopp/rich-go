package client

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hugolgst/rich-go/ipc"
	"gopkg.in/natefinch/npipe.v2"
	"log"
	"os"
	"time"
)

// Login sends a handshake in the socket and returns an error or nil
func Login(clientId string) error {
	payload, err := json.Marshal(Handshake{"1", clientId})
	if err != nil {
		return err
	}

	err = ipc.OpenSocket()
	if err != nil {
		return err
	}

	_, err = ipc.Send(0, string(payload))
	if err != nil {
		return err
	}

	return nil
}

func SetActivity(clientId string, activity Activity) error {

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			mapActivity(&activity),
		},
		getNonce(),
	})

	if err != nil {
		return err
	}

	_, err = ipc.Send(1, string(payload))
	if errors.Is(err, npipe.ErrClosed) {
		log.Println("pipe closed, try reconnect")
		time.Sleep(time.Second * 10)
		err := Login(clientId)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func getNonce() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
