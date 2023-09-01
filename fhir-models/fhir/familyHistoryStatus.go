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

// THIS FILE IS GENERATED BY https://github.com/stat/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// FamilyHistoryStatus is documented here http://hl7.org/fhir/ValueSet/history-status
type FamilyHistoryStatus int

const (
	FamilyHistoryStatusPartial FamilyHistoryStatus = iota
	FamilyHistoryStatusCompleted
	FamilyHistoryStatusEnteredInError
	FamilyHistoryStatusHealthUnknown
)

 func (code *FamilyHistoryStatus)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FamilyHistoryStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "partial":
		*code = FamilyHistoryStatusPartial
	case "completed":
		*code = FamilyHistoryStatusCompleted
	case "entered-in-error":
		*code = FamilyHistoryStatusEnteredInError
	case "health-unknown":
		*code = FamilyHistoryStatusHealthUnknown
	default:
		return fmt.Errorf("unknown FamilyHistoryStatus code `%s`", s)
	}
	return nil
}
func (code FamilyHistoryStatus) String() string {
	return code.Code()
}
func (code FamilyHistoryStatus) Code() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "partial"
	case FamilyHistoryStatusCompleted:
		return "completed"
	case FamilyHistoryStatusEnteredInError:
		return "entered-in-error"
	case FamilyHistoryStatusHealthUnknown:
		return "health-unknown"
	}
	return "<unknown>"
}
func (code FamilyHistoryStatus) Display() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "Partial"
	case FamilyHistoryStatusCompleted:
		return "Completed"
	case FamilyHistoryStatusEnteredInError:
		return "Entered in Error"
	case FamilyHistoryStatusHealthUnknown:
		return "Health Unknown"
	}
	return "<unknown>"
}
func (code FamilyHistoryStatus) Definition() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "Some health information is known and captured, but not complete - see notes for details."
	case FamilyHistoryStatusCompleted:
		return "All available related health information is captured as of the date (and possibly time) when the family member history was taken."
	case FamilyHistoryStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case FamilyHistoryStatusHealthUnknown:
		return "Health information for this family member is unavailable/unknown."
	}
	return "<unknown>"
}
