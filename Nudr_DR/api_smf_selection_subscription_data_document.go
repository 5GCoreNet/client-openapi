/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type SMFSelectionSubscriptionDataDocumentApi interface {

	/*
	QuerySmfSelectData Retrieves the SMF selection subscription data of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@param servingPlmnId PLMN ID
	@return ApiQuerySmfSelectDataRequest
	*/
	QuerySmfSelectData(ctx context.Context, ueId string, servingPlmnId string) ApiQuerySmfSelectDataRequest

	// QuerySmfSelectDataExecute executes the request
	//  @return SmfSelectionSubscriptionData
	QuerySmfSelectDataExecute(r ApiQuerySmfSelectDataRequest) (*SmfSelectionSubscriptionData, *http.Response, error)
}

// SMFSelectionSubscriptionDataDocumentApiService SMFSelectionSubscriptionDataDocumentApi service
type SMFSelectionSubscriptionDataDocumentApiService service

type ApiQuerySmfSelectDataRequest struct {
	ctx context.Context
	ApiService SMFSelectionSubscriptionDataDocumentApi
	ueId string
	servingPlmnId string
	fields *[]string
	supportedFeatures *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// attributes to be retrieved
func (r ApiQuerySmfSelectDataRequest) Fields(fields []string) ApiQuerySmfSelectDataRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQuerySmfSelectDataRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmfSelectDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQuerySmfSelectDataRequest) IfNoneMatch(ifNoneMatch string) ApiQuerySmfSelectDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQuerySmfSelectDataRequest) IfModifiedSince(ifModifiedSince string) ApiQuerySmfSelectDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQuerySmfSelectDataRequest) Execute() (*SmfSelectionSubscriptionData, *http.Response, error) {
	return r.ApiService.QuerySmfSelectDataExecute(r)
}

/*
QuerySmfSelectData Retrieves the SMF selection subscription data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param servingPlmnId PLMN ID
 @return ApiQuerySmfSelectDataRequest
*/
func (a *SMFSelectionSubscriptionDataDocumentApiService) QuerySmfSelectData(ctx context.Context, ueId string, servingPlmnId string) ApiQuerySmfSelectDataRequest {
	return ApiQuerySmfSelectDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		servingPlmnId: servingPlmnId,
	}
}

// Execute executes the request
//  @return SmfSelectionSubscriptionData
func (a *SMFSelectionSubscriptionDataDocumentApiService) QuerySmfSelectDataExecute(r ApiQuerySmfSelectDataRequest) (*SmfSelectionSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmfSelectionSubscriptionData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFSelectionSubscriptionDataDocumentApiService.QuerySmfSelectData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/{servingPlmnId}/provisioned-data/smf-selection-subscription-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingPlmnId"+"}", url.PathEscape(parameterToString(r.servingPlmnId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
