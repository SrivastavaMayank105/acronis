package controller

import (
	"acronis/dto"
	"acronis/mapper"
	"acronis/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	svc            *mocks.MockStoreData
	dataController Controller
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.svc = mocks.NewMockStoreData(suite.T())
	suite.dataController = *NewController(suite.svc)
}
func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, &ControllerTestSuite{})
}

func (suite *ControllerTestSuite) TestSetData() {
	testCases := []struct {
		Name         string
		requestBody  mapper.DataInfo
		initfunc     func()
		expectedResp dto.StoredDataInfo
		statusCode   int
	}{
		{
			Name: "Successfully inserted the data",
			requestBody: mapper.DataInfo{
				Data: []string{"hello", "how are you"},
			},
			initfunc: func() {
				suite.svc.EXPECT().InsertData(mock.MatchedBy(func(arg mapper.DataInfo) bool {
					i := arg.Data.([]interface{})
					return i[0] == "hello"
				})).RunAndReturn(func(di mapper.DataInfo) (dto.StoredDataInfo, error) {
					return dto.StoredDataInfo{
						DataValue: []string{"hello", "how are you"},
					}, nil
				})
			},
			expectedResp: dto.StoredDataInfo{
				DataValue: []interface{}{"hello", "how are you"},
			},
			statusCode: http.StatusCreated,
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.Name, func() {
			defer gock.Off()
			suite.SetupTest()
			testCase.initfunc()
			reqBody, _ := json.Marshal(testCase.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/api/data", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			router := gin.Default()
			router.POST("/api/data", suite.dataController.SetData)
			router.ServeHTTP(recorder, req)
			suite.Equal(testCase.statusCode, recorder.Code)

			var actualResp dto.StoredDataInfo
			err := json.Unmarshal(recorder.Body.Bytes(), &actualResp)
			suite.NoError(err)
			suite.Equal(testCase.expectedResp.DataValue, actualResp.DataValue)

		})
	}
}
