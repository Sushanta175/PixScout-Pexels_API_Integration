# Go Pexels API Client

## Overview

This is a Go client library for interacting with the Pexels API, allowing you to easily search and retrieve photos and videos from the Pexels platform.

## Features

- Search photos by query
- Search videos by query
- Retrieve curated photos
- Get popular videos
- Fetch random photos and videos
- Rate limit tracking

## Prerequisites

- Go 1.16 or higher
- Pexels API Token

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Sushanta175/Go_Pexels_API.git
cd Go_Pexels_API
```

2. Install dependencies:
```bash
go mod tidy
```

## Configuration

1. Create a `.env` file in the root directory
2. Add your Pexels API token:
```
Pexels_Token=your_api_token_here
```

## Usage Examples

### Searching Photos

```go
package main

import (
    "fmt"
    "log"

    "github.com/Sushanta175/Go_Pexels_API/client"
    "github.com/Sushanta175/Go_Pexels_API/config"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Create a new Pexels client
    cfg := config.LoadConfig()
    c := client.NewClient(cfg.ApiToken)

    // Search for photos
    result, err := c.SearchPhotos("waves", 15, 1)
    if err != nil {
        log.Fatalf("Search Error: %v", err)
    }

    // Print remaining API requests
    fmt.Printf("Remaining API Requests: %d\n", c.RemainingRequests())

    // Process the results
    for _, photo := range result.Photos {
        fmt.Printf("Photo URL: %s\n", photo.Url)
    }
}
```

### Searching Videos

```go
// Similar to photo search, but use SearchVideo method
result, err := c.SearchVideo("nature", 10, 1)
```

### Getting Random Photo or Video

```go
randomPhoto, err := c.GetRandomPhoto()
randomVideo, err := c.GetRandomVideo()
```

## Rate Limit Handling

The client automatically tracks remaining API requests:

```go
// Check remaining requests
remainingRequests := c.RemainingRequests()

// Check if requests are low
if c.IsRateLimitLow(10) {
    log.Println("Low API requests remaining")
}
```

## Error Handling

The library provides detailed error messages for:
- API request failures
- Authentication errors
- JSON parsing errors

## Dependencies

- `github.com/joho/godotenv` for environment variable management
- Standard Go libraries for HTTP requests and JSON parsing

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Sushanta Sahu - github.com/Sushanta175

Project Link: [https://github.com/Sushanta175/Go_Pexels_API](https://github.com/Sushanta175/Go_Pexels_API)

---

## 🚀 Web Explorer & DevOps Add-On

The project now includes a **web GUI** and container-ready deployment for showcasing DevOps workflows.

### Features

| Feature | Details |
|---------|---------|
| Web GUI | Search photos in the browser using htmx + Go templates |
| REST API | `/api/photos` and `/api/videos` endpoints |
| SPA | Fully served from Go (no external web server) |
| Docker | Multi-stage build, small distroless runtime image |
| Kubernetes | Deployment, Service, Secret manifests included |
| GitOps | Ready for Argo CD auto-sync |

### Quick Start (Web GUI)

```bash
go run ./cmd/web
# Open http://localhost:8080
```

### Docker

```bash
docker build -t pexels-app .
docker run -p 8080:8080 --env-file=.env pexels-app
```

### Kubernetes

```bash
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/
```

### Repository Layout Additions

```
cmd/web/         # Web server entry point
web/templates/   # HTML templates
web/static/      # CSS + JS
k8s/             # Kubernetes manifests
Dockerfile       # Multi-stage Docker build
```

