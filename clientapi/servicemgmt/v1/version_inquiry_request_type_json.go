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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/servicemgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalVersionInquiryRequest writes a value of the 'version_inquiry_request' type to the given writer.
func MarshalVersionInquiryRequest(object *VersionInquiryRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteVersionInquiryRequest(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteVersionInquiryRequest writes a value of the 'version_inquiry_request' type to the given stream.
func WriteVersionInquiryRequest(object *VersionInquiryRequest, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("service_type")
		stream.WriteString(object.serviceType)
	}
	stream.WriteObjectEnd()
}

// UnmarshalVersionInquiryRequest reads a value of the 'version_inquiry_request' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalVersionInquiryRequest(source interface{}) (object *VersionInquiryRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadVersionInquiryRequest(iterator)
	err = iterator.Error
	return
}

// ReadVersionInquiryRequest reads a value of the 'version_inquiry_request' type from the given iterator.
func ReadVersionInquiryRequest(iterator *jsoniter.Iterator) *VersionInquiryRequest {
	object := &VersionInquiryRequest{
		fieldSet_: make([]bool, 1),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "service_type":
			value := iterator.ReadString()
			object.serviceType = value
			object.fieldSet_[0] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
