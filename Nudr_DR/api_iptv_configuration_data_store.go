/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type IPTVConfigurationDataStoreApi interface {

	/*
	ReadIPTVCongifurationData Retrieve IPTV configuration Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReadIPTVCongifurationDataRequest
	*/
	ReadIPTVCongifurationData(ctx context.Context) ApiReadIPTVCongifurationDataRequest

	// ReadIPTVCongifurationDataExecute executes the request
	//  @return []IptvConfigData
	ReadIPTVCongifurationDataExecute(r ApiReadIPTVCongifurationDataRequest) ([]IptvConfigData, *http.Response, error)
}

// IPTVConfigurationDataStoreApiService IPTVConfigurationDataStoreApi service
type IPTVConfigurationDataStoreApiService service

type ApiReadIPTVCongifurationDataRequest struct {
	ctx context.Context
	ApiService IPTVConfigurationDataStoreApi
	configIds *[]string
	dnns *[]string
	snssais *[]Snssai
	supis *[]string
	interGroupIds *[]string
}

// Each element identifies a configuration.
func (r ApiReadIPTVCongifurationDataRequest) ConfigIds(configIds []string) ApiReadIPTVCongifurationDataRequest {
	r.configIds = &configIds
	return r
}

// Each element identifies a DNN.
func (r ApiReadIPTVCongifurationDataRequest) Dnns(dnns []string) ApiReadIPTVCongifurationDataRequest {
	r.dnns = &dnns
	return r
}

// Each element identifies a slice.
func (r ApiReadIPTVCongifurationDataRequest) Snssais(snssais []Snssai) ApiReadIPTVCongifurationDataRequest {
	r.snssais = &snssais
	return r
}

// Each element identifies the user.
func (r ApiReadIPTVCongifurationDataRequest) Supis(supis []string) ApiReadIPTVCongifurationDataRequest {
	r.supis = &supis
	return r
}

// Each element identifies a group of users.
func (r ApiReadIPTVCongifurationDataRequest) InterGroupIds(interGroupIds []string) ApiReadIPTVCongifurationDataRequest {
	r.interGroupIds = &interGroupIds
	return r
}

func (r ApiReadIPTVCongifurationDataRequest) Execute() ([]IptvConfigData, *http.Response, error) {
	return r.ApiService.ReadIPTVCongifurationDataExecute(r)
}

/*
ReadIPTVCongifurationData Retrieve IPTV configuration Data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReadIPTVCongifurationDataRequest
*/
func (a *IPTVConfigurationDataStoreApiService) ReadIPTVCongifurationData(ctx context.Context) ApiReadIPTVCongifurationDataRequest {
	return ApiReadIPTVCongifurationDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []IptvConfigData
func (a *IPTVConfigurationDataStoreApiService) ReadIPTVCongifurationDataExecute(r ApiReadIPTVCongifurationDataRequest) ([]IptvConfigData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []IptvConfigData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IPTVConfigurationDataStoreApiService.ReadIPTVCongifurationData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application-data/iptvConfigData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configIds != nil {
		localVarQueryParams.Add("config-ids", parameterToString(*r.configIds, "csv"))
	}
	if r.dnns != nil {
		localVarQueryParams.Add("dnns", parameterToString(*r.dnns, "csv"))
	}
	if r.snssais != nil {
		localVarQueryParams.Add("snssais", parameterToString(*r.snssais, "csv"))
	}
	if r.supis != nil {
		localVarQueryParams.Add("supis", parameterToString(*r.supis, "csv"))
	}
	if r.interGroupIds != nil {
		localVarQueryParams.Add("inter-group-ids", parameterToString(*r.interGroupIds, "csv"))
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
