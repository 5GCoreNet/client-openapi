/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type NIDDAuthorizationInfoDocumentApi interface {

	/*
	CreateNIDDAuthorizationInfo Create NIDD Authorization Info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@return ApiCreateNIDDAuthorizationInfoRequest
	*/
	CreateNIDDAuthorizationInfo(ctx context.Context, ueId string) ApiCreateNIDDAuthorizationInfoRequest

	// CreateNIDDAuthorizationInfoExecute executes the request
	//  @return NiddAuthorizationInfo
	CreateNIDDAuthorizationInfoExecute(r ApiCreateNIDDAuthorizationInfoRequest) (*NiddAuthorizationInfo, *http.Response, error)

	/*
	GetNiddAuthorizationInfo Retrieve NIDD Authorization Info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@return ApiGetNiddAuthorizationInfoRequest
	*/
	GetNiddAuthorizationInfo(ctx context.Context, ueId string) ApiGetNiddAuthorizationInfoRequest

	// GetNiddAuthorizationInfoExecute executes the request
	//  @return NiddAuthorizationInfo
	GetNiddAuthorizationInfoExecute(r ApiGetNiddAuthorizationInfoRequest) (*NiddAuthorizationInfo, *http.Response, error)

	/*
	ModifyNiddAuthorizationInfo Modify NIDD Authorization Info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@return ApiModifyNiddAuthorizationInfoRequest
	*/
	ModifyNiddAuthorizationInfo(ctx context.Context, ueId string) ApiModifyNiddAuthorizationInfoRequest

	// ModifyNiddAuthorizationInfoExecute executes the request
	//  @return PatchResult
	ModifyNiddAuthorizationInfoExecute(r ApiModifyNiddAuthorizationInfoRequest) (*PatchResult, *http.Response, error)

	/*
	RemoveNiddAuthorizationInfo Delete NIDD Authorization Info

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@return ApiRemoveNiddAuthorizationInfoRequest
	*/
	RemoveNiddAuthorizationInfo(ctx context.Context, ueId string) ApiRemoveNiddAuthorizationInfoRequest

	// RemoveNiddAuthorizationInfoExecute executes the request
	RemoveNiddAuthorizationInfoExecute(r ApiRemoveNiddAuthorizationInfoRequest) (*http.Response, error)
}

// NIDDAuthorizationInfoDocumentApiService NIDDAuthorizationInfoDocumentApi service
type NIDDAuthorizationInfoDocumentApiService service

type ApiCreateNIDDAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService NIDDAuthorizationInfoDocumentApi
	ueId string
	niddAuthorizationInfo *NiddAuthorizationInfo
}

func (r ApiCreateNIDDAuthorizationInfoRequest) NiddAuthorizationInfo(niddAuthorizationInfo NiddAuthorizationInfo) ApiCreateNIDDAuthorizationInfoRequest {
	r.niddAuthorizationInfo = &niddAuthorizationInfo
	return r
}

func (r ApiCreateNIDDAuthorizationInfoRequest) Execute() (*NiddAuthorizationInfo, *http.Response, error) {
	return r.ApiService.CreateNIDDAuthorizationInfoExecute(r)
}

/*
CreateNIDDAuthorizationInfo Create NIDD Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @return ApiCreateNIDDAuthorizationInfoRequest
*/
func (a *NIDDAuthorizationInfoDocumentApiService) CreateNIDDAuthorizationInfo(ctx context.Context, ueId string) ApiCreateNIDDAuthorizationInfoRequest {
	return ApiCreateNIDDAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return NiddAuthorizationInfo
func (a *NIDDAuthorizationInfoDocumentApiService) CreateNIDDAuthorizationInfoExecute(r ApiCreateNIDDAuthorizationInfoRequest) (*NiddAuthorizationInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NiddAuthorizationInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NIDDAuthorizationInfoDocumentApiService.CreateNIDDAuthorizationInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.niddAuthorizationInfo == nil {
		return localVarReturnValue, nil, reportError("niddAuthorizationInfo is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.niddAuthorizationInfo
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

type ApiGetNiddAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService NIDDAuthorizationInfoDocumentApi
	ueId string
}

func (r ApiGetNiddAuthorizationInfoRequest) Execute() (*NiddAuthorizationInfo, *http.Response, error) {
	return r.ApiService.GetNiddAuthorizationInfoExecute(r)
}

/*
GetNiddAuthorizationInfo Retrieve NIDD Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @return ApiGetNiddAuthorizationInfoRequest
*/
func (a *NIDDAuthorizationInfoDocumentApiService) GetNiddAuthorizationInfo(ctx context.Context, ueId string) ApiGetNiddAuthorizationInfoRequest {
	return ApiGetNiddAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return NiddAuthorizationInfo
func (a *NIDDAuthorizationInfoDocumentApiService) GetNiddAuthorizationInfoExecute(r ApiGetNiddAuthorizationInfoRequest) (*NiddAuthorizationInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NiddAuthorizationInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NIDDAuthorizationInfoDocumentApiService.GetNiddAuthorizationInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)

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

type ApiModifyNiddAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService NIDDAuthorizationInfoDocumentApi
	ueId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyNiddAuthorizationInfoRequest) PatchItem(patchItem []PatchItem) ApiModifyNiddAuthorizationInfoRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyNiddAuthorizationInfoRequest) SupportedFeatures(supportedFeatures string) ApiModifyNiddAuthorizationInfoRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyNiddAuthorizationInfoRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyNiddAuthorizationInfoExecute(r)
}

/*
ModifyNiddAuthorizationInfo Modify NIDD Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @return ApiModifyNiddAuthorizationInfoRequest
*/
func (a *NIDDAuthorizationInfoDocumentApiService) ModifyNiddAuthorizationInfo(ctx context.Context, ueId string) ApiModifyNiddAuthorizationInfoRequest {
	return ApiModifyNiddAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *NIDDAuthorizationInfoDocumentApiService) ModifyNiddAuthorizationInfoExecute(r ApiModifyNiddAuthorizationInfoRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NIDDAuthorizationInfoDocumentApiService.ModifyNiddAuthorizationInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}

	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

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
	localVarPostBody = r.patchItem
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

type ApiRemoveNiddAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService NIDDAuthorizationInfoDocumentApi
	ueId string
}

func (r ApiRemoveNiddAuthorizationInfoRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveNiddAuthorizationInfoExecute(r)
}

/*
RemoveNiddAuthorizationInfo Delete NIDD Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @return ApiRemoveNiddAuthorizationInfoRequest
*/
func (a *NIDDAuthorizationInfoDocumentApiService) RemoveNiddAuthorizationInfo(ctx context.Context, ueId string) ApiRemoveNiddAuthorizationInfoRequest {
	return ApiRemoveNiddAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
func (a *NIDDAuthorizationInfoDocumentApiService) RemoveNiddAuthorizationInfoExecute(r ApiRemoveNiddAuthorizationInfoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NIDDAuthorizationInfoDocumentApiService.RemoveNiddAuthorizationInfo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/nidd-authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
