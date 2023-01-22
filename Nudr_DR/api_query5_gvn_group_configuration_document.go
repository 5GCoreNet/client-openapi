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
	"strings"
)


// Query5GVnGroupConfigurationDocumentApiService Query5GVnGroupConfigurationDocumentApi service
type Query5GVnGroupConfigurationDocumentApiService service

type ApiGet5GVnGroupConfigurationRequest struct {
	ctx context.Context
	ApiService *Query5GVnGroupConfigurationDocumentApiService
	externalGroupId string
}

func (r ApiGet5GVnGroupConfigurationRequest) Execute() (*Model5GVnGroupConfiguration, *http.Response, error) {
	return r.ApiService.Get5GVnGroupConfigurationExecute(r)
}

/*
Get5GVnGroupConfiguration Retrieve a 5GVnGroup configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGroupId
 @return ApiGet5GVnGroupConfigurationRequest
*/
func (a *Query5GVnGroupConfigurationDocumentApiService) Get5GVnGroupConfiguration(ctx context.Context, externalGroupId string) ApiGet5GVnGroupConfigurationRequest {
	return ApiGet5GVnGroupConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		externalGroupId: externalGroupId,
	}
}

// Execute executes the request
//  @return Model5GVnGroupConfiguration
func (a *Query5GVnGroupConfigurationDocumentApiService) Get5GVnGroupConfigurationExecute(r ApiGet5GVnGroupConfigurationRequest) (*Model5GVnGroupConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Model5GVnGroupConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Query5GVnGroupConfigurationDocumentApiService.Get5GVnGroupConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/5g-vn-groups/{externalGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterToString(r.externalGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
