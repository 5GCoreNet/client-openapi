/*
M5_NetworkAssistance

5GMS AF M5 Network Assistance API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_NetworkAssistance

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type DefaultApi interface {

	/*
	CreateNetworkAssistanceSession Create a new Network Assistance Session.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateNetworkAssistanceSessionRequest
	*/
	CreateNetworkAssistanceSession(ctx context.Context) ApiCreateNetworkAssistanceSessionRequest

	// CreateNetworkAssistanceSessionExecute executes the request
	//  @return NetworkAssistanceSession
	CreateNetworkAssistanceSessionExecute(r ApiCreateNetworkAssistanceSessionRequest) (*NetworkAssistanceSession, *http.Response, error)

	/*
	DestroyNetworkAssistanceSession Destroy an existing Network Assistance Session resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param naSessionId The resource identifier of an existing Network Assistance Session resource
	@return ApiDestroyNetworkAssistanceSessionRequest
	*/
	DestroyNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiDestroyNetworkAssistanceSessionRequest

	// DestroyNetworkAssistanceSessionExecute executes the request
	DestroyNetworkAssistanceSessionExecute(r ApiDestroyNetworkAssistanceSessionRequest) (*http.Response, error)

	/*
	PatchNetworkAssistanceSession Patch an existing Network Assistance Session resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param naSessionId The resource identifier of an existing Network Assistance Session resource
	@return ApiPatchNetworkAssistanceSessionRequest
	*/
	PatchNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiPatchNetworkAssistanceSessionRequest

	// PatchNetworkAssistanceSessionExecute executes the request
	//  @return NetworkAssistanceSession
	PatchNetworkAssistanceSessionExecute(r ApiPatchNetworkAssistanceSessionRequest) (*NetworkAssistanceSession, *http.Response, error)

	/*
	RequestBitRateRecommendation Obtain a bit rate recommendation for the next recommendation window

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param naSessionId The resource identifier of an existing Network Assistance Session resource
	@return ApiRequestBitRateRecommendationRequest
	*/
	RequestBitRateRecommendation(ctx context.Context, naSessionId string) ApiRequestBitRateRecommendationRequest

	// RequestBitRateRecommendationExecute executes the request
	//  @return M5QoSSpecification
	RequestBitRateRecommendationExecute(r ApiRequestBitRateRecommendationRequest) (*M5QoSSpecification, *http.Response, error)

	/*
	RequestDeliveryBoost Request a delivery boost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param naSessionId The resource identifier of an existing Network Assistance Session resource
	@return ApiRequestDeliveryBoostRequest
	*/
	RequestDeliveryBoost(ctx context.Context, naSessionId string) ApiRequestDeliveryBoostRequest

	// RequestDeliveryBoostExecute executes the request
	//  @return OperationSuccessResponse
	RequestDeliveryBoostExecute(r ApiRequestDeliveryBoostRequest) (*OperationSuccessResponse, *http.Response, error)

	/*
	RetrieveNetworkAssistanceSession Retrieve an existing Network Assistance Session resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param naSessionId The resource identifier of an existing Network Assistance Session resource
	@return ApiRetrieveNetworkAssistanceSessionRequest
	*/
	RetrieveNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiRetrieveNetworkAssistanceSessionRequest

	// RetrieveNetworkAssistanceSessionExecute executes the request
	//  @return NetworkAssistanceSession
	RetrieveNetworkAssistanceSessionExecute(r ApiRetrieveNetworkAssistanceSessionRequest) (*NetworkAssistanceSession, *http.Response, error)

	/*
	UpdateNetworkAssistanceSession Update an existing Network Assistance Session resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param naSessionId The resource identifier of an existing Network Assistance Session resource
	@return ApiUpdateNetworkAssistanceSessionRequest
	*/
	UpdateNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiUpdateNetworkAssistanceSessionRequest

	// UpdateNetworkAssistanceSessionExecute executes the request
	UpdateNetworkAssistanceSessionExecute(r ApiUpdateNetworkAssistanceSessionRequest) (*http.Response, error)
}

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateNetworkAssistanceSessionRequest struct {
	ctx context.Context
	ApiService DefaultApi
}

func (r ApiCreateNetworkAssistanceSessionRequest) Execute() (*NetworkAssistanceSession, *http.Response, error) {
	return r.ApiService.CreateNetworkAssistanceSessionExecute(r)
}

