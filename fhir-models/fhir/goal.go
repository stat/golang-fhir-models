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

// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
	Id                   *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	LifecycleStatus      GoalLifecycleStatus `bson:"lifecycleStatus" json:"lifecycleStatus"`
	AchievementStatus    *CodeableConcept    `bson:"achievementStatus,omitempty" json:"achievementStatus,omitempty"`
	Category             []CodeableConcept   `bson:"category,omitempty" json:"category,omitempty"`
	Priority             *CodeableConcept    `bson:"priority,omitempty" json:"priority,omitempty"`
	Description          CodeableConcept     `bson:"description" json:"description"`
	Subject              Reference           `bson:"subject" json:"subject"`
	StartDate            *string             `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept    `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty"`
	Target               []GoalTarget        `bson:"target,omitempty" json:"target,omitempty"`
	StatusDate           *string             `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	StatusReason         *string             `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	ExpressedBy          *Reference          `bson:"expressedBy,omitempty" json:"expressedBy,omitempty"`
	Addresses            []Reference         `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Note                 []Annotation        `bson:"note,omitempty" json:"note,omitempty"`
	OutcomeCode          []CodeableConcept   `bson:"outcomeCode,omitempty" json:"outcomeCode,omitempty"`
	OutcomeReference     []Reference         `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
}
type GoalTarget struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty"`
	DetailString          *string          `bson:"detailString,omitempty" json:"detailString,omitempty"`
	DetailBoolean         *bool            `bson:"detailBoolean,omitempty" json:"detailBoolean,omitempty"`
	DetailInteger         *int             `bson:"detailInteger,omitempty" json:"detailInteger,omitempty"`
	DetailRatio           *Ratio           `bson:"detailRatio,omitempty" json:"detailRatio,omitempty"`
	DueDate               *string          `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
	DueDuration           *Duration        `bson:"dueDuration,omitempty" json:"dueDuration,omitempty"`
}
type OtherGoal Goal

// MarshalJSON marshals the given Goal as JSON into a byte slice
 func (r *Goal)MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}

// UnmarshalGoal unmarshals a Goal.
func UnmarshalGoal(b []byte) (Goal, error) {
	var goal Goal
	if err := json.Unmarshal(b, &goal); err != nil {
		return goal, err
	}
	return goal, nil
}
