// Copyright 2019 - 2022 The Samply Community
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

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// SubscriptionStatus is documented here http://hl7.org/fhir/ValueSet/subscription-status
type SubscriptionStatus int

const (
	SubscriptionStatusRequested SubscriptionStatus = iota
	SubscriptionStatusActive
	SubscriptionStatusError
	SubscriptionStatusOff
)

 func (code *SubscriptionStatus)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SubscriptionStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "requested":
		*code = SubscriptionStatusRequested
	case "active":
		*code = SubscriptionStatusActive
	case "error":
		*code = SubscriptionStatusError
	case "off":
		*code = SubscriptionStatusOff
	default:
		return fmt.Errorf("unknown SubscriptionStatus code `%s`", s)
	}
	return nil
}
func (code SubscriptionStatus) String() string {
	return code.Code()
}
func (code SubscriptionStatus) Code() string {
	switch code {
	case SubscriptionStatusRequested:
		return "requested"
	case SubscriptionStatusActive:
		return "active"
	case SubscriptionStatusError:
		return "error"
	case SubscriptionStatusOff:
		return "off"
	}
	return "<unknown>"
}
func (code SubscriptionStatus) Display() string {
	switch code {
	case SubscriptionStatusRequested:
		return "Requested"
	case SubscriptionStatusActive:
		return "Active"
	case SubscriptionStatusError:
		return "Error"
	case SubscriptionStatusOff:
		return "Off"
	}
	return "<unknown>"
}
func (code SubscriptionStatus) Definition() string {
	switch code {
	case SubscriptionStatusRequested:
		return "The client has requested the subscription, and the server has not yet set it up."
	case SubscriptionStatusActive:
		return "The subscription is active."
	case SubscriptionStatusError:
		return "The server has an error executing the notification."
	case SubscriptionStatusOff:
		return "Too many errors have occurred or the subscription has expired."
	}
	return "<unknown>"
}
