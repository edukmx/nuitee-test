# Nuitee Test

# Installation

1. Edit .env
```
HOTELBEDS_HOST=https://api.test.hotelbeds.com
HOTELBEDS_API_KEY=3a2a7512aca242d176e7f78705d0d890
HOTELBEDS_SECRET=24d8734c7c
SERVER_PORT=:8080
TIMEOUT=4
```

2. Build and Run Docker
```
docker-compose build
docker-compose up -d
```

# Usage

- Get Cheapest Rates Endpoint
```
curl --location -g --request GET 'localhost:8080/hotels/?checkin=2025-06-15&checkout=2025-06-16&currency=EUR&guestNationality=US&hotelIds=77,168,264,265,29&occupancies=[{"rooms":1, "adults": 2},{"rooms":2, "adults": 2}]' \
--header 'x-liteapi-supplier-config: {{liteapi_config}
```

# Run Unit Test

```
cd src && make test
```