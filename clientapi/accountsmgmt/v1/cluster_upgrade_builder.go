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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/accountsmgmt/v1

import (
	time "time"
)

type ClusterUpgradeBuilder struct {
	fieldSet_        []bool
	state            string
	updatedTimestamp time.Time
	version          string
	available        bool
}

// NewClusterUpgrade creates a new builder of 'cluster_upgrade' objects.
func NewClusterUpgrade() *ClusterUpgradeBuilder {
	return &ClusterUpgradeBuilder{
		fieldSet_: make([]bool, 4),
	}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ClusterUpgradeBuilder) Empty() bool {
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

// Available sets the value of the 'available' attribute to the given value.
func (b *ClusterUpgradeBuilder) Available(value bool) *ClusterUpgradeBuilder {
	b.available = value
	b.fieldSet_[0] = true
	return b
}

// State sets the value of the 'state' attribute to the given value.
func (b *ClusterUpgradeBuilder) State(value string) *ClusterUpgradeBuilder {
	b.state = value
	b.fieldSet_[1] = true
	return b
}

// UpdatedTimestamp sets the value of the 'updated_timestamp' attribute to the given value.
func (b *ClusterUpgradeBuilder) UpdatedTimestamp(value time.Time) *ClusterUpgradeBuilder {
	b.updatedTimestamp = value
	b.fieldSet_[2] = true
	return b
}

// Version sets the value of the 'version' attribute to the given value.
func (b *ClusterUpgradeBuilder) Version(value string) *ClusterUpgradeBuilder {
	b.version = value
	b.fieldSet_[3] = true
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterUpgradeBuilder) Copy(object *ClusterUpgrade) *ClusterUpgradeBuilder {
	if object == nil {
		return b
	}
	if len(object.fieldSet_) > 0 {
		b.fieldSet_ = make([]bool, len(object.fieldSet_))
		copy(b.fieldSet_, object.fieldSet_)
	}
	b.available = object.available
	b.state = object.state
	b.updatedTimestamp = object.updatedTimestamp
	b.version = object.version
	return b
}

// Build creates a 'cluster_upgrade' object using the configuration stored in the builder.
func (b *ClusterUpgradeBuilder) Build() (object *ClusterUpgrade, err error) {
	object = new(ClusterUpgrade)
	if len(b.fieldSet_) > 0 {
		object.fieldSet_ = make([]bool, len(b.fieldSet_))
		copy(object.fieldSet_, b.fieldSet_)
	}
	object.available = b.available
	object.state = b.state
	object.updatedTimestamp = b.updatedTimestamp
	object.version = b.version
	return
}
