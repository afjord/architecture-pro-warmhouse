package com.smarthome.dm.db;

import jakarta.persistence.PrePersist;
import jakarta.persistence.PreUpdate;
import org.jspecify.annotations.NonNull;

import java.time.OffsetDateTime;

public class DeviceEntityListener {

    @PrePersist
    void prePersist(@NonNull Device device) {
        if (device.getStatus() == null) {
            device.setStatus("inactive");
        }
        device.setLastUpdated(OffsetDateTime.now());
        device.setCreatedAt(OffsetDateTime.now());
    }

    @PreUpdate
    void preUpdate(@NonNull Device device) {
        device.setLastUpdated(OffsetDateTime.now());
    }
}
