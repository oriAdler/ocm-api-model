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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/jobqueue/v1

import (
	time "time"
)

type QueueBuilder struct {
	fieldSet_   []bool
	id          string
	href        string
	createdAt   time.Time
	maxAttempts int
	maxRunTime  int
	name        string
	updatedAt   time.Time
}

// NewQueue creates a new builder of 'queue' objects.
func NewQueue() *QueueBuilder {
	return &QueueBuilder{
		fieldSet_: make([]bool, 8),
	}
}

// Link sets the flag that indicates if this is a link.
func (b *QueueBuilder) Link(value bool) *QueueBuilder {
	b.fieldSet_[0] = true
	return b
}

// ID sets the identifier of the object.
func (b *QueueBuilder) ID(value string) *QueueBuilder {
	b.id = value
	b.fieldSet_[1] = true
	return b
}

// HREF sets the link to the object.
func (b *QueueBuilder) HREF(value string) *QueueBuilder {
	b.href = value
	b.fieldSet_[2] = true
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *QueueBuilder) Empty() bool {
	if b == nil || len(b.fieldSet_) == 0 {
		return true
	}
	// Check all fields except the link flag (index 0)
	for i := 1; i < len(b.fieldSet_); i++ {
		if b.fieldSet_[i] {
			return false
		}
	}
	return true
}

// CreatedAt sets the value of the 'created_at' attribute to the given value.
func (b *QueueBuilder) CreatedAt(value time.Time) *QueueBuilder {
	b.createdAt = value
	b.fieldSet_[3] = true
	return b
}

// MaxAttempts sets the value of the 'max_attempts' attribute to the given value.
func (b *QueueBuilder) MaxAttempts(value int) *QueueBuilder {
	b.maxAttempts = value
	b.fieldSet_[4] = true
	return b
}

// MaxRunTime sets the value of the 'max_run_time' attribute to the given value.
func (b *QueueBuilder) MaxRunTime(value int) *QueueBuilder {
	b.maxRunTime = value
	b.fieldSet_[5] = true
	return b
}

// Name sets the value of the 'name' attribute to the given value.
func (b *QueueBuilder) Name(value string) *QueueBuilder {
	b.name = value
	b.fieldSet_[6] = true
	return b
}

// UpdatedAt sets the value of the 'updated_at' attribute to the given value.
func (b *QueueBuilder) UpdatedAt(value time.Time) *QueueBuilder {
	b.updatedAt = value
	b.fieldSet_[7] = true
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *QueueBuilder) Copy(object *Queue) *QueueBuilder {
	if object == nil {
		return b
	}
	if len(object.fieldSet_) > 0 {
		b.fieldSet_ = make([]bool, len(object.fieldSet_))
		copy(b.fieldSet_, object.fieldSet_)
	}
	b.id = object.id
	b.href = object.href
	b.createdAt = object.createdAt
	b.maxAttempts = object.maxAttempts
	b.maxRunTime = object.maxRunTime
	b.name = object.name
	b.updatedAt = object.updatedAt
	return b
}

// Build creates a 'queue' object using the configuration stored in the builder.
func (b *QueueBuilder) Build() (object *Queue, err error) {
	object = new(Queue)
	object.id = b.id
	object.href = b.href
	if len(b.fieldSet_) > 0 {
		object.fieldSet_ = make([]bool, len(b.fieldSet_))
		copy(object.fieldSet_, b.fieldSet_)
	}
	object.createdAt = b.createdAt
	object.maxAttempts = b.maxAttempts
	object.maxRunTime = b.maxRunTime
	object.name = b.name
	object.updatedAt = b.updatedAt
	return
}
