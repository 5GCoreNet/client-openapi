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


type AuthenticationSubscriptionDocumentApi interface {

	/*
	ModifyAuthenticationSubscription modify the authentication subscription data of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiModifyAuthenticationSubscriptionRequest
	*/
	ModifyAuthenticationSubscription(ctx context.Context, ueId string) ApiModifyAuthenticationSubscriptionRequest

	// ModifyAuthenticationSubscriptionExecute executes the request
	//  @return PatchResult
	ModifyAuthenticationSubscriptionExecute(r ApiModifyAuthenticationSubscriptionRequest) (*PatchResult, *http.Response, error)
}

// AuthenticationSubscriptionDocumentApiService AuthenticationSubscriptionDocumentApi service
type AuthenticationSubscriptionDocumentApiService service

type ApiModifyAuthenticationSubscriptionRequest struct {
	ctx context.Context
	ApiService AuthenticationSubscriptionDocumentApi
	ueId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyAuthenticationSubscriptionRequest) PatchItem(patchItem []PatchItem) ApiModifyAuthenticationSubscriptionRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyAuthenticationSubscriptionRequest) SupportedFeatures(supportedFeatures string) ApiModifyAuthenticationSubscriptionRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyAuthenticationSubscriptionRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyAuthenticationSubscriptionExecute(r)
}

/*
ModifyAuthenticationSubscription modify the authentication subscription data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiModifyAuthenticationSubscriptionRequest
*/
func (a *AuthenticationSubscriptionDocumentApiService) ModifyAuthenticationSubscription(ctx context.Context, ueId string) ApiModifyAuthenticationSubscriptionRequest {
	return ApiModifyAuthenticationSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *AuthenticationSubscriptionDocumentApiService) ModifyAuthenticationSubscriptionExecute(r ApiModifyAuthenticationSubscriptionRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationSubscriptionDocumentApiService.ModifyAuthenticationSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/authentication-data/authentication-subscription"
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
	localVarHTTPHeaderAccepts := []string{"application/problem+json", "application/json"}

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
