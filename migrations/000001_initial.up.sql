-- Initial migration: Create all tables for the Go backend
-- Migration: 000001_initial

-- Clubs table
CREATE TABLE IF NOT EXISTS clubs (
    id SERIAL PRIMARY KEY,
    club_name VARCHAR(100) NOT NULL UNIQUE,
    categories JSONB DEFAULT '[]',
    club_page TEXT,
    ig TEXT,
    discord TEXT,
    club_type VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_clubs_club_name ON clubs(club_name);
CREATE INDEX IF NOT EXISTS idx_clubs_deleted_at ON clubs(deleted_at);

-- Events table
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    location TEXT,
    categories JSONB DEFAULT '[]',
    status VARCHAR(32),
    source_url TEXT,
    source_image_url TEXT,
    reactions JSONB DEFAULT '{}',
    posted_at TIMESTAMP WITH TIME ZONE,
    comments_count INTEGER DEFAULT 0,
    likes_count INTEGER DEFAULT 0,
    food VARCHAR(255),
    registration BOOLEAN DEFAULT false,
    added_at TIMESTAMP WITH TIME ZONE,
    price DOUBLE PRECISION,
    school VARCHAR(255),
    club_type VARCHAR(50),
    ig_handle VARCHAR(100),
    discord_handle VARCHAR(100),
    x_handle VARCHAR(100),
    tiktok_handle VARCHAR(100),
    fb_handle VARCHAR(100),
    other_handle VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_events_deleted_at ON events(deleted_at);

-- Event dates table
CREATE TABLE IF NOT EXISTS event_dates (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    dtstart_utc TIMESTAMP WITH TIME ZONE NOT NULL,
    dtend_utc TIMESTAMP WITH TIME ZONE,
    duration BIGINT, -- Duration in nanoseconds
    tz VARCHAR(64),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_eventdates_dtstart_utc ON event_dates(dtstart_utc);
CREATE INDEX IF NOT EXISTS idx_eventdates_dtend_utc ON event_dates(dtend_utc);
CREATE INDEX IF NOT EXISTS idx_eventdates_event_dtstart ON event_dates(event_id, dtstart_utc);
CREATE INDEX IF NOT EXISTS idx_event_dates_deleted_at ON event_dates(deleted_at);

-- Event submissions table
CREATE TABLE IF NOT EXISTS event_submissions (
    id SERIAL PRIMARY KEY,
    submitted_by VARCHAR(255) NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    reviewed_at TIMESTAMP WITH TIME ZONE,
    reviewed_by VARCHAR(255),
    created_event_id INTEGER NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_event_submissions_submitted_at ON event_submissions(submitted_at);
CREATE INDEX IF NOT EXISTS idx_event_submissions_deleted_at ON event_submissions(deleted_at);

-- Event interests table
CREATE TABLE IF NOT EXISTS event_interests (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_event_interests_event_id ON event_interests(event_id);
CREATE INDEX IF NOT EXISTS idx_event_interests_user_id ON event_interests(user_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_event_interests_event_user ON event_interests(event_id, user_id);
CREATE INDEX IF NOT EXISTS idx_event_interests_deleted_at ON event_interests(deleted_at);

-- Ignored posts table
CREATE TABLE IF NOT EXISTS ignored_posts (
    id SERIAL PRIMARY KEY,
    shortcode VARCHAR(32) NOT NULL UNIQUE,
    added_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_ignored_posts_deleted_at ON ignored_posts(deleted_at);

-- Newsletter subscribers table
CREATE TABLE IF NOT EXISTS newsletter_subscribers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_newsletter_subscribers_deleted_at ON newsletter_subscribers(deleted_at);

-- Payments table
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    currency VARCHAR(3) DEFAULT 'CAD',
    status VARCHAR(32) NOT NULL,
    payment_method VARCHAR(64),
    transaction_id VARCHAR(255) UNIQUE,
    stripe_session_id VARCHAR(255),
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_payments_user_id ON payments(user_id);
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at ON payments(deleted_at);

-- Promotions table
CREATE TABLE IF NOT EXISTS promotions (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url TEXT,
    link_url TEXT,
    active BOOLEAN DEFAULT true,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    priority INTEGER DEFAULT 0,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_promotions_deleted_at ON promotions(deleted_at);

-- Waitlist entries table
CREATE TABLE IF NOT EXISTS waitlist_entries (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255),
    school VARCHAR(255),
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_waitlist_entries_deleted_at ON waitlist_entries(deleted_at);

-- Users table (for Clerk integration)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    clerk_id VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255),
    name VARCHAR(255),
    role VARCHAR(32) DEFAULT 'user',
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);
