# GeO-Locator

GeO-Locator is a Go-based application designed to provide geolocation services by leveraging configurable providers. The application offers a simple API to retrieve location data.

## Features
- **Configurable Geolocation Providers:** Easily switch between different geolocation providers.
- **Robust Error Handling:** Custom error management for improved debugging and error response.
- **Modular Structure:** Organized into packages for configuration, API handling, and provider management.

## Project Structure
- **cmd/geoLocator**: Main application entry point.
- **internal/config**: Manages configuration constants and loading.
- **internal/customErrors**: Contains custom error types for better error handling.
- **api**: Handles API routes, middleware, and request handling.
- **pkg/geoLocationProvider**: Provides an interface for geolocation services, with a dummy provider for testing.

## Prerequisites
- Go 1.22 or later
- Docker (for containerized development environment)
