// Copyright 2015-2016 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package producer

import (
	"github.com/trivago/gollum/core"
	"sync"
	"time"
)

// Null producer plugin
// This producer does nothing and provides only bare-bone configuration (i.e.
// enabled and streams). Use this producer to test consumer performance.
// This producer does not implement a fuse breaker.
type Null struct {
	control chan core.PluginControl
	streams []core.MessageStreamID
}

func init() {
	core.TypeRegistry.Register(Null{})
}

// Configure initializes the basic members
func (prod *Null) Configure(conf core.PluginConfigReader) error {
	prod.control = make(chan core.PluginControl, 1)
	prod.streams = conf.GetStreamArray("Stream", []core.MessageStreamID{core.GetStreamID("conf.ID")})

	return conf.Errors.OrNil()
}

// GetState always returns PluginStateActive
func (prod *Null) GetState() core.PluginState {
	return core.PluginStateActive
}

// IsActive always returns true
func (prod *Null) IsActive() bool {
	return true
}

// IsBlocked always returns false
func (prod *Null) IsBlocked() bool {
	return false
}

// Streams returns the streams this producer is listening to.
func (prod *Null) Streams() []core.MessageStreamID {
	return prod.streams
}

// GetDropStreamID returns the id of the stream to drop messages to.
func (prod *Null) GetDropStreamID() core.MessageStreamID {
	return core.DroppedStreamID
}

// Control returns write access to this producer's control channel.
func (prod *Null) Control() chan<- core.PluginControl {
	return prod.control
}

// Enqueue simply ignores the message
func (prod *Null) Enqueue(msg core.Message, timeout *time.Duration) {
}

// Produce writes to a buffer that is dumped to a file.
func (prod *Null) Produce(threads *sync.WaitGroup) {
	for {
		command := <-prod.control
		if command == core.PluginControlStopConsumer {
			return // ### return ###
		}
	}
}
