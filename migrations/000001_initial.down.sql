-- Rollback initial migration
-- Migration: 000001_initial

-- Drop tables in reverse order of creation (due to foreign key constraints)
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS waitlist_entries;
DROP TABLE IF EXISTS promotions;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS newsletter_subscribers;
DROP TABLE IF EXISTS ignored_posts;
DROP TABLE IF EXISTS event_interests;
DROP TABLE IF EXISTS event_submissions;
DROP TABLE IF EXISTS event_dates;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS clubs;
