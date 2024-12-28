# TikTok Playwright Project Documentation

## Overview

The TikTok Playwright project is a web application that allows users to search for TikTok videos and play them. It utilizes the Playwright library to automate browser interactions and fetch video data. The application is built using the Go programming language and the Gin web framework.

## Project Structure

```
tiktok-playwright/
├── go.mod
├── go.sum
├── main.go
├── handlers/
│   ├── play_handler.go
│   └── search_handler.go
├── services/
│   └── video_service.go
├── utils/
│   ├── redis.go
│   └── response.go
└── README.md
```

## Dependencies

The project uses the following dependencies, as specified in `go.mod`:

- `github.com/gin-gonic/gin`: A web framework for Go.
- `github.com/playwright-community/playwright-go`: A Go client for Playwright.
- `github.com/go-redis/redis/v8`: A Redis client for Go.

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd tiktok-playwright
   ```

2. Install the dependencies:
   ```bash
   go mod tidy
   ```

3. Ensure you have Redis running locally or update the connection settings in `utils/redis.go`.

## Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

The server will start on port `3030`. You can access the API endpoints at `http://localhost:3030`.

## API Endpoints

### Search Videos

- **Endpoint**: `GET /search`
- **Query Parameters**:
  - `keyword`: The search term for TikTok videos.
- **Response**: Returns a JSON object containing a list of videos matching the search term.

### Play Video

- **Endpoint**: `GET /play`
- **Query Parameters**:
  - `url`: The URL of the TikTok video to play.
- **Response**: Redirects to the video URL. If the URL is cached in Redis, it will redirect to the cached URL.

## Code Explanation

### Main Application (`main.go`)

The main application initializes the Redis client, sets up the Gin router, and defines the routes for searching and playing videos.

### Handlers

- **SearchHandler**: Handles video search requests. It checks for cached results in Redis and fetches new results if not found.
- **PlayHandler**: Manages video playback requests. It checks for cached video URLs and redirects accordingly.

### Services

- **VideoService**: Contains the logic for fetching videos from TikTok using Playwright. It launches a headless browser, navigates to the search page, and extracts video details.

### Utilities

- **Redis Utilities**: Functions for interacting with Redis, including setting and getting cached values.
- **Response Utilities**: Functions for sending standardized JSON responses for success and error cases.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments

- [Playwright](https://playwright.dev/) for browser automation.
- [Gin](https://gin-gonic.com/) for the web framework.
- [Redis](https://redis.io/) for caching video URLs.
"# go-tiktok-playwright-scraper" 
"# go-tiktok-playwright-scraper"  git init git add README.md git commit -m "first commit" git branch -M main git remote add origin https://github.com/haidarrais/go-tiktok-playwright-scraper.git git push -u origin mainwh
