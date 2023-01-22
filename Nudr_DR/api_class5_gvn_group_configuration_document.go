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


// Class5GVnGroupConfigurationDocumentApiService Class5GVnGroupConfigurationDocumentApi service
type Class5GVnGroupConfigurationDocumentApiService service

type ApiCreate5GVnGroupRequest struct {
	ctx context.Context
	ApiService *Class5GVnGroupConfigurationDocumentApiService
	externalGroupId string
	model5GVnGroupConfiguration *Model5GVnGroupConfiguration
}

func (r ApiCreate5GVnGroupRequest) Model5GVnGroupConfiguration(model5GVnGroupConfiguration Model5GVnGroupConfiguration) ApiCreate5GVnGroupRequest {
	r.model5GVnGroupConfiguration = &model5GVnGroupConfiguration
	return r
}

func (r ApiCreate5GVnGroupRequest) Execute() (*Model5GVnGroupConfiguration, *http.Response, error) {
	return r.ApiService.Create5GVnGroupExecute(r)
}

/*
Create5GVnGroup Create an individual 5G VN Grouop

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGroupId
 @return ApiCreate5GVnGroupRequest
*/
func (a *Class5GVnGroupConfigurationDocumentApiService) Create5GVnGroup(ctx context.Context, externalGroupId string) ApiCreate5GVnGroupRequest {
	return ApiCreate5GVnGroupRequest{
		ApiService: a,
		ctx: ctx,
		externalGroupId: externalGroupId,
	}
}

// Execute executes the request
//  @return Model5GVnGroupConfiguration
func (a *Class5GVnGroupConfigurationDocumentApiService) Create5GVnGroupExecute(r ApiCreate5GVnGroupRequest) (*Model5GVnGroupConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Model5GVnGroupConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class5GVnGroupConfigurationDocumentApiService.Create5GVnGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/5g-vn-groups/{externalGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterToString(r.externalGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model5GVnGroupConfiguration == nil {
		return localVarReturnValue, nil, reportError("model5GVnGroupConfiguration is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.model5GVnGroupConfiguration
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
