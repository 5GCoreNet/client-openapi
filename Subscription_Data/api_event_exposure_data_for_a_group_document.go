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


type EventExposureDataForAGroupDocumentApi interface {

	/*
	QueryGroupEEData Retrieves the ee profile data profile data of a group or anyUE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueGroupId Group of UEs or any UE
	@return ApiQueryGroupEEDataRequest
	*/
	QueryGroupEEData(ctx context.Context, ueGroupId string) ApiQueryGroupEEDataRequest

	// QueryGroupEEDataExecute executes the request
	//  @return EeGroupProfileData
	QueryGroupEEDataExecute(r ApiQueryGroupEEDataRequest) (*EeGroupProfileData, *http.Response, error)
}

// EventExposureDataForAGroupDocumentApiService EventExposureDataForAGroupDocumentApi service
type EventExposureDataForAGroupDocumentApiService service

type ApiQueryGroupEEDataRequest struct {
	ctx context.Context
	ApiService EventExposureDataForAGroupDocumentApi
	ueGroupId string
	supportedFeatures *string
}

// Supported Features
func (r ApiQueryGroupEEDataRequest) SupportedFeatures(supportedFeatures string) ApiQueryGroupEEDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQueryGroupEEDataRequest) Execute() (*EeGroupProfileData, *http.Response, error) {
	return r.ApiService.QueryGroupEEDataExecute(r)
}

/*
QueryGroupEEData Retrieves the ee profile data profile data of a group or anyUE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId Group of UEs or any UE
 @return ApiQueryGroupEEDataRequest
*/
func (a *EventExposureDataForAGroupDocumentApiService) QueryGroupEEData(ctx context.Context, ueGroupId string) ApiQueryGroupEEDataRequest {
	return ApiQueryGroupEEDataRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
	}
}

// Execute executes the request
//  @return EeGroupProfileData
func (a *EventExposureDataForAGroupDocumentApiService) QueryGroupEEDataExecute(r ApiQueryGroupEEDataRequest) (*EeGroupProfileData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EeGroupProfileData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventExposureDataForAGroupDocumentApiService.QueryGroupEEData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-profile-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterToString(r.ueGroupId, "")), -1)

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
