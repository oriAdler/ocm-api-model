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

// AddOnInstallMode represents the values of the 'add_on_install_mode' enumerated type.
type AddOnInstallMode string

const (
	// This mode means that the addon is deployed in all namespaces.
	// However, the addon status is retrieved from the target namespace
	AddOnInstallModeAllNamespaces AddOnInstallMode = "all_namespaces"
	// This mode means that the the addon CRD exists in a single specific namespace.
	// This namespace is reflected by the TargetNamespace addon field
	AddOnInstallModeOwnNamespace AddOnInstallMode = "own_namespace"
)
