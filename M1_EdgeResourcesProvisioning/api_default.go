/*
M1_EdgeResourcesProvisioning

5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_EdgeResourcesProvisioning

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateEdgeResourcesConfigurationRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	provisioningSessionId string
	edgeResourcesConfiguration *EdgeResourcesConfiguration
}

// A JSON representation of an Edge Resources Configuration
func (r ApiCreateEdgeResourcesConfigurationRequest) EdgeResourcesConfiguration(edgeResourcesConfiguration EdgeResourcesConfiguration) ApiCreateEdgeResourcesConfigurationRequest {
	r.edgeResourcesConfiguration = &edgeResourcesConfiguration
	return r
}

func (r ApiCreateEdgeResourcesConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateEdgeResourcesConfigurationExecute(r)
}

/*
CreateEdgeResourcesConfiguration Create an Edge Resources Configuration within the scope of the specified Provisioning Session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioningSessionId The resource identifier of an existing Provisioning Session.
 @return ApiCreateEdgeResourcesConfigurationRequest
*/
func (a *DefaultApiService) CreateEdgeResourcesConfiguration(ctx context.Context, provisioningSessionId string) ApiCreateEdgeResourcesConfigurationRequest {
	return ApiCreateEdgeResourcesConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		provisioningSessionId: provisioningSessionId,
	}
}

// Execute executes the request
func (a *DefaultApiService) CreateEdgeResourcesConfigurationExecute(r ApiCreateEdgeResourcesConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateEdgeResourcesConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/edge-resources-configurations"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterToString(r.provisioningSessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.edgeResourcesConfiguration == nil {
		return nil, reportError("edgeResourcesConfiguration is required and must be specified")
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
	localVarPostBody = r.edgeResourcesConfiguration
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

type ApiDestroyEdgeResourcesConfigurationRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	provisioningSessionId string
	edgeResourcesConfigurationId string
}

func (r ApiDestroyEdgeResourcesConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DestroyEdgeResourcesConfigurationExecute(r)
}

/*
DestroyEdgeResourcesConfiguration Method for DestroyEdgeResourcesConfiguration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioningSessionId The resource identifier of an existing Provisioning Session.
 @param edgeResourcesConfigurationId The resource identifier of an existing Edge Resources Configuration.
 @return ApiDestroyEdgeResourcesConfigurationRequest
*/
func (a *DefaultApiService) DestroyEdgeResourcesConfiguration(ctx context.Context, provisioningSessionId string, edgeResourcesConfigurationId string) ApiDestroyEdgeResourcesConfigurationRequest {
	return ApiDestroyEdgeResourcesConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		provisioningSessionId: provisioningSessionId,
		edgeResourcesConfigurationId: edgeResourcesConfigurationId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DestroyEdgeResourcesConfigurationExecute(r ApiDestroyEdgeResourcesConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DestroyEdgeResourcesConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterToString(r.provisioningSessionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"edgeResourcesConfigurationId"+"}", url.PathEscape(parameterToString(r.edgeResourcesConfigurationId, "")), -1)

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

type ApiPatchEdgeResourcesConfigurationRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	provisioningSessionId string
	edgeResourcesConfigurationId string
	edgeResourcesConfiguration *EdgeResourcesConfiguration
}

// A JSON representation of a Edge Resources Configuration
func (r ApiPatchEdgeResourcesConfigurationRequest) EdgeResourcesConfiguration(edgeResourcesConfiguration EdgeResourcesConfiguration) ApiPatchEdgeResourcesConfigurationRequest {
	r.edgeResourcesConfiguration = &edgeResourcesConfiguration
	return r
}

func (r ApiPatchEdgeResourcesConfigurationRequest) Execute() (*EdgeResourcesConfiguration, *http.Response, error) {
	return r.ApiService.PatchEdgeResourcesConfigurationExecute(r)
}

/*
PatchEdgeResourcesConfiguration Patch the Edge Resources Configuration for the specified Provisioning Session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioningSessionId The resource identifier of an existing Provisioning Session.
 @param edgeResourcesConfigurationId The resource identifier of an existing Edge Resources Configuration.
 @return ApiPatchEdgeResourcesConfigurationRequest
*/
func (a *DefaultApiService) PatchEdgeResourcesConfiguration(ctx context.Context, provisioningSessionId string, edgeResourcesConfigurationId string) ApiPatchEdgeResourcesConfigurationRequest {
	return ApiPatchEdgeResourcesConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		provisioningSessionId: provisioningSessionId,
		edgeResourcesConfigurationId: edgeResourcesConfigurationId,
	}
}

