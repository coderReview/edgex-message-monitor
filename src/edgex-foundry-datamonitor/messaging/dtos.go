// Copyright 2021 Alessandro De Blasis <alex@deblasis.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package messaging

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

func ParseEvent(eventPayloadBytes []byte) (*dtos.Event, error) {
	type eventEnvelope struct {
		Event *dtos.Event `json:"event"`
	}
	e := &eventEnvelope{}
	if err := json.Unmarshal(eventPayloadBytes, e); err != nil {
		return nil, err
	}
	return e.Event, nil
}

type Envelope struct {
	Serial  int64
	EventId string
	Topic   string
	Content interface{}
}