/*
CreateNetworkAssistanceSession Create a new Network Assistance Session.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNetworkAssistanceSessionRequest
*/
func (a *DefaultApiService) CreateNetworkAssistanceSession(ctx context.Context) ApiCreateNetworkAssistanceSessionRequest {
	return ApiCreateNetworkAssistanceSessionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkAssistanceSession
func (a *DefaultApiService) CreateNetworkAssistanceSessionExecute(r ApiCreateNetworkAssistanceSessionRequest) (*NetworkAssistanceSession, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NetworkAssistanceSession
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateNetworkAssistanceSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/"

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

type ApiDestroyNetworkAssistanceSessionRequest struct {
	ctx context.Context
	ApiService DefaultApi
	naSessionId string
}

func (r ApiDestroyNetworkAssistanceSessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DestroyNetworkAssistanceSessionExecute(r)
}

/*
DestroyNetworkAssistanceSession Destroy an existing Network Assistance Session resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param naSessionId The resource identifier of an existing Network Assistance Session resource
 @return ApiDestroyNetworkAssistanceSessionRequest
*/
func (a *DefaultApiService) DestroyNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiDestroyNetworkAssistanceSessionRequest {
	return ApiDestroyNetworkAssistanceSessionRequest{
		ApiService: a,
		ctx: ctx,
		naSessionId: naSessionId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DestroyNetworkAssistanceSessionExecute(r ApiDestroyNetworkAssistanceSessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DestroyNetworkAssistanceSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/{naSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"naSessionId"+"}", url.PathEscape(parameterToString(r.naSessionId, "")), -1)

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

type ApiPatchNetworkAssistanceSessionRequest struct {
	ctx context.Context
	ApiService DefaultApi
	naSessionId string
	networkAssistanceSession *NetworkAssistanceSession
}

// A JSON patch to a Network Assistance Session resource
func (r ApiPatchNetworkAssistanceSessionRequest) NetworkAssistanceSession(networkAssistanceSession NetworkAssistanceSession) ApiPatchNetworkAssistanceSessionRequest {
	r.networkAssistanceSession = &networkAssistanceSession
	return r
}

func (r ApiPatchNetworkAssistanceSessionRequest) Execute() (*NetworkAssistanceSession, *http.Response, error) {
	return r.ApiService.PatchNetworkAssistanceSessionExecute(r)
}

/*
PatchNetworkAssistanceSession Patch an existing Network Assistance Session resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param naSessionId The resource identifier of an existing Network Assistance Session resource
 @return ApiPatchNetworkAssistanceSessionRequest
*/
func (a *DefaultApiService) PatchNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiPatchNetworkAssistanceSessionRequest {
	return ApiPatchNetworkAssistanceSessionRequest{
		ApiService: a,
		ctx: ctx,
		naSessionId: naSessionId,
	}
}

// Execute executes the request
//  @return NetworkAssistanceSession
func (a *DefaultApiService) PatchNetworkAssistanceSessionExecute(r ApiPatchNetworkAssistanceSessionRequest) (*NetworkAssistanceSession, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NetworkAssistanceSession
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.PatchNetworkAssistanceSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/{naSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"naSessionId"+"}", url.PathEscape(parameterToString(r.naSessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.networkAssistanceSession == nil {
		return localVarReturnValue, nil, reportError("networkAssistanceSession is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json", "application/json-patch+json"}

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
	localVarPostBody = r.networkAssistanceSession
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

type ApiRequestBitRateRecommendationRequest struct {
	ctx context.Context
	ApiService DefaultApi
	naSessionId string
}

func (r ApiRequestBitRateRecommendationRequest) Execute() (*M5QoSSpecification, *http.Response, error) {
	return r.ApiService.RequestBitRateRecommendationExecute(r)
}

/*
RequestBitRateRecommendation Obtain a bit rate recommendation for the next recommendation window

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param naSessionId The resource identifier of an existing Network Assistance Session resource
 @return ApiRequestBitRateRecommendationRequest
*/
func (a *DefaultApiService) RequestBitRateRecommendation(ctx context.Context, naSessionId string) ApiRequestBitRateRecommendationRequest {
	return ApiRequestBitRateRecommendationRequest{
		ApiService: a,
		ctx: ctx,
		naSessionId: naSessionId,
	}
}

// Execute executes the request
//  @return M5QoSSpecification
func (a *DefaultApiService) RequestBitRateRecommendationExecute(r ApiRequestBitRateRecommendationRequest) (*M5QoSSpecification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *M5QoSSpecification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RequestBitRateRecommendation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/{naSessionId}/recommendation"
	localVarPath = strings.Replace(localVarPath, "{"+"naSessionId"+"}", url.PathEscape(parameterToString(r.naSessionId, "")), -1)

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

type ApiRequestDeliveryBoostRequest struct {
	ctx context.Context
	ApiService DefaultApi
	naSessionId string
}

func (r ApiRequestDeliveryBoostRequest) Execute() (*OperationSuccessResponse, *http.Response, error) {
	return r.ApiService.RequestDeliveryBoostExecute(r)
}

/*
RequestDeliveryBoost Request a delivery boost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param naSessionId The resource identifier of an existing Network Assistance Session resource
 @return ApiRequestDeliveryBoostRequest
*/
func (a *DefaultApiService) RequestDeliveryBoost(ctx context.Context, naSessionId string) ApiRequestDeliveryBoostRequest {
	return ApiRequestDeliveryBoostRequest{
		ApiService: a,
		ctx: ctx,
		naSessionId: naSessionId,
	}
}

// Execute executes the request
//  @return OperationSuccessResponse
func (a *DefaultApiService) RequestDeliveryBoostExecute(r ApiRequestDeliveryBoostRequest) (*OperationSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OperationSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RequestDeliveryBoost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/{naSessionId}/boost-request"
	localVarPath = strings.Replace(localVarPath, "{"+"naSessionId"+"}", url.PathEscape(parameterToString(r.naSessionId, "")), -1)

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

type ApiRetrieveNetworkAssistanceSessionRequest struct {
	ctx context.Context
	ApiService DefaultApi
	naSessionId string
}

func (r ApiRetrieveNetworkAssistanceSessionRequest) Execute() (*NetworkAssistanceSession, *http.Response, error) {
	return r.ApiService.RetrieveNetworkAssistanceSessionExecute(r)
}

/*
RetrieveNetworkAssistanceSession Retrieve an existing Network Assistance Session resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param naSessionId The resource identifier of an existing Network Assistance Session resource
 @return ApiRetrieveNetworkAssistanceSessionRequest
*/
func (a *DefaultApiService) RetrieveNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiRetrieveNetworkAssistanceSessionRequest {
	return ApiRetrieveNetworkAssistanceSessionRequest{
		ApiService: a,
		ctx: ctx,
		naSessionId: naSessionId,
	}
}

// Execute executes the request
//  @return NetworkAssistanceSession
func (a *DefaultApiService) RetrieveNetworkAssistanceSessionExecute(r ApiRetrieveNetworkAssistanceSessionRequest) (*NetworkAssistanceSession, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NetworkAssistanceSession
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RetrieveNetworkAssistanceSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/{naSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"naSessionId"+"}", url.PathEscape(parameterToString(r.naSessionId, "")), -1)

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

type ApiUpdateNetworkAssistanceSessionRequest struct {
	ctx context.Context
	ApiService DefaultApi
	naSessionId string
	networkAssistanceSession *NetworkAssistanceSession
}

// A replacement JSON representation of a Network Assistance Session resource
func (r ApiUpdateNetworkAssistanceSessionRequest) NetworkAssistanceSession(networkAssistanceSession NetworkAssistanceSession) ApiUpdateNetworkAssistanceSessionRequest {
	r.networkAssistanceSession = &networkAssistanceSession
	return r
}

func (r ApiUpdateNetworkAssistanceSessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateNetworkAssistanceSessionExecute(r)
}

/*
UpdateNetworkAssistanceSession Update an existing Network Assistance Session resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param naSessionId The resource identifier of an existing Network Assistance Session resource
 @return ApiUpdateNetworkAssistanceSessionRequest
*/
func (a *DefaultApiService) UpdateNetworkAssistanceSession(ctx context.Context, naSessionId string) ApiUpdateNetworkAssistanceSessionRequest {
	return ApiUpdateNetworkAssistanceSessionRequest{
		ApiService: a,
		ctx: ctx,
		naSessionId: naSessionId,
	}
}

// Execute executes the request
func (a *DefaultApiService) UpdateNetworkAssistanceSessionExecute(r ApiUpdateNetworkAssistanceSessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateNetworkAssistanceSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-assistance/{naSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"naSessionId"+"}", url.PathEscape(parameterToString(r.naSessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.networkAssistanceSession == nil {
		return nil, reportError("networkAssistanceSession is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.networkAssistanceSession
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
