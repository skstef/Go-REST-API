# Event Booking REST API

## Introduction

This project is a Go-powered **Event Booking REST API** that allows users to manage events by performing CRUD (Create, Read, Update, Delete) operations. The API also includes secure authentication and authorization, ensuring that only authenticated users can create, update, or delete events. It is designed for scalability and efficiency, leveraging Go's concurrency features for optimal performance.

## Available Endpoints List

Example requests for each endpoint can be found in the `api-test` folder.

### Public endpoints:

- **GET** `/events` – Retrieves a list of all events.
- **GET** `/events/{id}` – Fetches details of a specific event by its ID.
- **POST** `/signup` – Registers a new user.
- **POST** `/login` – Authenticates a user and provides a token.

### Protected Endpoints (Authorization Required)

- **POST** `/events` – Creates a new event.
- **PUT** `/events/{id}` – Updates an existing event by its ID.
- **DELETE** `/events/{id}` – Deletes a specific event by its ID.
- **POST** `/events/{id}/register` – Registers a user for an event.
- **POST** `/events/{id}/cancel` – Cancels a user's registration for an event.

## How to Run the Project

1. Clone the repository:

```bash
git clone https://github.com/skstef/Go-REST-API.git
cd event-booking-api
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the project: You can run the server with the following command:

```bash
go run .
```
