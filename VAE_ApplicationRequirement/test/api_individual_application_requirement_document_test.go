/*
VAE_ApplicationRequirement

Testing IndividualApplicationRequirementDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package VAE_ApplicationRequirement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_VAE_ApplicationRequirement_IndividualApplicationRequirementDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualApplicationRequirementDocumentApiService DeleteApplicationRequirement", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var requirementId string

        resp, httpRes, err := apiClient.IndividualApplicationRequirementDocumentApi.DeleteApplicationRequirement(context.Background(), requirementId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationRequirementDocumentApiService ReadApplicationRequirement", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var requirementId string

        resp, httpRes, err := apiClient.IndividualApplicationRequirementDocumentApi.ReadApplicationRequirement(context.Background(), requirementId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
