-- =====================
-- USERS
-- Stores application users (customers, admins, theater managers)
-- =====================
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT UNIQUE, -- Nullable for OAuth-only accounts
    phone TEXT UNIQUE, -- Optional contact number
    password_hash TEXT, -- Nullable for OAuth-only
    oauth_provider TEXT, -- e.g. 'google', 'facebook', 'apple'
    oauth_provider_id TEXT, -- Provider's unique user ID
    role TEXT NOT NULL DEFAULT 'customer', -- customer, admin, theater_manager
    status TEXT NOT NULL DEFAULT 'active', -- active, disabled, banned
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_status ON users(status);

-- =====================
-- THEATERS
-- Stores information about theater locations
-- =====================
CREATE TABLE theaters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    address TEXT NOT NULL, -- Human-readable address
    location GEOGRAPHY(Point, 4326) NOT NULL, -- For distance-based queries
    contact_phone TEXT,
    contact_email TEXT,
    status TEXT NOT NULL DEFAULT 'active', -- active, closed, under_renovation
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_theaters_status ON theaters(status);
CREATE INDEX idx_theaters_location ON theaters USING GIST(location);

-- =====================
-- THEATER_MANAGERS
-- Maps managers to theaters they control
-- =====================
CREATE TABLE theater_managers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    theater_id UUID NOT NULL REFERENCES theaters(id) ON DELETE CASCADE,
    UNIQUE (user_id, theater_id)
);
CREATE INDEX idx_theater_managers_user_id ON theater_managers(user_id);
CREATE INDEX idx_theater_managers_theater_id ON theater_managers(theater_id);

-- =====================
-- SCREENS
-- Represents individual screens inside theaters
-- =====================
CREATE TABLE screens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    theater_id UUID NOT NULL REFERENCES theaters(id) ON DELETE CASCADE,
    name TEXT NOT NULL, -- e.g., "Screen 1"
    total_seats INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_screens_theater_id ON screens(theater_id);

-- =====================
-- MOVIES
-- Stores localized movie data (language-specific versions)
-- =====================
CREATE TABLE movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL, -- e.g., "Oppenheimer"
    language_code TEXT NOT NULL, -- ISO code, e.g., 'en', 'hi'
    synopsis TEXT,
    release_date DATE NOT NULL,
    runtime_minutes INT,
    poster_url TEXT,
    trailer_url TEXT,
    rating TEXT, -- e.g., PG-13, R, U/A
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_movies_title ON movies(title);
CREATE INDEX idx_movies_language_code ON movies(language_code);

-- =====================
-- FORMATS
-- Available screening formats (e.g., IMAX, Dolby Atmos)
-- =====================
CREATE TABLE formats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT
);
CREATE UNIQUE INDEX idx_formats_name ON formats(name);

-- =====================
-- SCREEN_FORMAT_CAPABILITIES
-- Which formats a screen supports
-- =====================
CREATE TABLE screen_format_capabilities (
    screen_id UUID NOT NULL REFERENCES screens(id) ON DELETE CASCADE,
    format_id UUID NOT NULL REFERENCES formats(id) ON DELETE CASCADE,
    PRIMARY KEY (screen_id, format_id)
);
CREATE INDEX idx_sfc_format_id ON screen_format_capabilities(format_id);

-- =====================
-- SHOWS
-- Scheduled movie showings
-- =====================
CREATE TABLE shows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    movie_id UUID NOT NULL REFERENCES movies(id),
    screen_id UUID NOT NULL REFERENCES screens(id),
    start_time TIMESTAMPTZ NOT NULL,
    status TEXT NOT NULL DEFAULT 'scheduled', -- scheduled, running, finished, cancelled
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_shows_movie_id ON shows(movie_id);
CREATE INDEX idx_shows_screen_id ON shows(screen_id);
CREATE INDEX idx_shows_start_time ON shows(start_time);
CREATE INDEX idx_shows_status ON shows(status);

-- =====================
-- SEAT_TYPES
-- Defines categories of seats
-- =====================
CREATE TABLE seat_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL, -- Regular, Premium, Recliner
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX idx_seat_types_name ON seat_types(name);

-- =====================
-- SEATS
-- Stores physical seats for a screen
-- =====================
CREATE TABLE seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    screen_id UUID NOT NULL REFERENCES screens(id) ON DELETE CASCADE,
    seat_row TEXT NOT NULL, -- e.g., 'A'
    seat_number INT NOT NULL, -- e.g., 1, 2, 3
    seat_type_id UUID NOT NULL REFERENCES seat_types(id),
    status TEXT NOT NULL DEFAULT 'active', -- active, inactive
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (screen_id, seat_row, seat_number)
);
CREATE INDEX idx_seats_screen_id ON seats(screen_id);
CREATE INDEX idx_seats_seat_type_id ON seats(seat_type_id);
CREATE INDEX idx_seats_status ON seats(status);

-- =====================
-- SHOW_SEAT_PRICING
-- Price per seat type for a specific show
-- =====================
CREATE TABLE show_seat_pricing (
    show_id UUID NOT NULL REFERENCES shows(id) ON DELETE CASCADE,
    seat_type_id UUID NOT NULL REFERENCES seat_types(id),
    price NUMERIC(10,2) NOT NULL,
    PRIMARY KEY (show_id, seat_type_id)
);
CREATE INDEX idx_ssp_seat_type_id ON show_seat_pricing(seat_type_id);

-- =====================
-- BOOKINGS
-- Represents a user's ticket purchase
-- =====================
CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    show_id UUID NOT NULL REFERENCES shows(id),
    total_amount NUMERIC(10,2) NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending', -- pending, confirmed, cancelled
    payment_reference TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_show_id ON bookings(show_id);
CREATE INDEX idx_bookings_status ON bookings(status);

-- =====================
-- BOOKING_ITEMS
-- Individual seat bookings within a booking
-- =====================
CREATE TABLE booking_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    seat_id UUID NOT NULL REFERENCES seats(id),
    price NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (booking_id, seat_id)
);
CREATE INDEX idx_booking_items_booking_id ON booking_items(booking_id);
CREATE INDEX idx_booking_items_seat_id ON booking_items(seat_id);

-- =====================
-- SEAT_ALLOCATIONS
-- Permanent record of seat usage per show
-- =====================
CREATE TABLE seat_allocations (
    show_id UUID NOT NULL REFERENCES shows(id),
    seat_id UUID NOT NULL REFERENCES seats(id),
    booking_id UUID, -- NULL until confirmed
    is_booked BOOLEAN DEFAULT false NOT NULL,
    PRIMARY KEY (show_id, seat_id)
);
CREATE INDEX idx_seat_allocations_booking_id ON seat_allocations(booking_id);
CREATE INDEX idx_seat_allocations_is_booked ON seat_allocations(is_booked);
