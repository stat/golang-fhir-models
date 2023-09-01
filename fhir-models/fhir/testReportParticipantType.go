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

// TestReportParticipantType is documented here http://hl7.org/fhir/ValueSet/report-participant-type
type TestReportParticipantType int

const (
	TestReportParticipantTypeTestEngine TestReportParticipantType = iota
	TestReportParticipantTypeClient
	TestReportParticipantTypeServer
)

 func (code *TestReportParticipantType)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TestReportParticipantType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "test-engine":
		*code = TestReportParticipantTypeTestEngine
	case "client":
		*code = TestReportParticipantTypeClient
	case "server":
		*code = TestReportParticipantTypeServer
	default:
		return fmt.Errorf("unknown TestReportParticipantType code `%s`", s)
	}
	return nil
}
func (code TestReportParticipantType) String() string {
	return code.Code()
}
func (code TestReportParticipantType) Code() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "test-engine"
	case TestReportParticipantTypeClient:
		return "client"
	case TestReportParticipantTypeServer:
		return "server"
	}
	return "<unknown>"
}
func (code TestReportParticipantType) Display() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "Test Engine"
	case TestReportParticipantTypeClient:
		return "Client"
	case TestReportParticipantTypeServer:
		return "Server"
	}
	return "<unknown>"
}
func (code TestReportParticipantType) Definition() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "The test execution engine."
	case TestReportParticipantTypeClient:
		return "A FHIR Client."
	case TestReportParticipantTypeServer:
		return "A FHIR Server."
	}
	return "<unknown>"
}
