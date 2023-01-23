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
	"strings"
)


type RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi interface {

	/*
	QueryPeiInformation Retrieves the PEI Information of the 5GC/EPC domains

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiQueryPeiInformationRequest
	*/
	QueryPeiInformation(ctx context.Context, ueId string) ApiQueryPeiInformationRequest

	// QueryPeiInformationExecute executes the request
	//  @return PeiUpdateInfo
	QueryPeiInformationExecute(r ApiQueryPeiInformationRequest) (*PeiUpdateInfo, *http.Response, error)
}

// RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApiService RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi service
type RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApiService service

type ApiQueryPeiInformationRequest struct {
	ctx context.Context
	ApiService RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi
	ueId string
}

func (r ApiQueryPeiInformationRequest) Execute() (*PeiUpdateInfo, *http.Response, error) {
	return r.ApiService.QueryPeiInformationExecute(r)
}

/*
QueryPeiInformation Retrieves the PEI Information of the 5GC/EPC domains

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiQueryPeiInformationRequest
*/
func (a *RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApiService) QueryPeiInformation(ctx context.Context, ueId string) ApiQueryPeiInformationRequest {
	return ApiQueryPeiInformationRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return PeiUpdateInfo
func (a *RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApiService) QueryPeiInformationExecute(r ApiQueryPeiInformationRequest) (*PeiUpdateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PeiUpdateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApiService.QueryPeiInformation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/pei-info"
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
