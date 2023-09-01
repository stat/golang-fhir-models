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

// FHIRSubstanceStatus is documented here http://hl7.org/fhir/ValueSet/substance-status
type FHIRSubstanceStatus int

const (
	FHIRSubstanceStatusActive FHIRSubstanceStatus = iota
	FHIRSubstanceStatusInactive
	FHIRSubstanceStatusEnteredInError
)

 func (code *FHIRSubstanceStatus)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRSubstanceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FHIRSubstanceStatusActive
	case "inactive":
		*code = FHIRSubstanceStatusInactive
	case "entered-in-error":
		*code = FHIRSubstanceStatusEnteredInError
	default:
		return fmt.Errorf("unknown FHIRSubstanceStatus code `%s`", s)
	}
	return nil
}
func (code FHIRSubstanceStatus) String() string {
	return code.Code()
}
func (code FHIRSubstanceStatus) Code() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "active"
	case FHIRSubstanceStatusInactive:
		return "inactive"
	case FHIRSubstanceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FHIRSubstanceStatus) Display() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "Active"
	case FHIRSubstanceStatusInactive:
		return "Inactive"
	case FHIRSubstanceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FHIRSubstanceStatus) Definition() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "The substance is considered for use or reference."
	case FHIRSubstanceStatusInactive:
		return "The substance is considered for reference, but not for use."
	case FHIRSubstanceStatusEnteredInError:
		return "The substance was entered in error."
	}
	return "<unknown>"
}
