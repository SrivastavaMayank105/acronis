package service

import (
	"acronis/dto"
	"acronis/mapper"
	"acronis/repository"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type StoreData interface {
	InsertData(data mapper.DataInfo) (dto.StoredDataInfo, error)
	GetAllData() ([]dto.StoredDataInfo, error)
	GetDataByKey(key string) (dto.StoredDataInfo, error)
	DeleteDataByKey(key string) error
	PushDataToList(key string, value interface{}) (dto.StoredDataInfo, error)
	PopDataFromList(key string, value interface{}) (dto.StoredDataInfo, error)
	UpdateDataByKey(key string, data mapper.DataInfo) (dto.StoredDataInfo, error)
}

type InMemmoryData struct {
	repository repository.DataRepository
}

func NewInMemoryStore(repo repository.DataRepository) StoreData {
	return &InMemmoryData{
		repository: repo,
	}
}

func (inMemmory *InMemmoryData) InsertData(data mapper.DataInfo) (dto.StoredDataInfo, error) {
	key := generateRandomKey()
	respDto := dto.StoredDataInfo{
		Key:            key,
		DataValue:      data.Data,
		CreationTime:   time.Now(),
		ExpirationTime: time.Now().Add(time.Duration(1) * time.Minute), // after 1 min this will get expire, and will get deleted from map
	}

	err := inMemmory.repository.StoreDataIntoDB(respDto, key)
	if err != nil {
		return dto.StoredDataInfo{}, errors.New("failed to insert data into db")
	}
	return respDto, nil
}

func (inMemmory *InMemmoryData) GetAllData() ([]dto.StoredDataInfo, error) {
	resp, err := inMemmory.repository.GetDataFromDB()
	if err != nil {
		return nil, errors.New("failed to retrieve data from the db")
	}

	respDto := []dto.StoredDataInfo{}
	for key, value := range resp {

		eachrespDto := dto.StoredDataInfo{}
		eachrespDto.Key = key
		eachrespDto.CreationTime = value.CreationTime
		eachrespDto.ExpirationTime = value.ExpireTime
		if value.DataType == "string" {
			eachrespDto.DataValue = value.StringValue
		} else {
			eachrespDto.DataValue = value.ListValue
		}

		respDto = append(respDto, eachrespDto)

	}
	return respDto, nil
}

func (inMemmory *InMemmoryData) GetDataByKey(key string) (dto.StoredDataInfo, error) {
	responseItem, err := inMemmory.repository.GetDataByKeyFromDB(key)
	if err != nil {
		return dto.StoredDataInfo{}, errors.New("failed to retrieve data from the db")
	}

	resp := MapDataFromRepoToDto(key, responseItem)

	return resp, nil
}

func (inMemmory *InMemmoryData) DeleteDataByKey(key string) error {
	err := inMemmory.repository.DeleteDataFromDB(key)
	if err != nil {
		return err
	}

	return nil
}

func (inMemmory *InMemmoryData) PushDataToList(key string, value interface{}) (dto.StoredDataInfo, error) {
	itemResp, err := inMemmory.repository.InsertDataIntoListInDB(key, value)
	if err != nil {
		return dto.StoredDataInfo{}, err
	}
	resp := MapDataFromRepoToDto(key, itemResp)

	return resp, nil
}

func (inMemmory *InMemmoryData) PopDataFromList(key string, value interface{}) (dto.StoredDataInfo, error) {

	itemResp, err := inMemmory.repository.DeleteDataFromListInDB(key, value)
	if err != nil {
		return dto.StoredDataInfo{}, err
	}
	resp := MapDataFromRepoToDto(key, itemResp)

	return resp, nil
}

func (inMemmory *InMemmoryData) UpdateDataByKey(key string, data mapper.DataInfo) (dto.StoredDataInfo, error) {
	itemResp, err := inMemmory.repository.UpdateDataIntoDB(key, data)
	if err != nil {
		return dto.StoredDataInfo{}, err
	}
	resp := MapDataFromRepoToDto(key, itemResp)
	return resp, nil

}

// TODO: good to use hash algo
func generateRandomKey() string {
	return "key_" + strconv.Itoa(rand.Intn(1000000))
}

func MapDataFromRepoToDto(key string, item repository.Item) dto.StoredDataInfo {
	respDto := dto.StoredDataInfo{}
	respDto.Key = key
	respDto.CreationTime = item.CreationTime
	respDto.ExpirationTime = item.ExpireTime
	if item.DataType == "string" {
		respDto.DataValue = item.StringValue
	} else {
		respDto.DataValue = item.ListValue
	}

	return respDto
}
