// Copyright 2006-2021 VMware, Inc.
// SPDX-License-Identifier: MIT

// Package v1alpha2 contains API Schema definitions for the tenancy v1alpha2 API group
//+kubebuilder:object:generate=true
//+groupName=tenancy.platform.cnr.vmware.com
package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "tenancy.platform.cnr.vmware.com", Version: "v1alpha2"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
