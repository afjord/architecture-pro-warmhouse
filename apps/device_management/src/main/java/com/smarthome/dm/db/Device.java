package com.smarthome.dm.db;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.time.OffsetDateTime;
import java.util.UUID;

@Getter
@Setter
@Entity
@Table(name = "devices")
@EntityListeners(DeviceEntityListener.class)
public class Device {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    private UUID homeId;
    private String name;
    private String type;
    private String location;
    private String unit;
    private Double value;
    private String status;
    private OffsetDateTime lastUpdated;
    private OffsetDateTime createdAt;
}