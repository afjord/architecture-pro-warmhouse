package org.example;

import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TemperatureController {

    @GetMapping("health")
    public void health() {
    }

    @GetMapping("temperature")
    public TemperatureData getTemperature(@RequestParam(required = false) String location,
                                          @RequestParam(value = "sensorID", required = false) String sensorId) {
        if (!StringUtils.hasText(location) && !StringUtils.hasText(sensorId)) {
            throw new IllegalArgumentException();
        }
        if (!StringUtils.hasText(location)) {
            location = switch (sensorId) {
                case "1" -> "Living Room";
                case "2" -> "Bedroom";
                case "3" -> "Kitchen";
                default -> "Unknown";
            };
        }
        if (!StringUtils.hasText(sensorId)) {
            sensorId = switch (location) {
                case "Living Room" -> "1";
                case "Bedroom" -> "2";
                case "Kitchen" -> "3";
                default -> "0";
            };
        }
        return TemperatureData.getRandom(location, sensorId);
    }

    @GetMapping("temperature/{id}")
    public TemperatureData getTemperatureBySensorId(@PathVariable String id) {
        String location = switch (id) {
            case "1" -> "Living Room";
            case "2" -> "Bedroom";
            case "3" -> "Kitchen";
            default -> "Unknown";
        };
        return TemperatureData.getRandom(location, id);
    }
}
