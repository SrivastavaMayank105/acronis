package repository

import (
	"acronis/dto"
	"acronis/mapper"
	"errors"
	"sync"
	"time"
)

// Mocking the DB , using map for the DB
// TODO: implement the actual DB here

type DataRepository interface {
	StoreDataIntoDB(data dto.StoredDataInfo, key string) error
	DeleteDataFromDB(key string) error
	UpdateDataIntoDB(key string, data mapper.DataInfo) (Item, error)
	GetDataFromDB() (map[string]Item, error)
	GetDataByKeyFromDB(key string) (Item, error)
	DeleteDataFromListInDB(key string, value interface{}) (Item, error)
	InsertDataIntoListInDB(key string, value interface{}) (Item, error)
	StartCleanupJob()
}

type StoreDataMap struct {
	mockDB map[string]Item
	mutex  sync.RWMutex
}

type Item struct {
	DataType     string
	StringValue  string
	ListValue    []interface{}
	CreationTime time.Time
	ExpireTime   time.Time
}

func NewStoreDataMap() DataRepository {
	return &StoreDataMap{
		mockDB: make(map[string]Item),
	}
}

func (repo *StoreDataMap) StoreDataIntoDB(data dto.StoredDataInfo, key string) error {

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	item := Item{}
	item.CreationTime = data.CreationTime
	item.ExpireTime = data.ExpirationTime

	switch val := data.DataValue.(type) {
	case string:
		item.DataType = "string"
		item.StringValue = val
	case []interface{}:
		item.DataType = "list"
		item.ListValue = val
	default:
		item.DataType = "unsupported"
		return errors.New("unsupported data type")
	}
	repo.mockDB[key] = item

	//TODO: if connection with db fails, return the error

	return nil

}

func (repo *StoreDataMap) DeleteDataFromDB(key string) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	_, ok := repo.mockDB[key]
	if ok {
		delete(repo.mockDB, key)
	} else {
		return errors.New("key does not exists")
	}
	return nil
}

func (repo *StoreDataMap) UpdateDataIntoDB(key string, data mapper.DataInfo) (Item, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	item, ok := repo.mockDB[key]
	if !ok {
		return Item{}, errors.New("key does not exist")
	}

	if item.DataType == "string" {
		item.StringValue = data.Data.(string)
	} else if item.DataType == "list" {
		item.ListValue = data.Data.([]interface{})
	} else {
		return Item{}, errors.New("unsupported data type")
	}

	repo.mockDB[key] = item

	return repo.mockDB[key], nil
}

func (repo *StoreDataMap) GetDataFromDB() (map[string]Item, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	//TODO: if connection with db fails, return the error
	return repo.mockDB, nil
}

func (repo *StoreDataMap) GetDataByKeyFromDB(key string) (Item, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	//TODO: if connection with db fails, return the error

	_, ok := repo.mockDB[key]
	if !ok {
		return Item{}, errors.New("key does not exists")
	}
	return repo.mockDB[key], nil
}

func (repo *StoreDataMap) DeleteDataFromListInDB(key string, value interface{}) (Item, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	itemValue, ok := repo.mockDB[key]
	if !ok {
		return Item{}, errors.New("key does not exists")
	}

	if itemValue.DataType != "list" {
		return Item{}, errors.New("incorrect data type passed")
	}

	index := -1
	for ind, val := range itemValue.ListValue {
		if value == val {
			index = ind
			break
		}
	}

	if index == -1 {
		return Item{}, errors.New("values does not exists in list")
	}

	itemValue.ListValue = append(itemValue.ListValue[:index], itemValue.ListValue[index+1:]...)
	repo.mockDB[key] = itemValue
	return itemValue, nil
}
func (repo *StoreDataMap) InsertDataIntoListInDB(key string, value interface{}) (Item, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	itemValue, ok := repo.mockDB[key]
	if !ok {
		return Item{}, errors.New("key does not exists")
	}

	if itemValue.DataType != "list" {
		return Item{}, errors.New("incorrect data type passed")
	}

	itemValue.ListValue = append(itemValue.ListValue, value)
	repo.mockDB[key] = itemValue
	return itemValue, nil
}

// creating a  function which will delete the data from this map after every 5 sec

func (repo *StoreDataMap) StartCleanupJob() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		repo.mutex.Lock()
		now := time.Now()
		for key, item := range repo.mockDB {
			if item.ExpireTime.Before(now) {
				delete(repo.mockDB, key)
			}
		}
		repo.mutex.Unlock()
	}
}