// Execute executes the request
//  @return EdgeResourcesConfiguration
func (a *DefaultApiService) PatchEdgeResourcesConfigurationExecute(r ApiPatchEdgeResourcesConfigurationRequest) (*EdgeResourcesConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EdgeResourcesConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.PatchEdgeResourcesConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterToString(r.provisioningSessionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"edgeResourcesConfigurationId"+"}", url.PathEscape(parameterToString(r.edgeResourcesConfigurationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.edgeResourcesConfiguration == nil {
		return localVarReturnValue, nil, reportError("edgeResourcesConfiguration is required and must be specified")
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
	localVarPostBody = r.edgeResourcesConfiguration
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

type ApiRetrieveEdgeResourcesConfigurationRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	provisioningSessionId string
	edgeResourcesConfigurationId string
}

func (r ApiRetrieveEdgeResourcesConfigurationRequest) Execute() (*EdgeResourcesConfiguration, *http.Response, error) {
	return r.ApiService.RetrieveEdgeResourcesConfigurationExecute(r)
}

/*
RetrieveEdgeResourcesConfiguration Retrieve the Edge Resources Configuration of the specified Provisioning Session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioningSessionId The resource identifier of an existing Provisioning Session.
 @param edgeResourcesConfigurationId The resource identifier of an existing Edge Resources Configuration.
 @return ApiRetrieveEdgeResourcesConfigurationRequest
*/
func (a *DefaultApiService) RetrieveEdgeResourcesConfiguration(ctx context.Context, provisioningSessionId string, edgeResourcesConfigurationId string) ApiRetrieveEdgeResourcesConfigurationRequest {
	return ApiRetrieveEdgeResourcesConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		provisioningSessionId: provisioningSessionId,
		edgeResourcesConfigurationId: edgeResourcesConfigurationId,
	}
}

// Execute executes the request
//  @return EdgeResourcesConfiguration
func (a *DefaultApiService) RetrieveEdgeResourcesConfigurationExecute(r ApiRetrieveEdgeResourcesConfigurationRequest) (*EdgeResourcesConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EdgeResourcesConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RetrieveEdgeResourcesConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterToString(r.provisioningSessionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"edgeResourcesConfigurationId"+"}", url.PathEscape(parameterToString(r.edgeResourcesConfigurationId, "")), -1)

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

type ApiUpdateEdgeResourcesConfigurationRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	provisioningSessionId string
	edgeResourcesConfigurationId string
	edgeResourcesConfiguration *EdgeResourcesConfiguration
}

// A JSON representation of an Edge Resources Configuration
func (r ApiUpdateEdgeResourcesConfigurationRequest) EdgeResourcesConfiguration(edgeResourcesConfiguration EdgeResourcesConfiguration) ApiUpdateEdgeResourcesConfigurationRequest {
	r.edgeResourcesConfiguration = &edgeResourcesConfiguration
	return r
}

func (r ApiUpdateEdgeResourcesConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateEdgeResourcesConfigurationExecute(r)
}

/*
UpdateEdgeResourcesConfiguration Update an Edge Resources Configuration for the specified Provisioning Session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioningSessionId The resource identifier of an existing Provisioning Session.
 @param edgeResourcesConfigurationId The resource identifier of an existing Edge Resources Configuration.
 @return ApiUpdateEdgeResourcesConfigurationRequest
*/
func (a *DefaultApiService) UpdateEdgeResourcesConfiguration(ctx context.Context, provisioningSessionId string, edgeResourcesConfigurationId string) ApiUpdateEdgeResourcesConfigurationRequest {
	return ApiUpdateEdgeResourcesConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		provisioningSessionId: provisioningSessionId,
		edgeResourcesConfigurationId: edgeResourcesConfigurationId,
	}
}

// Execute executes the request
func (a *DefaultApiService) UpdateEdgeResourcesConfigurationExecute(r ApiUpdateEdgeResourcesConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateEdgeResourcesConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterToString(r.provisioningSessionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"edgeResourcesConfigurationId"+"}", url.PathEscape(parameterToString(r.edgeResourcesConfigurationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.edgeResourcesConfiguration == nil {
		return nil, reportError("edgeResourcesConfiguration is required and must be specified")
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
	localVarPostBody = r.edgeResourcesConfiguration
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
