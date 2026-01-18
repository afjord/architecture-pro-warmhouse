package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"smarthome/models"
)

type DevicesAPIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewDevicesAPIClient(baseURL string) *DevicesAPIClient {
	return &DevicesAPIClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *DevicesAPIClient) GetDevices(ctx context.Context) ([]models.Device, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/devices",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sensor api returned %d", resp.StatusCode)
	}

	var devices []models.Device
	return devices, json.NewDecoder(resp.Body).Decode(&devices)
}

func (c *DevicesAPIClient) GetSensorByID(ctx context.Context, id int) (*models.Device, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/devices/%d", c.baseURL, id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sensor api returned %d", resp.StatusCode)
	}

	var sensor models.Device
	return &sensor, json.NewDecoder(resp.Body).Decode(&sensor)
}

func (c *DevicesAPIClient) CreateSensor(
	ctx context.Context,
	reqBody models.SensorCreate,
) (*models.Device, error) {

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/devices",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("sensor api returned %d", resp.StatusCode)
	}

	var sensor models.Device
	return &sensor, json.NewDecoder(resp.Body).Decode(&sensor)
}

func (c *DevicesAPIClient) UpdateSensor(
	ctx context.Context,
	id int,
	reqBody models.SensorUpdate,
) (*models.Device, error) {

	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s/devices/%d", c.baseURL, id),
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sensor api returned %d", resp.StatusCode)
	}

	var sensor models.Device
	return &sensor, json.NewDecoder(resp.Body).Decode(&sensor)
}

func (c *DevicesAPIClient) DeleteSensor(ctx context.Context, id int) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/devices/%d", c.baseURL, id),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("sensor api returned %d", resp.StatusCode)
	}

	return nil
}

func (c *DevicesAPIClient) UpdateSensorValue(
	ctx context.Context,
	id int,
	value float64,
	status string,
) error {

	body, _ := json.Marshal(map[string]interface{}{
		"value":  value,
		"status": status,
	})

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("%s/devices/%d/value", c.baseURL, id),
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("sensor api returned %d", resp.StatusCode)
	}

	return nil
}
