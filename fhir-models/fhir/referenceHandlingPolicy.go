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

// ReferenceHandlingPolicy is documented here http://hl7.org/fhir/ValueSet/reference-handling-policy
type ReferenceHandlingPolicy int

const (
	ReferenceHandlingPolicyLiteral ReferenceHandlingPolicy = iota
	ReferenceHandlingPolicyLogical
	ReferenceHandlingPolicyResolves
	ReferenceHandlingPolicyEnforced
	ReferenceHandlingPolicyLocal
)

 func (code *ReferenceHandlingPolicy)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ReferenceHandlingPolicy) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "literal":
		*code = ReferenceHandlingPolicyLiteral
	case "logical":
		*code = ReferenceHandlingPolicyLogical
	case "resolves":
		*code = ReferenceHandlingPolicyResolves
	case "enforced":
		*code = ReferenceHandlingPolicyEnforced
	case "local":
		*code = ReferenceHandlingPolicyLocal
	default:
		return fmt.Errorf("unknown ReferenceHandlingPolicy code `%s`", s)
	}
	return nil
}
func (code ReferenceHandlingPolicy) String() string {
	return code.Code()
}
func (code ReferenceHandlingPolicy) Code() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "literal"
	case ReferenceHandlingPolicyLogical:
		return "logical"
	case ReferenceHandlingPolicyResolves:
		return "resolves"
	case ReferenceHandlingPolicyEnforced:
		return "enforced"
	case ReferenceHandlingPolicyLocal:
		return "local"
	}
	return "<unknown>"
}
func (code ReferenceHandlingPolicy) Display() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "Literal References"
	case ReferenceHandlingPolicyLogical:
		return "Logical References"
	case ReferenceHandlingPolicyResolves:
		return "Resolves References"
	case ReferenceHandlingPolicyEnforced:
		return "Reference Integrity Enforced"
	case ReferenceHandlingPolicyLocal:
		return "Local References Only"
	}
	return "<unknown>"
}
func (code ReferenceHandlingPolicy) Definition() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "The server supports and populates Literal references (i.e. using Reference.reference) where they are known (this code does not guarantee that all references are literal; see 'enforced')."
	case ReferenceHandlingPolicyLogical:
		return "The server allows logical references (i.e. using Reference.identifier)."
	case ReferenceHandlingPolicyResolves:
		return "The server will attempt to resolve logical references to literal references - i.e. converting Reference.identifier to Reference.reference (if resolution fails, the server may still accept resources; see logical)."
	case ReferenceHandlingPolicyEnforced:
		return "The server enforces that references have integrity - e.g. it ensures that references can always be resolved. This is typically the case for clinical record systems, but often not the case for middleware/proxy systems."
	case ReferenceHandlingPolicyLocal:
		return "The server does not support references that point to other servers."
	}
	return "<unknown>"
}
