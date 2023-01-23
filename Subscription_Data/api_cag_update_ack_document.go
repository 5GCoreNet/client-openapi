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


type CAGUpdateAckDocumentApi interface {

	/*
	CreateCagUpdateAck To store the CAG update acknowledgement information of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiCreateCagUpdateAckRequest
	*/
	CreateCagUpdateAck(ctx context.Context, ueId string) ApiCreateCagUpdateAckRequest

	// CreateCagUpdateAckExecute executes the request
	CreateCagUpdateAckExecute(r ApiCreateCagUpdateAckRequest) (*http.Response, error)
}

// CAGUpdateAckDocumentApiService CAGUpdateAckDocumentApi service
type CAGUpdateAckDocumentApiService service

type ApiCreateCagUpdateAckRequest struct {
	ctx context.Context
	ApiService CAGUpdateAckDocumentApi
	ueId string
	cagAckData *CagAckData
	supportedFeatures *string
}

func (r ApiCreateCagUpdateAckRequest) CagAckData(cagAckData CagAckData) ApiCreateCagUpdateAckRequest {
	r.cagAckData = &cagAckData
	return r
}

// Supported Features
func (r ApiCreateCagUpdateAckRequest) SupportedFeatures(supportedFeatures string) ApiCreateCagUpdateAckRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiCreateCagUpdateAckRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateCagUpdateAckExecute(r)
}

/*
CreateCagUpdateAck To store the CAG update acknowledgement information of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiCreateCagUpdateAckRequest
*/
func (a *CAGUpdateAckDocumentApiService) CreateCagUpdateAck(ctx context.Context, ueId string) ApiCreateCagUpdateAckRequest {
	return ApiCreateCagUpdateAckRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
func (a *CAGUpdateAckDocumentApiService) CreateCagUpdateAckExecute(r ApiCreateCagUpdateAckRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CAGUpdateAckDocumentApiService.CreateCagUpdateAck")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/ue-update-confirmation-data/subscribed-cag"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterToString(r.ueId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cagAckData == nil {
		return nil, reportError("cagAckData is required and must be specified")
	}

	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
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
	localVarPostBody = r.cagAckData
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
