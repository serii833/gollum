package consumer

import (
	"github.com/trivago/gollum/log"
	"github.com/trivago/gollum/shared"
	"io"
	"os"
	"sync"
)

const (
	consoleBufferGrowSize = 256
)

// Console consumer plugin
// Configuration example
//
// - "consumer.Console":
//   Enable: true
//
// This consumer does not define any options beside the standard ones.
type Console struct {
	standardConsumer
}

func init() {
	shared.RuntimeType.Register(Console{})
}

func (cons *Console) readFrom(stream io.Reader, threads *sync.WaitGroup) {
	buffer := shared.CreateBufferedReader(consoleBufferGrowSize, cons.postMessageFromSlice)

	for {
		err := buffer.Read(stream, "\n")
		if err != nil {
			Log.Error.Print("Error reading stdin: ", err)
		}
	}
}

// Consume listens to stdin.
func (cons Console) Consume(threads *sync.WaitGroup) {
	go cons.readFrom(os.Stdin, threads)

	defer cons.markAsDone()
	cons.defaultControlLoop(threads)
}
