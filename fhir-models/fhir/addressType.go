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
	bson "go.mongodb.org/mongo-driver/bson"
	bsonrw "go.mongodb.org/mongo-driver/bson/bsonrw"
	bsontype "go.mongodb.org/mongo-driver/bson/bsontype"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AddressType is documented here http://hl7.org/fhir/ValueSet/address-type
type AddressType int

const (
	AddressTypePostal AddressType = iota
	AddressTypePhysical
	AddressTypeBoth
)

func (code *AddressType) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(code.Code())
}
func (code *AddressType) UnmarshalBSONValue(t bsontype.Type, bytes []byte) error {
	if t != bsontype.String {
		err := fmt.Errorf("UnmarshalBSONValue error: cannot unmarshal non string value")
		return err
	}
	reader := bsonrw.NewBSONValueReader(t, bytes)
	decoder, err := bson.NewDecoder(reader)
	if err != nil {
		return err
	}
	var s string
	err = decoder.Decode(&s)
	if err != nil {
		return err
	}
	switch s {
	case "postal":
		*code = AddressTypePostal
	case "physical":
		*code = AddressTypePhysical
	case "both":
		*code = AddressTypeBoth
	default:
		return fmt.Errorf("unknown AddressType code `%s`", s)
	}
	return nil
}
func (code *AddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AddressType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "postal":
		*code = AddressTypePostal
	case "physical":
		*code = AddressTypePhysical
	case "both":
		*code = AddressTypeBoth
	default:
		return fmt.Errorf("unknown AddressType code `%s`", s)
	}
	return nil
}
func (code AddressType) String() string {
	return code.Code()
}
func (code AddressType) Code() string {
	switch code {
	case AddressTypePostal:
		return "postal"
	case AddressTypePhysical:
		return "physical"
	case AddressTypeBoth:
		return "both"
	}
	return "<unknown>"
}
func (code AddressType) Display() string {
	switch code {
	case AddressTypePostal:
		return "Postal"
	case AddressTypePhysical:
		return "Physical"
	case AddressTypeBoth:
		return "Postal & Physical"
	}
	return "<unknown>"
}
func (code AddressType) Definition() string {
	switch code {
	case AddressTypePostal:
		return "Mailing addresses - PO Boxes and care-of addresses."
	case AddressTypePhysical:
		return "A physical address that can be visited."
	case AddressTypeBoth:
		return "An address that is both physical and postal."
	}
	return "<unknown>"
}
