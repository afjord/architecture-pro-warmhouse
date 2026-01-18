package com.smarthome.dm.mappers;

import com.smarthome.dm.clients.SensorClient;
import com.smarthome.dm.clients.TemperatureData;
import com.smarthome.model.Device;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class ValueReader {
    private final SensorClient client;

    public Device enrich(Device device) {
        if (device == null) {
            return null;
        }
        if (!"temperature".equals(device.getType())) {
            return device;
        }
        TemperatureData temperatureById = client.getTemperatureById(device.getId().toString());
        device.setValue(temperatureById.getValue());
        device.setStatus("active");
        device.setLastUpdated(temperatureById.getTimestamp());
        return device;

    }
}
