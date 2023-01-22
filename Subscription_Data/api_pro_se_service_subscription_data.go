/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ProSeServiceSubscriptionDataApiService ProSeServiceSubscriptionDataApi service
type ProSeServiceSubscriptionDataApiService service

type ApiQueryPorseDataRequest struct {
	ctx context.Context
	ApiService *ProSeServiceSubscriptionDataApiService
	ueId string
	supportedFeatures *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// Supported Features
func (r ApiQueryPorseDataRequest) SupportedFeatures(supportedFeatures string) ApiQueryPorseDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQueryPorseDataRequest) IfNoneMatch(ifNoneMatch string) ApiQueryPorseDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQueryPorseDataRequest) IfModifiedSince(ifModifiedSince string) ApiQueryPorseDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQueryPorseDataRequest) Execute() (*ProseSubscriptionData, *http.Response, error) {
	return r.ApiService.QueryPorseDataExecute(r)
}

/*
QueryPorseData Retrieves the subscribed ProSe service Data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiQueryPorseDataRequest
*/
func (a *ProSeServiceSubscriptionDataApiService) QueryPorseData(ctx context.Context, ueId string) ApiQueryPorseDataRequest {
	return ApiQueryPorseDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return ProseSubscriptionData
func (a *ProSeServiceSubscriptionDataApiService) QueryPorseDataExecute(r ApiQueryPorseDataRequest) (*ProseSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProseSubscriptionData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProSeServiceSubscriptionDataApiService.QueryPorseData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/prose-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
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
	if r.ifNoneMatch != nil {
		localVarHeaderParams["If-None-Match"] = parameterToString(*r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = parameterToString(*r.ifModifiedSince, "")
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