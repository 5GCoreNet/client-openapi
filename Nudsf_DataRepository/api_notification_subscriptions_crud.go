/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type NotificationSubscriptionsCRUDApi interface {

	/*
	GetNotificationSubscriptions Notification subscription retrieval

	retrieve all notification subscriptions of the storage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param realmId Identifier of the Realm
	@param storageId Identifier of the Storage
	@return ApiGetNotificationSubscriptionsRequest
	*/
	GetNotificationSubscriptions(ctx context.Context, realmId string, storageId string) ApiGetNotificationSubscriptionsRequest

	// GetNotificationSubscriptionsExecute executes the request
	//  @return []NotificationSubscription
	GetNotificationSubscriptionsExecute(r ApiGetNotificationSubscriptionsRequest) ([]NotificationSubscription, *http.Response, error)
}

// NotificationSubscriptionsCRUDApiService NotificationSubscriptionsCRUDApi service
type NotificationSubscriptionsCRUDApiService service

type ApiGetNotificationSubscriptionsRequest struct {
	ctx context.Context
	ApiService NotificationSubscriptionsCRUDApi
	realmId string
	storageId string
	limitRange *int32
	supportedFeatures *string
}

// The maximum number of NotificationSubscriptions to fetch
func (r ApiGetNotificationSubscriptionsRequest) LimitRange(limitRange int32) ApiGetNotificationSubscriptionsRequest {
	r.limitRange = &limitRange
	return r
}

// Features required to be supported by the target NF
func (r ApiGetNotificationSubscriptionsRequest) SupportedFeatures(supportedFeatures string) ApiGetNotificationSubscriptionsRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiGetNotificationSubscriptionsRequest) Execute() ([]NotificationSubscription, *http.Response, error) {
	return r.ApiService.GetNotificationSubscriptionsExecute(r)
}

/*
GetNotificationSubscriptions Notification subscription retrieval

retrieve all notification subscriptions of the storage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realmId Identifier of the Realm
 @param storageId Identifier of the Storage
 @return ApiGetNotificationSubscriptionsRequest
*/
func (a *NotificationSubscriptionsCRUDApiService) GetNotificationSubscriptions(ctx context.Context, realmId string, storageId string) ApiGetNotificationSubscriptionsRequest {
	return ApiGetNotificationSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		realmId: realmId,
		storageId: storageId,
	}
}

// Execute executes the request
//  @return []NotificationSubscription
func (a *NotificationSubscriptionsCRUDApiService) GetNotificationSubscriptionsExecute(r ApiGetNotificationSubscriptionsRequest) ([]NotificationSubscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []NotificationSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationSubscriptionsCRUDApiService.GetNotificationSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{realmId}/{storageId}/subs-to-notify"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", url.PathEscape(parameterToString(r.realmId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", url.PathEscape(parameterToString(r.storageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limitRange != nil {
		localVarQueryParams.Add("limit-range", parameterToString(*r.limitRange, ""))
	}
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
		if localVarHTTPResponse.StatusCode == 304 {
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
		if localVarHTTPResponse.StatusCode == 401 {
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
