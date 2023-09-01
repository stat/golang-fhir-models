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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicinalProductIndication is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductIndication
type MedicinalProductIndication struct {
	Id                      *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                               `bson:"text,omitempty" json:"text,omitempty"`
	Extension               []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Subject                 []Reference                              `bson:"subject,omitempty" json:"subject,omitempty"`
	DiseaseSymptomProcedure *CodeableConcept                         `bson:"diseaseSymptomProcedure,omitempty" json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableConcept                         `bson:"diseaseStatus,omitempty" json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableConcept                        `bson:"comorbidity,omitempty" json:"comorbidity,omitempty"`
	IntendedEffect          *CodeableConcept                         `bson:"intendedEffect,omitempty" json:"intendedEffect,omitempty"`
	Duration                *Quantity                                `bson:"duration,omitempty" json:"duration,omitempty"`
	OtherTherapy            []MedicinalProductIndicationOtherTherapy `bson:"otherTherapy,omitempty" json:"otherTherapy,omitempty"`
	UndesirableEffect       []Reference                              `bson:"undesirableEffect,omitempty" json:"undesirableEffect,omitempty"`
	Population              []Population                             `bson:"population,omitempty" json:"population,omitempty"`
}
type MedicinalProductIndicationOtherTherapy struct {
	Id                        *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept `bson:"therapyRelationshipType" json:"therapyRelationshipType"`
	MedicationCodeableConcept CodeableConcept `bson:"medicationCodeableConcept" json:"medicationCodeableConcept"`
	MedicationReference       Reference       `bson:"medicationReference" json:"medicationReference"`
}
type OtherMedicinalProductIndication MedicinalProductIndication

// MarshalJSON marshals the given MedicinalProductIndication as JSON into a byte slice
 func (r *MedicinalProductIndication)MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductIndication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductIndication: OtherMedicinalProductIndication(r),
		ResourceType:                    "MedicinalProductIndication",
	})
}

// UnmarshalMedicinalProductIndication unmarshals a MedicinalProductIndication.
func UnmarshalMedicinalProductIndication(b []byte) (MedicinalProductIndication, error) {
	var medicinalProductIndication MedicinalProductIndication
	if err := json.Unmarshal(b, &medicinalProductIndication); err != nil {
		return medicinalProductIndication, err
	}
	return medicinalProductIndication, nil
}
