package com.smarthome.dm.mappers;

import com.smarthome.dm.db.Device;
import com.smarthome.model.DeviceCreateRequest;
import com.smarthome.model.DevicePatchRequest;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper
public abstract class DeviceMapper {

    @Mapping(target = "value", ignore = true)
    @Mapping(target = "homeId", ignore = true)
    @Mapping(target = "status", ignore = true)
    @Mapping(target = "lastUpdated", ignore = true)
    @Mapping(target = "id", ignore = true)
    @Mapping(target = "createdAt", ignore = true)
    public abstract Device map(DeviceCreateRequest source);

    @Mapping(target = "unit", ignore = true)
    @Mapping(target = "type", ignore = true)
    @Mapping(target = "name", ignore = true)
    @Mapping(target = "location", ignore = true)
    @Mapping(target = "lastUpdated", ignore = true)
    @Mapping(target = "id", ignore = true)
    @Mapping(target = "homeId", ignore = true)
    @Mapping(target = "createdAt", ignore = true)
    public abstract Device map(DevicePatchRequest source);

    public abstract com.smarthome.model.Device map(Device device);
}
