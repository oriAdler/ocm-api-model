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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalAzureNodesOutboundConnectivity writes a value of the 'azure_nodes_outbound_connectivity' type to the given writer.
func MarshalAzureNodesOutboundConnectivity(object *AzureNodesOutboundConnectivity, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteAzureNodesOutboundConnectivity(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteAzureNodesOutboundConnectivity writes a value of the 'azure_nodes_outbound_connectivity' type to the given stream.
func WriteAzureNodesOutboundConnectivity(object *AzureNodesOutboundConnectivity, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("outbound_type")
		stream.WriteString(object.outboundType)
	}
	stream.WriteObjectEnd()
}

// UnmarshalAzureNodesOutboundConnectivity reads a value of the 'azure_nodes_outbound_connectivity' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAzureNodesOutboundConnectivity(source interface{}) (object *AzureNodesOutboundConnectivity, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadAzureNodesOutboundConnectivity(iterator)
	err = iterator.Error
	return
}

// ReadAzureNodesOutboundConnectivity reads a value of the 'azure_nodes_outbound_connectivity' type from the given iterator.
func ReadAzureNodesOutboundConnectivity(iterator *jsoniter.Iterator) *AzureNodesOutboundConnectivity {
	object := &AzureNodesOutboundConnectivity{
		fieldSet_: make([]bool, 1),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "outbound_type":
			value := iterator.ReadString()
			object.outboundType = value
			object.fieldSet_[0] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
