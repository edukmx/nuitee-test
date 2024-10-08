# Nuitee Test

Nuitee Test is an API designed to fetch hotel rates, including the cheapest available rates from various hotel IDs. It interacts with the HotelBeds API to provide real-time data for hotel pricing based on guest information, nationality, and other parameters.

# Installation

## Prerequisites
Ensure you have the following installed on your system before proceeding:

- Docker
- Docker Compose
- Go (if you plan to build or develop locally)

1. Clone the repo
```
git clone https://github.com/edukmx/nuitee-test.git
cd nuitee-test
```

2. Edit .env
```
HOTELBEDS_HOST=https://api.test.hotelbeds.com
HOTELBEDS_API_KEY=3a2a7512aca242d176e7f78705d0d890
HOTELBEDS_SECRET=24d8734c7c
SERVER_PORT=:8080
TIMEOUT=4
```

3. Build and Run Docker
```
docker-compose build
docker-compose up -d
```

4. Accessing the API

Once the containers are running, the application should be accessible at http://localhost:8080.

# Usage

- Get Cheapest Rates Endpoint

You can use the following curl command to test the Get Cheapest Rates endpoint:
```
curl --location -g --request GET 'localhost:8080/hotels/?checkin=2025-06-15&checkout=2025-06-16&currency=EUR&guestNationality=US&hotelIds=77,168,264,265,29&occupancies=[{"rooms":1, "adults": 2},{"rooms":2, "adults": 2}]' \
```

This request fetches the cheapest available rates for the specified hotels, dates, and occupancy details.

Parameters:
- checkin: Check-in date in the format YYYY-MM-DD.
- checkout: Check-out date in the format YYYY-MM-DD.
- currency: The currency code in which to return rates (e.g., EUR for Euros).
- guestNationality: The ISO code of the guest's nationality (e.g., US for United States).
- hotelIds: A comma-separated list of hotel IDs to fetch rates for.
- occupancies: A JSON array specifying the number of rooms and guests.

# Run Unit Test

Unit tests ensure the correctness of the core logic. You can run the tests using the following commands:

Change to the source directory:
```
cd src
```

Run the tests using make:

```
make test
```


## Project Structure
The project follows a clean and modular structure to separate concerns. Here is a quick overview of the key directories:

- src/cmd/api/main.go: This is the entry point of the API where the server starts.
- src/internal: The internal logic of the project is stored here. It includes domain services, repository interactions, and handlers.
- internal/app: Application layer, which includes use cases and business logic.
- internal/domain: Domain layer, responsible for the core entities and domain logic.
- internal/infra: Infrastructure layer, dealing with external systems (APIs, databases, etc.).


This command will execute all the unit tests and output the results.

## Contributing
Contributions to this project are welcome. If you'd like to contribute, please follow these steps:

- Fork the repository.
- Create a feature branch (git checkout -b feature/new-feature).
- Commit your changes (git commit -am 'Add new feature').
- Push the branch (git push origin feature/new-feature).
- Open a Pull Request.

## Notes
- API Rate Limiting: If you are using the test API for HotelBeds, be mindful of API rate limits.
- Timeout: The TIMEOUT environment variable specifies the request timeout in seconds for API calls. You can adjust it based on your use case.
