/*
CAPIF_Discover_Service_API

API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Discover_Service_API

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiAllServiceAPIsGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	apiInvokerId *string
	apiName *string
	apiVersion *string
	commType *CommunicationType
	protocol *Protocol
	aefId *string
	dataFormat *DataFormat
	apiCat *string
	preferredAefLoc *AefLocation
	supportedFeatures *string
	apiSupportedFeatures *string
}

// String identifying the API invoker assigned by the CAPIF core function. It also represents the CCF identifier in the CAPIF-6/6e interface. 
func (r ApiAllServiceAPIsGetRequest) ApiInvokerId(apiInvokerId string) ApiAllServiceAPIsGetRequest {
	r.apiInvokerId = &apiInvokerId
	return r
}

// API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122. 
func (r ApiAllServiceAPIsGetRequest) ApiName(apiName string) ApiAllServiceAPIsGetRequest {
	r.apiName = &apiName
	return r
}

// API major version the URI (e.g. v1).
func (r ApiAllServiceAPIsGetRequest) ApiVersion(apiVersion string) ApiAllServiceAPIsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

// Communication type used by the API (e.g. REQUEST_RESPONSE).
func (r ApiAllServiceAPIsGetRequest) CommType(commType CommunicationType) ApiAllServiceAPIsGetRequest {
	r.commType = &commType
	return r
}

// Protocol used by the API.
func (r ApiAllServiceAPIsGetRequest) Protocol(protocol Protocol) ApiAllServiceAPIsGetRequest {
	r.protocol = &protocol
	return r
}

// AEF identifer.
func (r ApiAllServiceAPIsGetRequest) AefId(aefId string) ApiAllServiceAPIsGetRequest {
	r.aefId = &aefId
	return r
}

// Data formats used by the API (e.g. serialization protocol JSON used).
func (r ApiAllServiceAPIsGetRequest) DataFormat(dataFormat DataFormat) ApiAllServiceAPIsGetRequest {
	r.dataFormat = &dataFormat
	return r
}

// The service API category to which the service API belongs to.
func (r ApiAllServiceAPIsGetRequest) ApiCat(apiCat string) ApiAllServiceAPIsGetRequest {
	r.apiCat = &apiCat
	return r
}

// The preferred AEF location.
func (r ApiAllServiceAPIsGetRequest) PreferredAefLoc(preferredAefLoc AefLocation) ApiAllServiceAPIsGetRequest {
	r.preferredAefLoc = &preferredAefLoc
	return r
}

// Features supported by the NF consumer for the CAPIF Discover Service API.
func (r ApiAllServiceAPIsGetRequest) SupportedFeatures(supportedFeatures string) ApiAllServiceAPIsGetRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Features supported by the discovered service API indicated by api-name parameter. This may only be present if api-name query parameter is present. 
func (r ApiAllServiceAPIsGetRequest) ApiSupportedFeatures(apiSupportedFeatures string) ApiAllServiceAPIsGetRequest {
	r.apiSupportedFeatures = &apiSupportedFeatures
	return r
}

func (r ApiAllServiceAPIsGetRequest) Execute() (*DiscoveredAPIs, *http.Response, error) {
	return r.ApiService.AllServiceAPIsGetExecute(r)
}

/*
AllServiceAPIsGet Method for AllServiceAPIsGet

Discover published service APIs and retrieve a collection of APIs according to certain filter criteria.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAllServiceAPIsGetRequest
*/
func (a *DefaultApiService) AllServiceAPIsGet(ctx context.Context) ApiAllServiceAPIsGetRequest {
	return ApiAllServiceAPIsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DiscoveredAPIs
func (a *DefaultApiService) AllServiceAPIsGetExecute(r ApiAllServiceAPIsGetRequest) (*DiscoveredAPIs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscoveredAPIs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AllServiceAPIsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/allServiceAPIs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiInvokerId == nil {
		return localVarReturnValue, nil, reportError("apiInvokerId is required and must be specified")
	}

	localVarQueryParams.Add("api-invoker-id", parameterToString(*r.apiInvokerId, ""))
	if r.apiName != nil {
		localVarQueryParams.Add("api-name", parameterToString(*r.apiName, ""))
	}
	if r.apiVersion != nil {
		localVarQueryParams.Add("api-version", parameterToString(*r.apiVersion, ""))
	}
	if r.commType != nil {
		localVarQueryParams.Add("comm-type", parameterToString(*r.commType, ""))
	}
	if r.protocol != nil {
		localVarQueryParams.Add("protocol", parameterToString(*r.protocol, ""))
	}
	if r.aefId != nil {
		localVarQueryParams.Add("aef-id", parameterToString(*r.aefId, ""))
	}
	if r.dataFormat != nil {
		localVarQueryParams.Add("data-format", parameterToString(*r.dataFormat, ""))
	}
	if r.apiCat != nil {
		localVarQueryParams.Add("api-cat", parameterToString(*r.apiCat, ""))
	}
	if r.preferredAefLoc != nil {
		localVarQueryParams.Add("preferred-aef-loc", parameterToString(*r.preferredAefLoc, ""))
	}
	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
	if r.apiSupportedFeatures != nil {
		localVarQueryParams.Add("api-supported-features", parameterToString(*r.apiSupportedFeatures, ""))
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
		if localVarHTTPResponse.StatusCode == 406 {
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
		if localVarHTTPResponse.StatusCode == 414 {
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
		if localVarHTTPResponse.StatusCode == 429 {
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
