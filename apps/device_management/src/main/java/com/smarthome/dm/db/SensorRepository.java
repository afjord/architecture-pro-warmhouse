package com.smarthome.dm.db;

import org.springframework.data.jpa.repository.JpaRepository;

public interface SensorRepository extends JpaRepository<Device, Integer> {
}