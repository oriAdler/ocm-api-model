/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// The rule that is applied to validate token claims to authenticate users.
type TokenClaimValidationRuleBuilder struct {
	fieldSet_     []bool
	claim         string
	requiredValue string
}

// NewTokenClaimValidationRule creates a new builder of 'token_claim_validation_rule' objects.
func NewTokenClaimValidationRule() *TokenClaimValidationRuleBuilder {
	return &TokenClaimValidationRuleBuilder{
		fieldSet_: make([]bool, 2),
	}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *TokenClaimValidationRuleBuilder) Empty() bool {
	if b == nil || len(b.fieldSet_) == 0 {
		return true
	}
	for _, set := range b.fieldSet_ {
		if set {
			return false
		}
	}
	return true
}

// Claim sets the value of the 'claim' attribute to the given value.
func (b *TokenClaimValidationRuleBuilder) Claim(value string) *TokenClaimValidationRuleBuilder {
	b.claim = value
	b.fieldSet_[0] = true
	return b
}

// RequiredValue sets the value of the 'required_value' attribute to the given value.
func (b *TokenClaimValidationRuleBuilder) RequiredValue(value string) *TokenClaimValidationRuleBuilder {
	b.requiredValue = value
	b.fieldSet_[1] = true
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *TokenClaimValidationRuleBuilder) Copy(object *TokenClaimValidationRule) *TokenClaimValidationRuleBuilder {
	if object == nil {
		return b
	}
	if len(object.fieldSet_) > 0 {
		b.fieldSet_ = make([]bool, len(object.fieldSet_))
		copy(b.fieldSet_, object.fieldSet_)
	}
	b.claim = object.claim
	b.requiredValue = object.requiredValue
	return b
}

// Build creates a 'token_claim_validation_rule' object using the configuration stored in the builder.
func (b *TokenClaimValidationRuleBuilder) Build() (object *TokenClaimValidationRule, err error) {
	object = new(TokenClaimValidationRule)
	if len(b.fieldSet_) > 0 {
		object.fieldSet_ = make([]bool, len(b.fieldSet_))
		copy(object.fieldSet_, b.fieldSet_)
	}
	object.claim = b.claim
	object.requiredValue = b.requiredValue
	return
}
