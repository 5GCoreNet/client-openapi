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


type SMFEventGroupSubscriptionInfoDocumentApi interface {

	/*
	CreateSmfGroupSubscriptions Create SMF Subscription Info for a group of UEs or any YE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueGroupId
	@param subsId
	@return ApiCreateSmfGroupSubscriptionsRequest
	*/
	CreateSmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) ApiCreateSmfGroupSubscriptionsRequest

	// CreateSmfGroupSubscriptionsExecute executes the request
	//  @return SmfSubscriptionInfo
	CreateSmfGroupSubscriptionsExecute(r ApiCreateSmfGroupSubscriptionsRequest) (*SmfSubscriptionInfo, *http.Response, error)
}

// SMFEventGroupSubscriptionInfoDocumentApiService SMFEventGroupSubscriptionInfoDocumentApi service
type SMFEventGroupSubscriptionInfoDocumentApiService service

type ApiCreateSmfGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService SMFEventGroupSubscriptionInfoDocumentApi
	ueGroupId string
	subsId string
	smfSubscriptionInfo *SmfSubscriptionInfo
}

func (r ApiCreateSmfGroupSubscriptionsRequest) SmfSubscriptionInfo(smfSubscriptionInfo SmfSubscriptionInfo) ApiCreateSmfGroupSubscriptionsRequest {
	r.smfSubscriptionInfo = &smfSubscriptionInfo
	return r
}

func (r ApiCreateSmfGroupSubscriptionsRequest) Execute() (*SmfSubscriptionInfo, *http.Response, error) {
	return r.ApiService.CreateSmfGroupSubscriptionsExecute(r)
}

/*
CreateSmfGroupSubscriptions Create SMF Subscription Info for a group of UEs or any YE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId
 @return ApiCreateSmfGroupSubscriptionsRequest
*/
func (a *SMFEventGroupSubscriptionInfoDocumentApiService) CreateSmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) ApiCreateSmfGroupSubscriptionsRequest {
	return ApiCreateSmfGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return SmfSubscriptionInfo
func (a *SMFEventGroupSubscriptionInfoDocumentApiService) CreateSmfGroupSubscriptionsExecute(r ApiCreateSmfGroupSubscriptionsRequest) (*SmfSubscriptionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmfSubscriptionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFEventGroupSubscriptionInfoDocumentApiService.CreateSmfGroupSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterToString(r.ueGroupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterToString(r.subsId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.smfSubscriptionInfo == nil {
		return localVarReturnValue, nil, reportError("smfSubscriptionInfo is required and must be specified")
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
	localVarPostBody = r.smfSubscriptionInfo
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
