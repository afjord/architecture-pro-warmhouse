package org.example;

import org.jspecify.annotations.NonNull;

import java.time.OffsetDateTime;
import java.util.concurrent.ThreadLocalRandom;

public record TemperatureData(
        double value,
        String unit,
        OffsetDateTime timestamp,
        String location,
        String status,
        String sensor_id,
        String sensor_type,
        String description
) {
    public static @NonNull TemperatureData getRandom(@NonNull String location, @NonNull String sensorId) {
        double value = ThreadLocalRandom.current().nextInt(180, 401) * .1;
        return new TemperatureData(value, "°C", OffsetDateTime.now(), location, "active", sensorId, "temperature", "Temperature sensor in " + location);
    }
}
