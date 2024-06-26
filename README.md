# Fetch Backend Challenge

This project implements the Fetch Backend Challenge using Go to create REST APIs.

## About

This web service processes receipts and calculates points based on the criteria provided in the challenge.

### Endpoints

1. **Process Receipts**
   - **Path:** `/receipts/process`
   - **Method:** `POST`
   - **Payload:** JSON receipt
   - **Response:** JSON containing an ID for the receipt.

2. **Get Points**
   - **Path:** `/receipts/{id}/points`
   - **Method:** `GET`
   - **Response:** JSON object with the number of points awarded.

### Features

- Comprehensive unit and integration tests
- Validation for every field in the input payload
- Thorough error handling

## Implementation

When a POST request with a receipt object is received:
1. Every field is validated using regex validators.
2. A UUID is generated and checked within the map of UUID: points to ensure no collisions.
3. Points are calculated and stored in the map with their respective UUID.
4. The UUID is sent back as a response.

For a GET request for points with a specific ID:
- The map is checked for the ID, and the points are returned.

### Security Measures

1. No payload is accepted for GET requests.
2. Only the defined API endpoints are allowed; all other methods return "Method Not Allowed."
3. Unidentified paths return a 404 error.

## Repository Structure

All implementations exist within the main package:

- `routes.go`: Contains the API routes.
- `handlers.go`: Contains the request handlers.
- `models.go`: Defines the data models.
- `points_calculator.go`: Core logic for points calculation.
- `utils.go`: Helper functions, including UUID generator and validators.
- `unit_test.go`: Unit tests for functions.
- `integration_test.go`: Integration tests to verify flow and point calculations.

## Prerequisites

- Go installed (version 1.20 used in this implementation).

## Starting the web server

   1.  Clone this repo 
   ```sh
   git clone https://github.com/swethapaturu9/fetch-backend-challenge.git
   ```
   2. Navigate to the root directory and then main package
   ```sh
   cd fetch-backend-challenge/main
   ```
   3. Run the command
   ```sh
   go run . 
   ```
   4. Run unit & integration tests
  ```sh
  go test -v   
   ```
