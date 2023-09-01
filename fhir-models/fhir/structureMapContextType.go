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

// StructureMapContextType is documented here http://hl7.org/fhir/ValueSet/map-context-type
type StructureMapContextType int

const (
	StructureMapContextTypeType StructureMapContextType = iota
	StructureMapContextTypeVariable
)

 func (code *StructureMapContextType)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapContextType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "type":
		*code = StructureMapContextTypeType
	case "variable":
		*code = StructureMapContextTypeVariable
	default:
		return fmt.Errorf("unknown StructureMapContextType code `%s`", s)
	}
	return nil
}
func (code StructureMapContextType) String() string {
	return code.Code()
}
func (code StructureMapContextType) Code() string {
	switch code {
	case StructureMapContextTypeType:
		return "type"
	case StructureMapContextTypeVariable:
		return "variable"
	}
	return "<unknown>"
}
func (code StructureMapContextType) Display() string {
	switch code {
	case StructureMapContextTypeType:
		return "Type"
	case StructureMapContextTypeVariable:
		return "Variable"
	}
	return "<unknown>"
}
func (code StructureMapContextType) Definition() string {
	switch code {
	case StructureMapContextTypeType:
		return "The context specifies a type."
	case StructureMapContextTypeVariable:
		return "The context specifies a variable."
	}
	return "<unknown>"
}
