/*
Niwmsc_SMService

SMS-IWMSC Short Message Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Niwmsc_SMService

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
)


type SendMOSMSMessageAndTheDeliveryReportApi interface {

	/*
	SendSMS Send SMS payload for a given UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param supi Subscription Permanent Identifier (SUPI)
	@return ApiSendSMSRequest
	*/
	SendSMS(ctx context.Context, supi string) ApiSendSMSRequest

	// SendSMSExecute executes the request
	//  @return SendSMS200Response
	SendSMSExecute(r ApiSendSMSRequest) (*SendSMS200Response, *http.Response, error)
}

// SendMOSMSMessageAndTheDeliveryReportApiService SendMOSMSMessageAndTheDeliveryReportApi service
type SendMOSMSMessageAndTheDeliveryReportApiService service

type ApiSendSMSRequest struct {
	ctx context.Context
	ApiService SendMOSMSMessageAndTheDeliveryReportApi
	supi string
	jsonData *SmsData
	binaryPayload **os.File
}

func (r ApiSendSMSRequest) JsonData(jsonData SmsData) ApiSendSMSRequest {
	r.jsonData = &jsonData
	return r
}

func (r ApiSendSMSRequest) BinaryPayload(binaryPayload *os.File) ApiSendSMSRequest {
	r.binaryPayload = &binaryPayload
	return r
}

func (r ApiSendSMSRequest) Execute() (*SendSMS200Response, *http.Response, error) {
	return r.ApiService.SendSMSExecute(r)
}

/*
SendSMS Send SMS payload for a given UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param supi Subscription Permanent Identifier (SUPI)
 @return ApiSendSMSRequest
*/
func (a *SendMOSMSMessageAndTheDeliveryReportApiService) SendSMS(ctx context.Context, supi string) ApiSendSMSRequest {
	return ApiSendSMSRequest{
		ApiService: a,
		ctx: ctx,
		supi: supi,
	}
}

// Execute executes the request
//  @return SendSMS200Response
func (a *SendMOSMSMessageAndTheDeliveryReportApiService) SendSMSExecute(r ApiSendSMSRequest) (*SendSMS200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendSMS200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendMOSMSMessageAndTheDeliveryReportApiService.SendSMS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mo-sm-infos/{supi}/sendsms"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", url.PathEscape(parameterToString(r.supi, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/related"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"multipart/related", "application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.jsonData != nil {
		paramJson, err := parameterToJson(*r.jsonData)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("jsonData", paramJson)
	}
	var binaryPayloadLocalVarFormFileName string
	var binaryPayloadLocalVarFileName     string
	var binaryPayloadLocalVarFileBytes    []byte

	binaryPayloadLocalVarFormFileName = "binaryPayload"

	var binaryPayloadLocalVarFile *os.File
	if r.binaryPayload != nil {
		binaryPayloadLocalVarFile = *r.binaryPayload
	}
	if binaryPayloadLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryPayloadLocalVarFile)
		binaryPayloadLocalVarFileBytes = fbs
		binaryPayloadLocalVarFileName = binaryPayloadLocalVarFile.Name()
		binaryPayloadLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryPayloadLocalVarFileBytes, fileName: binaryPayloadLocalVarFileName, formFileName: binaryPayloadLocalVarFormFileName})
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
