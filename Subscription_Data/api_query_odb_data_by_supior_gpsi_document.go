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


type QueryODBDataBySUPIOrGPSIDocumentApi interface {

	/*
	GetOdbData Retrieve ODB Data data by SUPI or GPSI

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE ID
	@return ApiGetOdbDataRequest
	*/
	GetOdbData(ctx context.Context, ueId string) ApiGetOdbDataRequest

	// GetOdbDataExecute executes the request
	//  @return OdbData
	GetOdbDataExecute(r ApiGetOdbDataRequest) (*OdbData, *http.Response, error)
}

// QueryODBDataBySUPIOrGPSIDocumentApiService QueryODBDataBySUPIOrGPSIDocumentApi service
type QueryODBDataBySUPIOrGPSIDocumentApiService service

type ApiGetOdbDataRequest struct {
	ctx context.Context
	ApiService QueryODBDataBySUPIOrGPSIDocumentApi
	ueId string
}

func (r ApiGetOdbDataRequest) Execute() (*OdbData, *http.Response, error) {
	return r.ApiService.GetOdbDataExecute(r)
}

/*
GetOdbData Retrieve ODB Data data by SUPI or GPSI

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE ID
 @return ApiGetOdbDataRequest
*/
func (a *QueryODBDataBySUPIOrGPSIDocumentApiService) GetOdbData(ctx context.Context, ueId string) ApiGetOdbDataRequest {
	return ApiGetOdbDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return OdbData
func (a *QueryODBDataBySUPIOrGPSIDocumentApiService) GetOdbDataExecute(r ApiGetOdbDataRequest) (*OdbData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OdbData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryODBDataBySUPIOrGPSIDocumentApiService.GetOdbData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/operator-determined-barring-data"
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
