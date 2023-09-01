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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/stat/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Communication is documented here http://hl7.org/fhir/StructureDefinition/Communication
type Communication struct {
	Id                    *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string               `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                []Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	InResponseTo          []Reference            `bson:"inResponseTo,omitempty" json:"inResponseTo,omitempty"`
	Status                EventStatus            `bson:"status" json:"status"`
	StatusReason          *CodeableConcept       `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category              []CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Priority              *RequestPriority       `bson:"priority,omitempty" json:"priority,omitempty"`
	Medium                []CodeableConcept      `bson:"medium,omitempty" json:"medium,omitempty"`
	Subject               *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	Topic                 *CodeableConcept       `bson:"topic,omitempty" json:"topic,omitempty"`
	About                 []Reference            `bson:"about,omitempty" json:"about,omitempty"`
	Encounter             *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Sent                  *string                `bson:"sent,omitempty" json:"sent,omitempty"`
	Received              *string                `bson:"received,omitempty" json:"received,omitempty"`
	Recipient             []Reference            `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Sender                *Reference             `bson:"sender,omitempty" json:"sender,omitempty"`
	ReasonCode            []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Payload               []CommunicationPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	Note                  []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
}
type CommunicationPayload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentString     string      `bson:"contentString" json:"contentString"`
	ContentAttachment Attachment  `bson:"contentAttachment" json:"contentAttachment"`
	ContentReference  Reference   `bson:"contentReference" json:"contentReference"`
}
type OtherCommunication Communication

// MarshalJSON marshals the given Communication as JSON into a byte slice
 func (r *Communication)MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunication
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunication: OtherCommunication(r),
		ResourceType:       "Communication",
	})
}

// UnmarshalCommunication unmarshals a Communication.
func UnmarshalCommunication(b []byte) (Communication, error) {
	var communication Communication
	if err := json.Unmarshal(b, &communication); err != nil {
		return communication, err
	}
	return communication, nil
}
