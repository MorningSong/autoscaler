// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetComputeGlobalImageCapabilitySchemaVersionRequest wrapper for the GetComputeGlobalImageCapabilitySchemaVersion operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetComputeGlobalImageCapabilitySchemaVersion.go.html to see an example of how to use GetComputeGlobalImageCapabilitySchemaVersionRequest.
type GetComputeGlobalImageCapabilitySchemaVersionRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute global image capability schema
	ComputeGlobalImageCapabilitySchemaId *string `mandatory:"true" contributesTo:"path" name:"computeGlobalImageCapabilitySchemaId"`

	// The name of the compute global image capability schema version
	ComputeGlobalImageCapabilitySchemaVersionName *string `mandatory:"true" contributesTo:"path" name:"computeGlobalImageCapabilitySchemaVersionName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetComputeGlobalImageCapabilitySchemaVersionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetComputeGlobalImageCapabilitySchemaVersionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetComputeGlobalImageCapabilitySchemaVersionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetComputeGlobalImageCapabilitySchemaVersionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetComputeGlobalImageCapabilitySchemaVersionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetComputeGlobalImageCapabilitySchemaVersionResponse wrapper for the GetComputeGlobalImageCapabilitySchemaVersion operation
type GetComputeGlobalImageCapabilitySchemaVersionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ComputeGlobalImageCapabilitySchemaVersion instance
	ComputeGlobalImageCapabilitySchemaVersion `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetComputeGlobalImageCapabilitySchemaVersionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetComputeGlobalImageCapabilitySchemaVersionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
