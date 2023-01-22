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


// SessionManagementSubscriptionDataApiService SessionManagementSubscriptionDataApi service
type SessionManagementSubscriptionDataApiService service

type ApiQuerySmDataRequest struct {
	ctx context.Context
	ApiService *SessionManagementSubscriptionDataApiService
	ueId string
	servingPlmnId string
	singleNssai *VarSnssai
	dnn *string
	fields *[]string
	supportedFeatures *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// single NSSAI
func (r ApiQuerySmDataRequest) SingleNssai(singleNssai VarSnssai) ApiQuerySmDataRequest {
	r.singleNssai = &singleNssai
	return r
}

// DNN
func (r ApiQuerySmDataRequest) Dnn(dnn string) ApiQuerySmDataRequest {
	r.dnn = &dnn
	return r
}

// attributes to be retrieved
func (r ApiQuerySmDataRequest) Fields(fields []string) ApiQuerySmDataRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQuerySmDataRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQuerySmDataRequest) IfNoneMatch(ifNoneMatch string) ApiQuerySmDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQuerySmDataRequest) IfModifiedSince(ifModifiedSince string) ApiQuerySmDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQuerySmDataRequest) Execute() (*SmSubsData, *http.Response, error) {
	return r.ApiService.QuerySmDataExecute(r)
}

/*
QuerySmData Retrieves the Session Management subscription data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param servingPlmnId PLMN ID
 @return ApiQuerySmDataRequest
*/
func (a *SessionManagementSubscriptionDataApiService) QuerySmData(ctx context.Context, ueId string, servingPlmnId string) ApiQuerySmDataRequest {
	return ApiQuerySmDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		servingPlmnId: servingPlmnId,
	}
}

// Execute executes the request
//  @return SmSubsData
func (a *SessionManagementSubscriptionDataApiService) QuerySmDataExecute(r ApiQuerySmDataRequest) (*SmSubsData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmSubsData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionManagementSubscriptionDataApiService.QuerySmData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/{servingPlmnId}/provisioned-data/sm-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingPlmnId"+"}", url.PathEscape(parameterToString(r.servingPlmnId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.singleNssai != nil {
		localVarQueryParams.Add("single-nssai", parameterToString(*r.singleNssai, ""))
	}
	if r.dnn != nil {
		localVarQueryParams.Add("dnn", parameterToString(*r.dnn, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, "csv"))
	}
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
