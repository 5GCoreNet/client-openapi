/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// SubsToNotifyDocumentApiService SubsToNotifyDocumentApi service
type SubsToNotifyDocumentApiService service

type ApiModifysubscriptionDataSubscriptionRequest struct {
	ctx context.Context
	ApiService *SubsToNotifyDocumentApiService
	subsId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifysubscriptionDataSubscriptionRequest) PatchItem(patchItem []PatchItem) ApiModifysubscriptionDataSubscriptionRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifysubscriptionDataSubscriptionRequest) SupportedFeatures(supportedFeatures string) ApiModifysubscriptionDataSubscriptionRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifysubscriptionDataSubscriptionRequest) Execute() (*ModifysubscriptionDataSubscription200Response, *http.Response, error) {
	return r.ApiService.ModifysubscriptionDataSubscriptionExecute(r)
}

/*
ModifysubscriptionDataSubscription Modify an individual subscriptionDataSubscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subsId
 @return ApiModifysubscriptionDataSubscriptionRequest
*/
func (a *SubsToNotifyDocumentApiService) ModifysubscriptionDataSubscription(ctx context.Context, subsId string) ApiModifysubscriptionDataSubscriptionRequest {
	return ApiModifysubscriptionDataSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return ModifysubscriptionDataSubscription200Response
func (a *SubsToNotifyDocumentApiService) ModifysubscriptionDataSubscriptionExecute(r ApiModifysubscriptionDataSubscriptionRequest) (*ModifysubscriptionDataSubscription200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModifysubscriptionDataSubscription200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubsToNotifyDocumentApiService.ModifysubscriptionDataSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/subs-to-notify/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterToString(r.subsId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}
	if len(*r.patchItem) < 1 {
		return localVarReturnValue, nil, reportError("patchItem must have at least 1 elements")
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

type ApiQuerySubscriptionDataSubscriptionsRequest struct {
	ctx context.Context
	ApiService *SubsToNotifyDocumentApiService
	subsId string
}

func (r ApiQuerySubscriptionDataSubscriptionsRequest) Execute() ([]SubscriptionDataSubscriptions, *http.Response, error) {
	return r.ApiService.QuerySubscriptionDataSubscriptionsExecute(r)
}

/*
QuerySubscriptionDataSubscriptions Retrieves a individual subscriptionDataSubscription identified by subsId

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subsId Unique ID of the subscription to retrieve
 @return ApiQuerySubscriptionDataSubscriptionsRequest
*/
func (a *SubsToNotifyDocumentApiService) QuerySubscriptionDataSubscriptions(ctx context.Context, subsId string) ApiQuerySubscriptionDataSubscriptionsRequest {
	return ApiQuerySubscriptionDataSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return []SubscriptionDataSubscriptions
func (a *SubsToNotifyDocumentApiService) QuerySubscriptionDataSubscriptionsExecute(r ApiQuerySubscriptionDataSubscriptionsRequest) ([]SubscriptionDataSubscriptions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SubscriptionDataSubscriptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubsToNotifyDocumentApiService.QuerySubscriptionDataSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/subs-to-notify/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterToString(r.subsId, "")), -1)

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

type ApiRemovesubscriptionDataSubscriptionsRequest struct {
	ctx context.Context
	ApiService *SubsToNotifyDocumentApiService
	subsId string
}

func (r ApiRemovesubscriptionDataSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemovesubscriptionDataSubscriptionsExecute(r)
}

/*
RemovesubscriptionDataSubscriptions Deletes a subscriptionDataSubscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subsId Unique ID of the subscription to remove
 @return ApiRemovesubscriptionDataSubscriptionsRequest
*/
func (a *SubsToNotifyDocumentApiService) RemovesubscriptionDataSubscriptions(ctx context.Context, subsId string) ApiRemovesubscriptionDataSubscriptionsRequest {
	return ApiRemovesubscriptionDataSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *SubsToNotifyDocumentApiService) RemovesubscriptionDataSubscriptionsExecute(r ApiRemovesubscriptionDataSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubsToNotifyDocumentApiService.RemovesubscriptionDataSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/subs-to-notify/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterToString(r.subsId, "")), -1)

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
