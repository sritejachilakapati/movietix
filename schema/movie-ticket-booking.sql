CREATE TYPE user_role AS ENUM ('admin', 'user');
CREATE TYPE video_format AS ENUM ('2D', '3D', 'IMAX 2D', 'IMAX 3D', '4DX', 'ICE');
CREATE TYPE audio_format AS ENUM ('5.1', '7.1', 'ATMOS');
CREATE TYPE show_status AS ENUM ('OPEN', 'CLOSED', 'CANCELLED');
CREATE TYPE availability_status AS ENUM ('AVAILABLE', 'FAST FILLING', 'SOLD OUT');
CREATE TYPE booking_status AS ENUM ('PENDING', 'SUCCESS', 'FAILED', 'CANCELLED');
CREATE TYPE payment_status AS ENUM ('created', 'paid', 'failed', 'refunded');

CREATE TABLE users (
  id uuid PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  role user_role NOT NULL DEFAULT 'user',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE movies (
  id uuid PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  audio_language CHAR(2) NOT NULL,
  subtitle_language CHAR(2) NOT NULL,
  formats VARCHAR(10)[] NOT NULL, -- array of `video_format` enum
  poster VARCHAR(255) NOT NULL,
  release_date DATE NOT NULL,
  synopsis TEXT NOT NULL,
  trailer VARCHAR(255) NOT NULL,
  cast VARCHAR(255)[] NOT NULL,
  runtime_minutes INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE cinemas (
  id uuid PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  owner_id uuid REFERENCES users(id),
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  location GEOMETRY(Point, 4326) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE auditoriums (
  id uuid PRIMARY KEY,
  cinema_id uuid REFERENCES cinemas(id),
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  subtitles BOOLEAN NOT NULL DEFAULT FALSE,
  video_format video_format NOT NULL,
  audio_format audio_format NOT NULL,
  total_rows INT NOT NULL,
  total_seats_per_row INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE sections (
  id uuid PRIMARY KEY,
  auditorium_id uuid REFERENCES auditoriums(id),
  name VARCHAR(255) NOT NULL, -- 'Prime' | 'Gold' | 'Recliner' etc
  price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE seats (
  id uuid PRIMARY KEY,
  auditorium_id uuid REFERENCES auditoriums(id),
  section_id uuid REFERENCES sections(id),
  row_key VARCHAR(2) NOT NULL,  -- 'A' | 'B' | 'C' etc till the total rows
  seat_number INT NOT NULL,
  order INT NOT NULL, -- seat order always starts from the left side of the first row when facing the screen. Order will go from `1` till the capacity of the auditorium.
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE (auditorium_id, row_key, seat_number)
);

CREATE TABLE shows (
  id uuid PRIMARY KEY,
  auditorium_id uuid REFERENCES auditoriums(id),
  movie_id uuid REFERENCES movies(id),
  show_time TIMESTAMP NOT NULL,
  status show_status NOT NULL,
  availability availability_status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE payments (
  id uuid PRIMARY KEY,
  order_id varchar(60) NOT NULL, -- third party payment gateway order id
  amount_rupees DECIMAL(10, 2) NOT NULL,
  status payment_status NOT NULL,
  payment_time TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE bookings (
  id uuid PRIMARY KEY,
  user_id uuid REFERENCES users(id),
  show_id uuid REFERENCES shows(id),
  seat_id uuid REFERENCES seats(id),
  status booking_status NOT NULL,
  payment_id uuid REFERENCES payments(id),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);