package com.smarthome.dm;

import com.smarthome.DevicesApi;
import com.smarthome.dm.mappers.DeviceMapper;
import com.smarthome.dm.mappers.ValueReader;
import com.smarthome.dm.service.DeviceService;
import com.smarthome.model.Device;
import com.smarthome.model.DeviceCreateRequest;
import com.smarthome.model.DevicePatchRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequiredArgsConstructor
public class DeviceController implements DevicesApi {
    private final DeviceService service;
    private final DeviceMapper mapper;
    private final ValueReader valueReader;

    @Override
    public ResponseEntity<Device> devicesDeviceIdGet(Integer deviceId) {
        return service.getById(deviceId)
                .map(mapper::map)
                .map(valueReader::enrich)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.status(404).body(null));
    }

    @Override
    public ResponseEntity<Device> devicesDeviceIdPatch(Integer deviceId, DevicePatchRequest devicePatchRequest) {
        com.smarthome.dm.db.Device mapped = mapper.map(devicePatchRequest);
        mapped.setId(deviceId);
        com.smarthome.dm.db.Device device = service.patchDevice(mapped);
        return ResponseEntity.ok(mapper.map(device));
    }

    @Override
    public ResponseEntity<Device> devicesDeviceIdPut(Integer deviceId, DeviceCreateRequest deviceCreateRequest) {
        com.smarthome.dm.db.Device mapped = mapper.map(deviceCreateRequest);
        mapped.setId(deviceId);
        com.smarthome.dm.db.Device updated = service.update(mapped);
        return ResponseEntity.ok(mapper.map(updated));
    }

    @Override
    public ResponseEntity<List<Device>> devicesGet() {
        List<Device> result = service.getAll().stream()
                .map(mapper::map)
                .map(valueReader::enrich)
                .toList();
        return ResponseEntity.ok(result);
    }

    @Override
    public ResponseEntity<Device> devicesPost(DeviceCreateRequest deviceCreateRequest) {
        com.smarthome.dm.db.Device sensor = service.createSensor(mapper.map(deviceCreateRequest));
        return ResponseEntity.status(HttpStatus.CREATED).body(mapper.map(sensor));
    }
}
