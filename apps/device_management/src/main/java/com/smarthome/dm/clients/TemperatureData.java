package com.smarthome.dm.clients;

import lombok.Getter;
import lombok.Setter;

import java.time.OffsetDateTime;

@Setter
@Getter
public class TemperatureData {
    private String location;
    private Double value;
    private String unit;
    private String status;
    private OffsetDateTime timestamp;
    private String description;
}
