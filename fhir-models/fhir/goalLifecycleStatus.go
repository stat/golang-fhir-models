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

// GoalLifecycleStatus is documented here http://hl7.org/fhir/ValueSet/goal-status
type GoalLifecycleStatus int

const (
	GoalLifecycleStatusProposed GoalLifecycleStatus = iota
	GoalLifecycleStatusPlanned
	GoalLifecycleStatusAccepted
	GoalLifecycleStatusActive
	GoalLifecycleStatusOnHold
	GoalLifecycleStatusCompleted
	GoalLifecycleStatusCancelled
	GoalLifecycleStatusEnteredInError
	GoalLifecycleStatusRejected
)

 func (code *GoalLifecycleStatus)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GoalLifecycleStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposed":
		*code = GoalLifecycleStatusProposed
	case "planned":
		*code = GoalLifecycleStatusPlanned
	case "accepted":
		*code = GoalLifecycleStatusAccepted
	case "active":
		*code = GoalLifecycleStatusActive
	case "on-hold":
		*code = GoalLifecycleStatusOnHold
	case "completed":
		*code = GoalLifecycleStatusCompleted
	case "cancelled":
		*code = GoalLifecycleStatusCancelled
	case "entered-in-error":
		*code = GoalLifecycleStatusEnteredInError
	case "rejected":
		*code = GoalLifecycleStatusRejected
	default:
		return fmt.Errorf("unknown GoalLifecycleStatus code `%s`", s)
	}
	return nil
}
func (code GoalLifecycleStatus) String() string {
	return code.Code()
}
func (code GoalLifecycleStatus) Code() string {
	switch code {
	case GoalLifecycleStatusProposed:
		return "proposed"
	case GoalLifecycleStatusPlanned:
		return "planned"
	case GoalLifecycleStatusAccepted:
		return "accepted"
	case GoalLifecycleStatusActive:
		return "active"
	case GoalLifecycleStatusOnHold:
		return "on-hold"
	case GoalLifecycleStatusCompleted:
		return "completed"
	case GoalLifecycleStatusCancelled:
		return "cancelled"
	case GoalLifecycleStatusEnteredInError:
		return "entered-in-error"
	case GoalLifecycleStatusRejected:
		return "rejected"
	}
	return "<unknown>"
}
func (code GoalLifecycleStatus) Display() string {
	switch code {
	case GoalLifecycleStatusProposed:
		return "Proposed"
	case GoalLifecycleStatusPlanned:
		return "Planned"
	case GoalLifecycleStatusAccepted:
		return "Accepted"
	case GoalLifecycleStatusActive:
		return "Active"
	case GoalLifecycleStatusOnHold:
		return "On Hold"
	case GoalLifecycleStatusCompleted:
		return "Completed"
	case GoalLifecycleStatusCancelled:
		return "Cancelled"
	case GoalLifecycleStatusEnteredInError:
		return "Entered in Error"
	case GoalLifecycleStatusRejected:
		return "Rejected"
	}
	return "<unknown>"
}
func (code GoalLifecycleStatus) Definition() string {
	switch code {
	case GoalLifecycleStatusProposed:
		return "A goal is proposed for this patient."
	case GoalLifecycleStatusPlanned:
		return "A goal is planned for this patient."
	case GoalLifecycleStatusAccepted:
		return "A proposed goal was accepted or acknowledged."
	case GoalLifecycleStatusActive:
		return "The goal is being sought actively."
	case GoalLifecycleStatusOnHold:
		return "The goal remains a long term objective but is no longer being actively pursued for a temporary period of time."
	case GoalLifecycleStatusCompleted:
		return "The goal is no longer being sought."
	case GoalLifecycleStatusCancelled:
		return "The goal has been abandoned."
	case GoalLifecycleStatusEnteredInError:
		return "The goal was entered in error and voided."
	case GoalLifecycleStatusRejected:
		return "A proposed goal was rejected."
	}
	return "<unknown>"
}
