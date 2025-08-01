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

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalFeatureReviewRequest writes a value of the 'feature_review_request' type to the given writer.
func MarshalFeatureReviewRequest(object *FeatureReviewRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteFeatureReviewRequest(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteFeatureReviewRequest writes a value of the 'feature_review_request' type to the given stream.
func WriteFeatureReviewRequest(object *FeatureReviewRequest, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("account_username")
		stream.WriteString(object.accountUsername)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_id")
		stream.WriteString(object.clusterId)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("feature")
		stream.WriteString(object.feature)
		count++
	}
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("organization_id")
		stream.WriteString(object.organizationId)
	}
	stream.WriteObjectEnd()
}

// UnmarshalFeatureReviewRequest reads a value of the 'feature_review_request' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalFeatureReviewRequest(source interface{}) (object *FeatureReviewRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadFeatureReviewRequest(iterator)
	err = iterator.Error
	return
}

// ReadFeatureReviewRequest reads a value of the 'feature_review_request' type from the given iterator.
func ReadFeatureReviewRequest(iterator *jsoniter.Iterator) *FeatureReviewRequest {
	object := &FeatureReviewRequest{
		fieldSet_: make([]bool, 4),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "account_username":
			value := iterator.ReadString()
			object.accountUsername = value
			object.fieldSet_[0] = true
		case "cluster_id":
			value := iterator.ReadString()
			object.clusterId = value
			object.fieldSet_[1] = true
		case "feature":
			value := iterator.ReadString()
			object.feature = value
			object.fieldSet_[2] = true
		case "organization_id":
			value := iterator.ReadString()
			object.organizationId = value
			object.fieldSet_[3] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
