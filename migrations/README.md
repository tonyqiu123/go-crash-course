# Database Migrations

This directory contains database migration files for the Go backend.

## Migration Tool

We recommend using [golang-migrate](https://github.com/golang-migrate/migrate) for running migrations.

### Installation

```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Go install
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Running Migrations

```bash
# Up migration (apply)
migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# Down migration (rollback)
migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down

# Specific version
migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" goto 1
```

### Creating New Migrations

```bash
migrate create -ext sql -dir migrations -seq migration_name
```

This will create two files:
- `XXXXXX_migration_name.up.sql` - Forward migration
- `XXXXXX_migration_name.down.sql` - Rollback migration

## Migration Files

- `000001_initial.up.sql` - Initial database schema with all tables
- `000001_initial.down.sql` - Rollback for initial schema
