
# Vehicle Lookup and PubSub Integration

This project implements a vehicle lookup system using an API, followed by a mocked Pub/Sub message publish if certain criteria are met. The project performs vehicle searches and publishes results to a mock Pub/Sub system when a valid "hire vehicle" is identified.

---

## Project Overview

This application simulates searching for a vehicle via the Transfer360 Sandbox API and publishes the search results to a Pub/Sub service if the vehicle meets certain conditions (e.g., the vehicle is a hire vehicle). It uses Go’s concurrency features to perform searches for multiple companies concurrently. Mock implementations for Pub/Sub are used to simulate message publishing during development.

## Technologies Used
- **Go 1.24.1+**: Main programming language.
- **Pub/Sub**: Simulated using a `MockPublisher` (in the absence of an actual Google Cloud Pub/Sub or similar service).
- **JSON**: For request/response data formatting.

## Installation

To run the project, you need to install Go and clone the repository.

### Prerequisites

- Go 1.24+ (for compatibility with Go Modules)
- Git for cloning the repository

### Steps

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/jamesoneill997/t360.git
   cd t360
   ```

2. **Install Dependencies:**
   Go modules will handle all dependencies. Run the following command to install them:
   ```bash
   go mod download
   ```

---

## Running the Application

### To Run the Application:

The application can be run via CLI with arguments for the vehicle VRM and contravention date.

1. **Run with Flags**:
   ```bash
   go run main.go --vrm GV01TPH --contravention_date 2025-03-18T20:16:39.387Z
   ```

   The flags `--vrm` and `--contravention_date` are required.

---
