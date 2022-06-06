### WEATHER

Access current weather data for any location.

```
GET /api/v1/weather?lat=35&lon=139
```

- **lat, lon** - Geographical coordinates (latitude, longitude).

```
GET /api/v1/weather?q=London
```

- **q** - City name, state code and country code.

```
GET /api/v1/weather?id=2172797
```

- **id** - City ID.

```
GET /api/v1/weather?zip=634000,ru
```

- **zip** - Zip code.

### FORECAST

5 day forecast is available at any location on the globe. It includes weather forecast data with 3-hour step. Default cnt = 40.

```
GET /api/v1/forecast?lat=35&lon=139&cnt=3
```

- **lat, lon** - Geographical coordinates (latitude, longitude).
- **cnt [optional]** - A number of timestamps.

```
GET /api/v1/forecast?q=London&cnt=3
```

- **q** - City name, state code and country code.
- **cnt [optional]** - A number of timestamps.

```
GET /api/v1/forecast?id=2172797&cnt=3
```

- **id** - City ID.
- **cnt [optional]** - A number of timestamps.

```
GET /api/v1/forecast?zip=634000,ru&cnt=3
```

- **zip** - Zip code.
- **cnt [optional]** - A number of timestamps.