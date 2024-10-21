DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP TRIGGER IF EXISTS update_movies_updated_at ON movies;
DROP TRIGGER IF EXISTS update_cinemas_updated_at ON cinemas;
DROP TRIGGER IF EXISTS update_auditoriums_updated_at ON auditoriums;
DROP TRIGGER IF EXISTS update_sections_updated_at ON sections;
DROP TRIGGER IF EXISTS update_seats_updated_at ON seats;
DROP TRIGGER IF EXISTS update_shows_updated_at ON shows;
DROP TRIGGER IF EXISTS update_payments_updated_at ON payments;
DROP TRIGGER IF EXISTS update_bookings_updated_at ON bookings;

DROP FUNCTION IF EXISTS update_updated_at_column();
