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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/addonsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalAddonInstallationBilling writes a value of the 'addon_installation_billing' type to the given writer.
func MarshalAddonInstallationBilling(object *AddonInstallationBilling, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteAddonInstallationBilling(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteAddonInstallationBilling writes a value of the 'addon_installation_billing' type to the given stream.
func WriteAddonInstallationBilling(object *AddonInstallationBilling, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("billing_marketplace_account")
		stream.WriteString(object.billingMarketplaceAccount)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("billing_model")
		stream.WriteString(string(object.billingModel))
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	present_ = len(object.fieldSet_) > 4 && object.fieldSet_[4]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("kind")
		stream.WriteString(object.kind)
	}
	stream.WriteObjectEnd()
}

// UnmarshalAddonInstallationBilling reads a value of the 'addon_installation_billing' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAddonInstallationBilling(source interface{}) (object *AddonInstallationBilling, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadAddonInstallationBilling(iterator)
	err = iterator.Error
	return
}

// ReadAddonInstallationBilling reads a value of the 'addon_installation_billing' type from the given iterator.
func ReadAddonInstallationBilling(iterator *jsoniter.Iterator) *AddonInstallationBilling {
	object := &AddonInstallationBilling{
		fieldSet_: make([]bool, 5),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "billing_marketplace_account":
			value := iterator.ReadString()
			object.billingMarketplaceAccount = value
			object.fieldSet_[0] = true
		case "billing_model":
			text := iterator.ReadString()
			value := BillingModel(text)
			object.billingModel = value
			object.fieldSet_[1] = true
		case "href":
			value := iterator.ReadString()
			object.href = value
			object.fieldSet_[2] = true
		case "id":
			value := iterator.ReadString()
			object.id = value
			object.fieldSet_[3] = true
		case "kind":
			value := iterator.ReadString()
			object.kind = value
			object.fieldSet_[4] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
