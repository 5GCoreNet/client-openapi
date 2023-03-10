/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type AccessAndMobilitySubscriptionDataRetrievalApi interface {

	/*
	GetAmData retrieve a UE's Access and Mobility Subscription Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param supi Identifier of the UE
	@return ApiGetAmDataRequest
	*/
	GetAmData(ctx context.Context, supi string) ApiGetAmDataRequest

	// GetAmDataExecute executes the request
	//  @return AccessAndMobilitySubscriptionData
	GetAmDataExecute(r ApiGetAmDataRequest) (*AccessAndMobilitySubscriptionData, *http.Response, error)
}

// AccessAndMobilitySubscriptionDataRetrievalApiService AccessAndMobilitySubscriptionDataRetrievalApi service
type AccessAndMobilitySubscriptionDataRetrievalApiService service

type ApiGetAmDataRequest struct {
	ctx context.Context
	ApiService AccessAndMobilitySubscriptionDataRetrievalApi
	supi string
	supportedFeatures *string
	plmnId *PlmnIdNid
	adjacentPlmns *[]PlmnId
	disasterRoamingInd *bool
	ifNoneMatch *string
	ifModifiedSince *string
}

// Supported Features
func (r ApiGetAmDataRequest) SupportedFeatures(supportedFeatures string) ApiGetAmDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Serving PLMN ID or SNPN ID
func (r ApiGetAmDataRequest) PlmnId(plmnId PlmnIdNid) ApiGetAmDataRequest {
	r.plmnId = &plmnId
	return r
}

// List of PLMNs adjacent to the UE&#39;s serving PLMN
func (r ApiGetAmDataRequest) AdjacentPlmns(adjacentPlmns []PlmnId) ApiGetAmDataRequest {
	r.adjacentPlmns = &adjacentPlmns
	return r
}

// Indication whether Disaster Roaming service is applied or not
func (r ApiGetAmDataRequest) DisasterRoamingInd(disasterRoamingInd bool) ApiGetAmDataRequest {
	r.disasterRoamingInd = &disasterRoamingInd
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiGetAmDataRequest) IfNoneMatch(ifNoneMatch string) ApiGetAmDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiGetAmDataRequest) IfModifiedSince(ifModifiedSince string) ApiGetAmDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiGetAmDataRequest) Execute() (*AccessAndMobilitySubscriptionData, *http.Response, error) {
	return r.ApiService.GetAmDataExecute(r)
}

/*
GetAmData retrieve a UE's Access and Mobility Subscription Data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param supi Identifier of the UE
 @return ApiGetAmDataRequest
*/
func (a *AccessAndMobilitySubscriptionDataRetrievalApiService) GetAmData(ctx context.Context, supi string) ApiGetAmDataRequest {
	return ApiGetAmDataRequest{
		ApiService: a,
		ctx: ctx,
		supi: supi,
	}
}

// Execute executes the request
//  @return AccessAndMobilitySubscriptionData
func (a *AccessAndMobilitySubscriptionDataRetrievalApiService) GetAmDataExecute(r ApiGetAmDataRequest) (*AccessAndMobilitySubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessAndMobilitySubscriptionData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessAndMobilitySubscriptionDataRetrievalApiService.GetAmData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{supi}/am-data"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", url.PathEscape(parameterToString(r.supi, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
	if r.plmnId != nil {
		localVarQueryParams.Add("plmn-id", parameterToString(*r.plmnId, ""))
	}
	if r.adjacentPlmns != nil {
		localVarQueryParams.Add("adjacent-plmns", parameterToString(*r.adjacentPlmns, "csv"))
	}
	if r.disasterRoamingInd != nil {
		localVarQueryParams.Add("disaster-roaming-ind", parameterToString(*r.disasterRoamingInd, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		if localVarHTTPResponse.StatusCode == 500 {
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
		if localVarHTTPResponse.StatusCode == 503 {
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
