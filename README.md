# Gator

A CLI RSS feed aggregator built in Go and backed by PostgreSQL.

## Prerequisites

Ensure you have the following installed on your machine:
* [Go](https://go.dev/) (v1.20+)
* [PostgreSQL](https://www.postgresql.org/)

## Installation

Install the `gator` CLI directly using `go install`:

```bash
go install github.com/samiiac/bootdev-gator@latest
```

## Setup

1. Make sure PostgreSQL is running.
2. Create `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Commands

* **`gator register <username>`** – Creates a new user account and sets them as current.
* **`gator login <username>`** – Switches the active user logged into the CLI.
* **`gator addfeed <name> <url>`** – Adds a new feed and follows it automatically.
* **`gator follow <url>`** – Follows feed using its RSS URL.
* **`gator agg <time_duration>`** – Runs the continuous feed scraper loop.
* **`gator browse [limit]`** – Displays posts from followed feeds up to an optional limit.