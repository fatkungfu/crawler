# Go Web Crawler

_Made with help form boot.dev_
A concurrent web crawler written in Go. Crawls a website, counts internal links to each page, and outputs a sorted report to the console and a CSV file.

## Features

- Concurrent crawling with configurable concurrency
- Limits the number of pages crawled
- Prints a sorted report of internal links
- Saves the report as `report.csv`

## Usage

Build the crawler:

```sh
go build -o crawler
```

Run the crawler:

```sh
./crawler <BASE_URL> <maxConcurrency> <maxPages>
```

Or run directly with Go:

```sh
go run . <BASE_URL> <maxConcurrency> <maxPages>
```

- <BASE_URL>: The root URL to start crawling (e.g., https://example.com)
- <maxConcurrency>: Maximum number of concurrent requests (e.g., 5)
- <maxPages>: Maximum number of pages to crawl (e.g., 20)
