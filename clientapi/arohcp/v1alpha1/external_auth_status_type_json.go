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

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalExternalAuthStatus writes a value of the 'external_auth_status' type to the given writer.
func MarshalExternalAuthStatus(object *ExternalAuthStatus, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteExternalAuthStatus(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteExternalAuthStatus writes a value of the 'external_auth_status' type to the given stream.
func WriteExternalAuthStatus(object *ExternalAuthStatus, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("message")
		stream.WriteString(object.message)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1] && object.state != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("state")
		WriteExternalAuthState(object.state, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalExternalAuthStatus reads a value of the 'external_auth_status' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalExternalAuthStatus(source interface{}) (object *ExternalAuthStatus, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadExternalAuthStatus(iterator)
	err = iterator.Error
	return
}

// ReadExternalAuthStatus reads a value of the 'external_auth_status' type from the given iterator.
func ReadExternalAuthStatus(iterator *jsoniter.Iterator) *ExternalAuthStatus {
	object := &ExternalAuthStatus{
		fieldSet_: make([]bool, 2),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "message":
			value := iterator.ReadString()
			object.message = value
			object.fieldSet_[0] = true
		case "state":
			value := ReadExternalAuthState(iterator)
			object.state = value
			object.fieldSet_[1] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
