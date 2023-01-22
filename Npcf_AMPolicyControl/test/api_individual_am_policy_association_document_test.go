/*
Npcf_AMPolicyControl

Testing IndividualAMPolicyAssociationDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Npcf_AMPolicyControl

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Npcf_AMPolicyControl_IndividualAMPolicyAssociationDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualAMPolicyAssociationDocumentApiService DeleteIndividualAMPolicyAssociation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var polAssoId string

        resp, httpRes, err := apiClient.IndividualAMPolicyAssociationDocumentApi.DeleteIndividualAMPolicyAssociation(context.Background(), polAssoId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualAMPolicyAssociationDocumentApiService ReadIndividualAMPolicyAssociation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var polAssoId string

        resp, httpRes, err := apiClient.IndividualAMPolicyAssociationDocumentApi.ReadIndividualAMPolicyAssociation(context.Background(), polAssoId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualAMPolicyAssociationDocumentApiService ReportObservedEventTriggersForIndividualAMPolicyAssociation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var polAssoId string

        resp, httpRes, err := apiClient.IndividualAMPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualAMPolicyAssociation(context.Background(), polAssoId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}