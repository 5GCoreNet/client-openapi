/*
MSGS_ASRegistration

Testing ASRegistrationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package MSGS_ASRegistration

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_MSGS_ASRegistration_ASRegistrationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ASRegistrationApiService RegistrationsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ASRegistrationApi.RegistrationsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}