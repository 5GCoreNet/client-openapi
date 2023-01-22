# Go API client for Nnef_SMContext

Nnef SMContext Service.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import Nnef_SMContext "//"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), Nnef_SMContext.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), Nnef_SMContext.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), Nnef_SMContext.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), Nnef_SMContext.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nnef-smcontext/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualSMContextDocumentApi* | [**Delete**](docs/IndividualSMContextDocumentApi.md#delete) | **Post** /sm-contexts/{smContextId}/release | Delete SM Context
*IndividualSMContextDocumentApi* | [**Deliver**](docs/IndividualSMContextDocumentApi.md#deliver) | **Post** /sm-contexts/{smContextId}/deliver | Deliver Uplink MO Data
*IndividualSMContextDocumentApi* | [**Update**](docs/IndividualSMContextDocumentApi.md#update) | **Post** /sm-contexts/{smContextId}/update | Update SM Context
*SMContextsCollectionCollectionApi* | [**Create**](docs/SMContextsCollectionCollectionApi.md#create) | **Post** /sm-contexts | Create SM Context


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [ApnRateStatus](docs/ApnRateStatus.md)
 - [DeliverReqData](docs/DeliverReqData.md)
 - [DeliverRequest](docs/DeliverRequest.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [NFType](docs/NFType.md)
 - [NFTypeAnyOf](docs/NFTypeAnyOf.md)
 - [NiddInformation](docs/NiddInformation.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RefToBinaryData](docs/RefToBinaryData.md)
 - [ReleaseCause](docs/ReleaseCause.md)
 - [ReleaseCauseAnyOf](docs/ReleaseCauseAnyOf.md)
 - [SmContextConfiguration](docs/SmContextConfiguration.md)
 - [SmContextCreateData](docs/SmContextCreateData.md)
 - [SmContextCreatedData](docs/SmContextCreatedData.md)
 - [SmContextReleaseData](docs/SmContextReleaseData.md)
 - [SmContextReleasedData](docs/SmContextReleasedData.md)
 - [SmContextStatus](docs/SmContextStatus.md)
 - [SmContextStatusAnyOf](docs/SmContextStatusAnyOf.md)
 - [SmContextStatusNotification](docs/SmContextStatusNotification.md)
 - [SmContextUpdateData](docs/SmContextUpdateData.md)
 - [SmallDataRateControl](docs/SmallDataRateControl.md)
 - [SmallDataRateControlTimeUnit](docs/SmallDataRateControlTimeUnit.md)
 - [SmallDataRateControlTimeUnitAnyOf](docs/SmallDataRateControlTimeUnitAnyOf.md)
 - [SmallDataRateStatus](docs/SmallDataRateStatus.md)
 - [Snssai](docs/Snssai.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nnef-smcontext**: Access to the Nnef_SMContext API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


