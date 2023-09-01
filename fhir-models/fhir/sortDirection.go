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

// SortDirection is documented here http://hl7.org/fhir/ValueSet/sort-direction
type SortDirection int

const (
	SortDirectionAscending SortDirection = iota
	SortDirectionDescending
)

 func (code *SortDirection)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SortDirection) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "ascending":
		*code = SortDirectionAscending
	case "descending":
		*code = SortDirectionDescending
	default:
		return fmt.Errorf("unknown SortDirection code `%s`", s)
	}
	return nil
}
func (code SortDirection) String() string {
	return code.Code()
}
func (code SortDirection) Code() string {
	switch code {
	case SortDirectionAscending:
		return "ascending"
	case SortDirectionDescending:
		return "descending"
	}
	return "<unknown>"
}
func (code SortDirection) Display() string {
	switch code {
	case SortDirectionAscending:
		return "Ascending"
	case SortDirectionDescending:
		return "Descending"
	}
	return "<unknown>"
}
func (code SortDirection) Definition() string {
	switch code {
	case SortDirectionAscending:
		return "Sort by the value ascending, so that lower values appear first."
	case SortDirectionDescending:
		return "Sort by the value descending, so that lower values appear last."
	}
	return "<unknown>"
}
