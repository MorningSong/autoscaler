// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttachParavirtualizedVolumeDetails The representation of AttachParavirtualizedVolumeDetails
type AttachParavirtualizedVolumeDetails struct {

	// The OCID of the instance. For AttachVolume operation, this is a required field for the request,
	// see AttachVolume.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The OCID of the volume. If CreateVolumeDetails is specified, this field must be omitted from the request.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// The device name. To retrieve a list of devices for a given instance, see ListInstanceDevices.
	Device *string `mandatory:"false" json:"device"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `mandatory:"false" json:"isShareable"`

	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`
}

// GetDevice returns Device
func (m AttachParavirtualizedVolumeDetails) GetDevice() *string {
	return m.Device
}

// GetDisplayName returns DisplayName
func (m AttachParavirtualizedVolumeDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetInstanceId returns InstanceId
func (m AttachParavirtualizedVolumeDetails) GetInstanceId() *string {
	return m.InstanceId
}

// GetIsReadOnly returns IsReadOnly
func (m AttachParavirtualizedVolumeDetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

// GetIsShareable returns IsShareable
func (m AttachParavirtualizedVolumeDetails) GetIsShareable() *bool {
	return m.IsShareable
}

// GetVolumeId returns VolumeId
func (m AttachParavirtualizedVolumeDetails) GetVolumeId() *string {
	return m.VolumeId
}

func (m AttachParavirtualizedVolumeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttachParavirtualizedVolumeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AttachParavirtualizedVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAttachParavirtualizedVolumeDetails AttachParavirtualizedVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAttachParavirtualizedVolumeDetails
	}{
		"paravirtualized",
		(MarshalTypeAttachParavirtualizedVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
