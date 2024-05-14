# Log Ingestor and Query Interface

## Overview
This project consists of a log ingestor built with Go and a query interface built with React.js and Tailwind CSS. The ingestor handles vast volumes of log data, and the interface allows querying this data using full-text search or specific field filters.

## Features
- **Log Ingestor**
  - Ingest logs via HTTP POST requests on port 3000.
  - Store logs in Elasticsearch for efficient search capabilities.
- **Query Interface**
  - Full-text search across logs.
  - Filters based on various fields (level, message, resourceId, etc.).
  - Real-time log ingestion and searching.

## Getting Started

### Prerequisites
- Go 1.23+
- Node.js 14+
- Elasticsearch 7.10+

### Running the Log Ingestor
1. Clone the repository:
    ```bash
    git clone https://github.com/chrisdamba/go-react-log-seer.git
    cd log-ingestor
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Start the log ingestor server:
    ```bash
    go run main.go
    ```

### Running the Query Interface
1. Navigate to the React project directory:
    ```bash
    cd log-query-interface
    ```

2. Install dependencies:
    ```bash
    npm install
    ```

3. Start the development server:
    ```bash
    npm start
    ```

### Sample Data
To ingest sample logs, send a POST request to `http://localhost:3000/logs` with the log data in JSON format:
```json
{
    "level": "error",
    "message": "Failed to connect to DB",
    "resourceId": "server-1234",
    "timestamp": "2023-09-15T08:00:00Z",
    "traceId": "abc-xyz-123",
    "spanId": "span-456",
    "commit": "5e5342f",
    "metadata": {
        "parentResourceId": "server-0987"
    }
}
