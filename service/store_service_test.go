package service

import (
	"acronis/dto"
	"acronis/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type StoreServiceTestSuite struct {
	suite.Suite
	svc  StoreData
	repo repository.DataRepository
}

func TestStoreServiceTestSuite(t *testing.T) {
	suite.Run(t, &StoreServiceTestSuite{})
}

func (suite *StoreServiceTestSuite) SetupTest() {
	suite.repo = repository.NewStoreDataMap()

	//creating mock data to test
	mockdto := dto.StoredDataInfo{
		Key:            "key_1214",
		DataValue:      []interface{}{"hi", 1},
		CreationTime:   time.Now(),
		ExpirationTime: time.Now().Add(time.Duration(1) * time.Second),
	}
	suite.repo.StoreDataIntoDB(mockdto, "key_1214")
	suite.svc = NewInMemoryStore(suite.repo)
}

func (suite *StoreServiceTestSuite) TestPushDataToList() {
	expectedResp := dto.StoredDataInfo{
		Key:            "key_1214",
		DataValue:      []interface{}{"hi", 1, "ok"},
		CreationTime:   time.Now(),
		ExpirationTime: time.Now().Add(time.Duration(1) * time.Second),
	}
	suite.SetupTest()
	actualResp, actualErr := suite.svc.PushDataToList("key_1214", "ok")
	suite.Nil(actualErr)
	suite.Equal(expectedResp.DataValue, actualResp.DataValue)
	suite.Equal(expectedResp.Key, actualResp.Key)
}
