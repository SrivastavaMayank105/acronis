package client

import (
	"acronis/dto"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type DataStoreCLient interface {
	SetData(ctx context.Context, requestBody interface{}) (*dto.StoredDataInfo, error)
	GetAllData(ctx context.Context) ([]dto.StoredDataInfo, error)
	GetDataByKey(ctx context.Context, key string) (*dto.StoredDataInfo, error)
	UpdateData(ctx context.Context, key string, request interface{}) (*dto.StoredDataInfo, error)
	DeleteData(ctx context.Context, key string) error
	PushDataIntoList(ctx context.Context, key string, value interface{}) (*dto.StoredDataInfo, error)
	PopDataFromList(ctx context.Context, key string, value interface{}) (*dto.StoredDataInfo, error)
}

type dataStoreClient struct {
	HttpClinet *http.Client
	BaseUrl    string
}

func NewDataStoreClient(baseUrl string) DataStoreCLient {
	return &dataStoreClient{BaseUrl: baseUrl, HttpClinet: &http.Client{}}
}

func (dtc *dataStoreClient) SetData(ctx context.Context, request interface{}) (*dto.StoredDataInfo, error) {

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, dtc.BaseUrl+"/api/data", bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("failed to set data")
	}
	respBody, _ := io.ReadAll(resp.Body)
	var result dto.StoredDataInfo
	err = json.Unmarshal(respBody, &result)

	if err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	return &result, nil
}

func (dtc *dataStoreClient) GetAllData(ctx context.Context) ([]dto.StoredDataInfo, error) {
	req, err := http.NewRequest(http.MethodGet, dtc.BaseUrl+"/api/data", nil)

	if err != nil {
		return nil, err
	}

	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch all data")
	}

	respBody, _ := io.ReadAll(resp.Body)
	var result []dto.StoredDataInfo
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, errors.New("failed to unmarshal all data response")
	}

	return result, nil
}

func (dtc *dataStoreClient) GetDataByKey(ctx context.Context, key string) (*dto.StoredDataInfo, error) {
	req, err := http.NewRequest(http.MethodGet, dtc.BaseUrl+"/api/data/"+key, nil)

	if err != nil {
		return nil, err
	}

	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch data")
	}

	respBody, _ := io.ReadAll(resp.Body)
	var result dto.StoredDataInfo
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, errors.New("failed to unmarshal all data response")
	}

	return &result, nil
}

func (dtc *dataStoreClient) UpdateData(ctx context.Context, key string, request interface{}) (*dto.StoredDataInfo, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, dtc.BaseUrl+"/api/data/"+key, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to update data")
	}
	respBody, _ := io.ReadAll(resp.Body)
	var result dto.StoredDataInfo
	err = json.Unmarshal(respBody, &result)

	if err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	return &result, nil
}

func (dtc *dataStoreClient) DeleteData(ctx context.Context, key string) error {
	req, err := http.NewRequest(http.MethodDelete, dtc.BaseUrl+"/api/data/"+key, nil)

	if err != nil {
		return err
	}

	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to delete data")
	}

	return nil

}

func (dtc *dataStoreClient) PushDataIntoList(ctx context.Context, key string, value interface{}) (*dto.StoredDataInfo, error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, dtc.BaseUrl+"/api/data/"+key+"/push", bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to push value into list")
	}

	respBody, _ := io.ReadAll(resp.Body)
	var result dto.StoredDataInfo
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, errors.New("failed to unmarshal push response")
	}

	return &result, nil

}

func (dtc *dataStoreClient) PopDataFromList(ctx context.Context, key string, value interface{}) (*dto.StoredDataInfo, error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, dtc.BaseUrl+"/api/data/"+key+"/pop", bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	resp, err := dtc.HttpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to push value into list")
	}

	respBody, _ := io.ReadAll(resp.Body)
	var result dto.StoredDataInfo
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, errors.New("failed to unmarshal push response")
	}

	return &result, nil

}
