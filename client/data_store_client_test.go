package client

import (
	"acronis/dto"
	"context"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
)

type DataStoreClientSuite struct {
	suite.Suite
	dataStoreClient DataStoreCLient
}

func TestDataStoreClientSuite(t *testing.T) {
	suite.Run(t, &DataStoreClientSuite{})
}

func (suite *DataStoreClientSuite) SetupTest() {
	suite.dataStoreClient = NewDataStoreClient("http://abc.com")
}

func (suite *DataStoreClientSuite) TestUpdateData() {
	testCases := []struct {
		Name         string
		initfunc     func()
		expectedResp *dto.StoredDataInfo
		requestData  interface{}
	}{
		{
			Name: "Happy path",
			initfunc: func() {
				gock.New("http://abc.com").
					Put("/api/data/Key_1234").
					BodyString(`"i have been updated"`).
					Reply(http.StatusOK).
					JSON(dto.StoredDataInfo{
						Key:       "Key_1234",
						DataValue: "hello",
					})
			},
			expectedResp: &dto.StoredDataInfo{
				Key:       "Key_1234",
				DataValue: "hello",
			},
			requestData: "i have been updated",
		},
	}

	for _, tcase := range testCases {
		suite.Run(tcase.Name, func() {
			defer gock.Off()

			tcase.initfunc()

			actualResp, actualErr := suite.dataStoreClient.UpdateData(context.TODO(), "Key_1234", tcase.requestData)

			suite.NoError(actualErr)
			suite.Equal(tcase.expectedResp, actualResp)
		})
	}
}
