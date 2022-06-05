# Nord - caching proxy server for [OpenWeather](https://openweathermap.org/)

Простой веб-сервер позволяющий кэшировать ответы от сервиса **OpenWeather**, чтобы не превышать лимит вызовов API. OpenWeather обновляет данные раз в 10 минут, но хранить можно опционально задав параметр RedisExpiration.

- data 2.5
- geo 1.0

**Caching proxy server:**

- [Current weather data](https://openweathermap.org/current) and [5 Day / 3 Hour Forecast](https://openweathermap.org/forecast5).
  ✅ by geographic coordinates
  ✅ by city name
  ✅ by city ID
  ✅ by ZIP code

  [swagger](swagger.md)

**OpenWeather client:**

- [Current weather data](https://openweathermap.org/current)
  ✅ by geographic coordinates
  ✅ by city name
  ✅ by city ID
  ✅ by ZIP code

  [example](./pkg/openweather/examples/currentWeather/main.go)

- [5 Day / 3 Hour Forecast](https://openweathermap.org/forecast5)
  ✅ by geographic coordinates
  ✅ by city name
  ✅ by city ID
  ✅ by ZIP code

  [example](./pkg/openweather/examples/forecast/main.go)

- [Geocoding](https://openweathermap.org/api/geocoding-api)
  ✅ Coordinates by location name
  ✅ Coordinates by zip/post code
  ✅ Reverse geocoding

  [example](./pkg/openweather/examples/geocoding/main.go)

### Используемые технологии:

- Go 1.17.8
- Redis 6.2.6
- Nginx 1.21.6
- Docker-compose 1.29.2

### Запуск сервера:

Установите [**api key**](./cmd/cache-server/main.go). Получить его можно [**здесь**](https://home.openweathermap.org/api_keys). Ключи в данном репозитории недействительны и оставлены, как пример.

- `make run` выполнит docker-compose up --detach для для данного проекта в dev моде.

  После чего можно выполнить запрос:

  ```
  GET http://<your ip address>/api/v1/ping
  ```

  И получить ответ:

  ```json
  HTTP/1.1 200 OK
  content-type: text/plain;charset=UTF-8

  pong
  ```

  [swagger](swagger.md)

- `make stop` выполнит docker-compose down.

- `make clear` удалит docker images подходящие под маску dev\*:latest.

  **Windows 10:** `make clear shell=powershell`
