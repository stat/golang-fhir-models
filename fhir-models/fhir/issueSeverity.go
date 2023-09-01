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

// IssueSeverity is documented here http://hl7.org/fhir/ValueSet/issue-severity
type IssueSeverity int

const (
	IssueSeverityFatal IssueSeverity = iota
	IssueSeverityError
	IssueSeverityWarning
	IssueSeverityInformation
)

 func (code *IssueSeverity)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IssueSeverity) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "fatal":
		*code = IssueSeverityFatal
	case "error":
		*code = IssueSeverityError
	case "warning":
		*code = IssueSeverityWarning
	case "information":
		*code = IssueSeverityInformation
	default:
		return fmt.Errorf("unknown IssueSeverity code `%s`", s)
	}
	return nil
}
func (code IssueSeverity) String() string {
	return code.Code()
}
func (code IssueSeverity) Code() string {
	switch code {
	case IssueSeverityFatal:
		return "fatal"
	case IssueSeverityError:
		return "error"
	case IssueSeverityWarning:
		return "warning"
	case IssueSeverityInformation:
		return "information"
	}
	return "<unknown>"
}
func (code IssueSeverity) Display() string {
	switch code {
	case IssueSeverityFatal:
		return "Fatal"
	case IssueSeverityError:
		return "Error"
	case IssueSeverityWarning:
		return "Warning"
	case IssueSeverityInformation:
		return "Information"
	}
	return "<unknown>"
}
func (code IssueSeverity) Definition() string {
	switch code {
	case IssueSeverityFatal:
		return "The issue caused the action to fail and no further checking could be performed."
	case IssueSeverityError:
		return "The issue is sufficiently important to cause the action to fail."
	case IssueSeverityWarning:
		return "The issue is not important enough to cause the action to fail but may cause it to be performed suboptimally or in a way that is not as desired."
	case IssueSeverityInformation:
		return "The issue has no relation to the degree of success of the action."
	}
	return "<unknown>"
}
