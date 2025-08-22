# Booking App (Go)

A simple command‑line application written in Go that simulates booking tickets for an event. It walks a user through entering their name, email, and number of tickets, validates input, records the booking, and prints friendly summaries.

> This repo is a learning project to practice Go basics (packages, variables/constants, loops, slices, maps, functions, and input validation).

## Features

* Prompt‑driven CLI flow for booking tickets
* Basic validation (name length, email contains `@`, remaining tickets bounds)
* In‑memory storage of bookings (slice of maps)
* Helpful summaries (first names list, tickets remaining)

## Tech Stack

* **Language:** Go (Golang)

## Project Structure

```
.
├── go.mod        # Go module definition
├── main.go       # Entry point and CLI flow
└── helper.go     # Validation and helper utilities
```

## Getting Started

### Prerequisites

* Go 1.20+ installed (check with `go version`)

### Clone the repository

```bash
git clone https://github.com/Raghav847/booking-app.git
cd booking-app
```

### Run

```bash
# Run directly
go run .

# Or build and run
go build -o booking
./booking
```

## Usage

When you run the program, you’ll be prompted for:

1. First name
2. Last name
3. Email
4. Number of tickets to book

If the inputs are valid and tickets are available, the booking is stored and you’ll see an updated list of first names and remaining tickets.

### Example

```
Welcome to our Go Conference booking application!
We have total of 50 tickets and 50 are still available.
Get your tickets here to attend.

Enter your first name: Raghav
Enter your last name: Khandelwal
Enter your email address: raghav@example.com
Enter number of tickets: 2

Thank you Raghav Khandelwal for booking 2 tickets. You will receive a confirmation email at raghav@example.com
first names: [Raghav]
```

## How It Works

* **Validation:** `helper.go` contains functions that check for a minimum name length, a basic email format, and whether the requested ticket quantity fits within the remaining pool.
* **State:** Bookings are held in a slice of maps (e.g., `[]map[string]string`), and remaining tickets are tracked with a `uint` counter.
* **Loop:** The main loop continues to accept bookings until tickets run out or the app is exited.

## Common Tasks

* **Run tests** *(if/when tests are added)*:

  ```bash
  go test ./...
  ```
* **Format & vet**:

  ```bash
  go fmt ./...
  go vet ./...
  ```

## Roadmap / Ideas

* Persist bookings to a file or database
* Stronger email validation and input parsing
* Graceful cancellation / editing of bookings
* Unit tests covering validation and booking logic
* Command‑line flags (e.g., total tickets, event name)

## Contributing

Issues and pull requests are welcome! If you’re practicing Go too, feel free to fork and experiment.

## License

This project currently does not include an explicit license. If you plan to use or distribute it, consider adding a license (e.g., MIT).
