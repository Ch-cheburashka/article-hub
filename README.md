# Documentation

A simple container for articles that allows you to store and download articles in it and search them by a unique identifier.

## Table of Contents
1. [Usage](#usage)
2. [API endpoints](#api-endpoints)
3. [Data Models](#data-models)

## Usage

### Prerequisites

- **Go**: Ensure Go (version 1.16 or later) is installed on your system.
### Steps

1. **Clone the Repository**
```bash
git clone https://github.com/Ch-cheburashka/article-hub
cd article-hub
```
2. **Install Dependencies**
Run the following command to install any required Go modules:
```bash
go mod tidy
```

3. **Run the Server**
To start the server, use the following command:
```bash
go run main.go
```

The server will start on http://localhost:8080 ***by default***. You can specify a different port using the `-port=... flag`.

## API Endpoints
***1. Add an Article***
- Endpoint: `/add`
- Method: `POST`
- Request: JSON-object with the following format: `{ "title": "Example Title", "content": "Example Content", "id"; 1234 }`
- Response:
    - 200 OK: "Article added successfully".

***2. Search***
- Endpoint: `/search`
- Method: `GET`
- Query Parameter: query - the search query string. `http://localhost:8080/search?query=1234`
- Response:
    - 200 OK: Returns a JSON-object of the following format: `{ "title": "Example Title", "content": "Example Content", "id"; 1234 }`


## Data Models
### Article
- ID: Unique identifier for the article (the number of added article).
- Title
- Content
- ID