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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/osdfleetmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalLabelRequestPayload writes a value of the 'label_request_payload' type to the given writer.
func MarshalLabelRequestPayload(object *LabelRequestPayload, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteLabelRequestPayload(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteLabelRequestPayload writes a value of the 'label_request_payload' type to the given stream.
func WriteLabelRequestPayload(object *LabelRequestPayload, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("key")
		stream.WriteString(object.key)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("value")
		stream.WriteString(object.value)
	}
	stream.WriteObjectEnd()
}

// UnmarshalLabelRequestPayload reads a value of the 'label_request_payload' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalLabelRequestPayload(source interface{}) (object *LabelRequestPayload, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadLabelRequestPayload(iterator)
	err = iterator.Error
	return
}

// ReadLabelRequestPayload reads a value of the 'label_request_payload' type from the given iterator.
func ReadLabelRequestPayload(iterator *jsoniter.Iterator) *LabelRequestPayload {
	object := &LabelRequestPayload{
		fieldSet_: make([]bool, 2),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "key":
			value := iterator.ReadString()
			object.key = value
			object.fieldSet_[0] = true
		case "value":
			value := iterator.ReadString()
			object.value = value
			object.fieldSet_[1] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
