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

// Contains the necessary attributes to support KMS encryption for Azure based clusters.
type AzureKmsEncryptionBuilder struct {
	fieldSet_ []bool
	activeKey *AzureKmsKeyBuilder
}

// NewAzureKmsEncryption creates a new builder of 'azure_kms_encryption' objects.
func NewAzureKmsEncryption() *AzureKmsEncryptionBuilder {
	return &AzureKmsEncryptionBuilder{
		fieldSet_: make([]bool, 1),
	}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureKmsEncryptionBuilder) Empty() bool {
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

// ActiveKey sets the value of the 'active_key' attribute to the given value.
//
// Contains the necessary attributes to support KMS encryption key for Azure based clusters
func (b *AzureKmsEncryptionBuilder) ActiveKey(value *AzureKmsKeyBuilder) *AzureKmsEncryptionBuilder {
	b.activeKey = value
	if value != nil {
		b.fieldSet_[0] = true
	} else {
		b.fieldSet_[0] = false
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureKmsEncryptionBuilder) Copy(object *AzureKmsEncryption) *AzureKmsEncryptionBuilder {
	if object == nil {
		return b
	}
	if len(object.fieldSet_) > 0 {
		b.fieldSet_ = make([]bool, len(object.fieldSet_))
		copy(b.fieldSet_, object.fieldSet_)
	}
	if object.activeKey != nil {
		b.activeKey = NewAzureKmsKey().Copy(object.activeKey)
	} else {
		b.activeKey = nil
	}
	return b
}

// Build creates a 'azure_kms_encryption' object using the configuration stored in the builder.
func (b *AzureKmsEncryptionBuilder) Build() (object *AzureKmsEncryption, err error) {
	object = new(AzureKmsEncryption)
	if len(b.fieldSet_) > 0 {
		object.fieldSet_ = make([]bool, len(b.fieldSet_))
		copy(object.fieldSet_, b.fieldSet_)
	}
	if b.activeKey != nil {
		object.activeKey, err = b.activeKey.Build()
		if err != nil {
			return
		}
	}
	return
}
