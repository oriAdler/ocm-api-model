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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/authorizations/v1

// Representation of Red Hat's Terms and Conditions for using OpenShift Dedicated and Amazon Red Hat OpenShift [Terms]
// review requests.
type SelfTermsReviewRequestBuilder struct {
	fieldSet_ []bool
	eventCode string
	siteCode  string
}

// NewSelfTermsReviewRequest creates a new builder of 'self_terms_review_request' objects.
func NewSelfTermsReviewRequest() *SelfTermsReviewRequestBuilder {
	return &SelfTermsReviewRequestBuilder{
		fieldSet_: make([]bool, 2),
	}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *SelfTermsReviewRequestBuilder) Empty() bool {
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

// EventCode sets the value of the 'event_code' attribute to the given value.
func (b *SelfTermsReviewRequestBuilder) EventCode(value string) *SelfTermsReviewRequestBuilder {
	b.eventCode = value
	b.fieldSet_[0] = true
	return b
}

// SiteCode sets the value of the 'site_code' attribute to the given value.
func (b *SelfTermsReviewRequestBuilder) SiteCode(value string) *SelfTermsReviewRequestBuilder {
	b.siteCode = value
	b.fieldSet_[1] = true
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *SelfTermsReviewRequestBuilder) Copy(object *SelfTermsReviewRequest) *SelfTermsReviewRequestBuilder {
	if object == nil {
		return b
	}
	if len(object.fieldSet_) > 0 {
		b.fieldSet_ = make([]bool, len(object.fieldSet_))
		copy(b.fieldSet_, object.fieldSet_)
	}
	b.eventCode = object.eventCode
	b.siteCode = object.siteCode
	return b
}

// Build creates a 'self_terms_review_request' object using the configuration stored in the builder.
func (b *SelfTermsReviewRequestBuilder) Build() (object *SelfTermsReviewRequest, err error) {
	object = new(SelfTermsReviewRequest)
	if len(b.fieldSet_) > 0 {
		object.fieldSet_ = make([]bool, len(b.fieldSet_))
		copy(object.fieldSet_, b.fieldSet_)
	}
	object.eventCode = b.eventCode
	object.siteCode = b.siteCode
	return
}
