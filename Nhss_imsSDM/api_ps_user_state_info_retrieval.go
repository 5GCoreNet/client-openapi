/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type PSUserStateInfoRetrievalApi interface {

	/*
	GetPsUserStateInfo Retrieve the user state information in PS domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param imsUeId IMS Identity
	@return ApiGetPsUserStateInfoRequest
	*/
	GetPsUserStateInfo(ctx context.Context, imsUeId string) ApiGetPsUserStateInfoRequest

	// GetPsUserStateInfoExecute executes the request
	//  @return PsUserState
	GetPsUserStateInfoExecute(r ApiGetPsUserStateInfoRequest) (*PsUserState, *http.Response, error)
}

// PSUserStateInfoRetrievalApiService PSUserStateInfoRetrievalApi service
type PSUserStateInfoRetrievalApiService service

type ApiGetPsUserStateInfoRequest struct {
	ctx context.Context
	ApiService PSUserStateInfoRetrievalApi
	imsUeId string
	requestedNodes *[]RequestedNode
	privateIdentity *string
}

// Indicates the serving node(s) for which the request is applicable.
func (r ApiGetPsUserStateInfoRequest) RequestedNodes(requestedNodes []RequestedNode) ApiGetPsUserStateInfoRequest {
	r.requestedNodes = &requestedNodes
	return r
}

// IMS Private Identity
func (r ApiGetPsUserStateInfoRequest) PrivateIdentity(privateIdentity string) ApiGetPsUserStateInfoRequest {
	r.privateIdentity = &privateIdentity
	return r
}

func (r ApiGetPsUserStateInfoRequest) Execute() (*PsUserState, *http.Response, error) {
	return r.ApiService.GetPsUserStateInfoExecute(r)
}

/*
GetPsUserStateInfo Retrieve the user state information in PS domain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imsUeId IMS Identity
 @return ApiGetPsUserStateInfoRequest
*/
func (a *PSUserStateInfoRetrievalApiService) GetPsUserStateInfo(ctx context.Context, imsUeId string) ApiGetPsUserStateInfoRequest {
	return ApiGetPsUserStateInfoRequest{
		ApiService: a,
		ctx: ctx,
		imsUeId: imsUeId,
	}
}

// Execute executes the request
//  @return PsUserState
func (a *PSUserStateInfoRetrievalApiService) GetPsUserStateInfoExecute(r ApiGetPsUserStateInfoRequest) (*PsUserState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PsUserState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PSUserStateInfoRetrievalApiService.GetPsUserStateInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{imsUeId}/access-data/ps-domain/user-state"
	localVarPath = strings.Replace(localVarPath, "{"+"imsUeId"+"}", url.PathEscape(parameterToString(r.imsUeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.requestedNodes != nil {
		localVarQueryParams.Add("requested-nodes", parameterToString(*r.requestedNodes, "csv"))
	}
	if r.privateIdentity != nil {
		localVarQueryParams.Add("private-identity", parameterToString(*r.privateIdentity, ""))
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
		if localVarHTTPResponse.StatusCode == 307 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 308 {
			var v RedirectResponse
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
		if localVarHTTPResponse.StatusCode == 504 {
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
