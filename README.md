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

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/movietix.git
    cd movietix
    ```

2. **Set up PostgreSQL:**

    Make sure you have PostgreSQL installed and running. Create a database for the project:

    ```sql
    CREATE DATABASE movietix;
    ```

3. **Configure environment variables:**

    Create a `.env` file in the project root with the following contents. Modify as per your configuration:

    ```plaintext
    DB_HOST=hostname
    DB_PORT=port
    DB_USERNAME=username
    DB_PASSWORD=password
    DB_NAME=database
    ```

4. **Run database migrations:**

    Use the following commands to setup your database schema and triggers:

    ```bash
    psql -h $DB_HOST -p $DB_PORT -U $DB_USERNAME -d $DB_NAME -f schema/ddl.sql
    psql -h $DB_HOST -p $DB_PORT -U $DB_USERNAME -d $DB_NAME -f schema/updated_at_trigger.sql
    ```

5. **Install dependencies:**

    ```bash
    go mod tidy
    ```

6. **Run the application:**

    ```bash
    go run cmd/main.go
    ```

