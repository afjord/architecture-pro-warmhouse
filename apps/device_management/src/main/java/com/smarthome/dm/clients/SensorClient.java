package com.smarthome.dm.clients;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

@Component
public class SensorClient {

    @Value("${temperature.api.url}")
    private String apiUrl;

    private final RestTemplate restTemplate = new RestTemplate();

    public TemperatureData getTemperatureById(String sensorId) {
        try {
            return restTemplate.getForObject(apiUrl + "/temperature/" + sensorId, TemperatureData.class);
        } catch (RestClientException e) {
            throw new RuntimeException("Failed to fetch temperature data for sensor ID: " + sensorId, e);
        }
    }

    public TemperatureData getTemperatureByLocation(String location) {
        try {
            return restTemplate.getForObject(apiUrl + "/temperature/location/" + location, TemperatureData.class);
        } catch (RestClientException e) {
            throw new RuntimeException("Failed to fetch temperature for location: " + location, e);
        }
    }
}