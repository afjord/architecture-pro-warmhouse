package com.smarthome.dm.service;

import com.smarthome.dm.db.Device;
import com.smarthome.dm.db.SensorRepository;
import jakarta.persistence.EntityNotFoundException;
import lombok.RequiredArgsConstructor;
import org.jspecify.annotations.NonNull;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class DeviceService {
    private final SensorRepository sensorRepository;

    public List<Device> getAll() {
        return sensorRepository.findAll();
    }

    public Optional<Device> getById(@NonNull Integer id) {
        return sensorRepository.findById(id);
    }

    public Device createSensor(Device createDto) {
        return sensorRepository.save(createDto);
    }

    public Device patchDevice(@NonNull Device device) {
        Device existed = getById(device.getId()).orElseThrow(EntityNotFoundException::new);
        existed.setValue(device.getValue());
        existed.setStatus(device.getStatus());
        return sensorRepository.save(existed);
    }

    public void deleteSensor(Integer id) {
        sensorRepository.deleteById(id);
    }

    public @NonNull Device update(@NonNull Device device) {
        Device existed = getById(device.getId()).orElseThrow(EntityNotFoundException::new);
        if (device.getName() != null) {
            existed.setName(device.getName());
        }
        if (device.getType() != null) {
            existed.setType(device.getType());
        }
        if (device.getLocation() != null) {
            existed.setLocation(device.getLocation());
        }
        if (device.getUnit() != null) {
            existed.setUnit(device.getUnit());
        }
        return sensorRepository.save(existed);
    }
}