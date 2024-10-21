# MovieTix

MovieTix is a movie ticket booking system built in Go. It allows users to book tickets for movies in various formats and audio configurations at different cinemas.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Database Schema](#database-schema)
- [Installation](#installation)

## Features

- User registration and authentication
- Movie listings with details
- Cinema and auditorium management
- Show scheduling
- Seat booking with payment processing

## Technologies

- **Go**: The main programming language.
- **PostgreSQL**: The database used for storing all the data.
- **pgx**: A PostgreSQL driver and toolkit for Go.
- **Gorilla Mux**: A powerful URL router and dispatcher for Golang.
- **sqlc**: A tool for generating type-safe Go code from SQL queries.

## Database Schema

The database schema includes the following tables:

- **users**: Stores user information.
- **movies**: Stores movie details.
- **cinemas**: Stores cinema details.
- **auditoriums**: Stores details of auditoriums within cinemas.
- **sections**: Stores information about different seating sections in auditoriums.
- **seats**: Stores details about individual seats.
- **shows**: Stores information about movie showtimes.
- **payments**: Stores payment details.
- **bookings**: Stores booking details.

## Installation

### 1. Clone the repository:

  ```bash
  git clone https://github.com/sritejachilakapati/movietix.git
  cd movietix
  ```

### 2. Set up PostgreSQL:

  Make sure you have PostgreSQL installed and running. Create a database for the project:

  ```sql
  CREATE DATABASE movietix;
  ```

### 3. Configure environment variables:

  Create a `.env` file in the project root with the following contents. Modify as per your configuration:

  ```plaintext
  DB_HOST=hostname
  DB_PORT=port
  DB_USERNAME=username
  DB_PASSWORD=password
  DB_NAME=database
  ```

### 4. Run database migrations:

  Install the `golang-migrate` package using one of the following methods:

  #### Homebrew (macOS):

  ```bash
  brew install golang-migrate
  ```

  #### Linux (apt/yum):

  - **For apt (Debian/Ubuntu):**

    ```bash
    sudo apt-get update
    sudo apt-get install -y golang-migrate
    ```

  - **For yum (CentOS/Fedora):**

    ```bash
    sudo yum install golang-migrate
    ```

  #### Go install (Go toolchain):

  ```bash
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```

  #### Download from GitHub:

  1. Download the binary for your OS from the [golang-migrate releases](https://github.com/golang-migrate/migrate/releases).
  2. Extract the binary and move it to your `PATH` (e.g., `/usr/local/bin`).

  #### Run the migrations:

  ```bash
  migrate -path ./migrations -database $DB_URL up
  ```

### 5. Install dependencies:

  ```bash
  go mod tidy
  ```

### 6. Run the application:

  ```bash
  go run cmd/main.go
  ```
