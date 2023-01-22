/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// Class5GVNGroupsInternalDocumentApiService Class5GVNGroupsInternalDocumentApi service
type Class5GVNGroupsInternalDocumentApiService service

type ApiQuery5GVnGroupInternalRequest struct {
	ctx context.Context
	ApiService *Class5GVNGroupsInternalDocumentApiService
	internalGroupIds *[]string
}

// List of Internal Group IDs
func (r ApiQuery5GVnGroupInternalRequest) InternalGroupIds(internalGroupIds []string) ApiQuery5GVnGroupInternalRequest {
	r.internalGroupIds = &internalGroupIds
	return r
}

func (r ApiQuery5GVnGroupInternalRequest) Execute() (*map[string]Model5GVnGroupConfiguration, *http.Response, error) {
	return r.ApiService.Query5GVnGroupInternalExecute(r)
}

/*
Query5GVnGroupInternal Retrieves the data of 5G VN Group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQuery5GVnGroupInternalRequest
*/
func (a *Class5GVNGroupsInternalDocumentApiService) Query5GVnGroupInternal(ctx context.Context) ApiQuery5GVnGroupInternalRequest {
	return ApiQuery5GVnGroupInternalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]Model5GVnGroupConfiguration
func (a *Class5GVNGroupsInternalDocumentApiService) Query5GVnGroupInternalExecute(r ApiQuery5GVnGroupInternalRequest) (*map[string]Model5GVnGroupConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]Model5GVnGroupConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class5GVNGroupsInternalDocumentApiService.Query5GVnGroupInternal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/5g-vn-groups/internal"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.internalGroupIds == nil {
		return localVarReturnValue, nil, reportError("internalGroupIds is required and must be specified")
	}

	localVarQueryParams.Add("internal-group-ids", parameterToString(*r.internalGroupIds, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
