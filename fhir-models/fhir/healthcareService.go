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

// HealthcareService is documented here http://hl7.org/fhir/StructureDefinition/HealthcareService
type HealthcareService struct {
	Id                     *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                 *bool                            `bson:"active,omitempty" json:"active,omitempty"`
	ProvidedBy             *Reference                       `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	Category               []CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Type                   []CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Specialty              []CodeableConcept                `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	Name                   *string                          `bson:"name,omitempty" json:"name,omitempty"`
	Comment                *string                          `bson:"comment,omitempty" json:"comment,omitempty"`
	ExtraDetails           *string                          `bson:"extraDetails,omitempty" json:"extraDetails,omitempty"`
	Photo                  *Attachment                      `bson:"photo,omitempty" json:"photo,omitempty"`
	Telecom                []ContactPoint                   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	CoverageArea           []Reference                      `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                `bson:"serviceProvisionCode,omitempty" json:"serviceProvisionCode,omitempty"`
	Eligibility            []HealthcareServiceEligibility   `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	Program                []CodeableConcept                `bson:"program,omitempty" json:"program,omitempty"`
	Characteristic         []CodeableConcept                `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Communication          []CodeableConcept                `bson:"communication,omitempty" json:"communication,omitempty"`
	ReferralMethod         []CodeableConcept                `bson:"referralMethod,omitempty" json:"referralMethod,omitempty"`
	AppointmentRequired    *bool                            `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTime `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailable  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                          `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                      `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type HealthcareServiceEligibility struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Comment           *string          `bson:"comment,omitempty" json:"comment,omitempty"`
}
type HealthcareServiceAvailableTime struct {
	Id                 *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DaysOfWeek         []DaysOfWeek `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool        `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *string      `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *string      `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}
type HealthcareServiceNotAvailable struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string      `bson:"description" json:"description"`
	During            *Period     `bson:"during,omitempty" json:"during,omitempty"`
}
type OtherHealthcareService HealthcareService

// MarshalJSON marshals the given HealthcareService as JSON into a byte slice
 func (r *HealthcareService)MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherHealthcareService
		ResourceType string `json:"resourceType"`
	}{
		OtherHealthcareService: OtherHealthcareService(r),
		ResourceType:           "HealthcareService",
	})
}

// UnmarshalHealthcareService unmarshals a HealthcareService.
func UnmarshalHealthcareService(b []byte) (HealthcareService, error) {
	var healthcareService HealthcareService
	if err := json.Unmarshal(b, &healthcareService); err != nil {
		return healthcareService, err
	}
	return healthcareService, nil
}
