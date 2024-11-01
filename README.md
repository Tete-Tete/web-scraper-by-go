# Go Wikipedia Scraper Tool

This repository contains a Go implementation for scraping Wikipedia pages, using the Colly framework for efficient data collection. It includes a main.go file, which serves as the entry point, and performs scraping on specified Wikipedia URLs to save the content in a JSON Lines format file.

## Overview
This project is a simple web scraper written in Go using the Colly framework. The scraper is used to gather information from a list of Wikipedia pages related to intelligent systems and robotics. The collected data is saved in a JSON Lines file (output.jl). The main goal is to provide a simple and effective tool for scraping Wikipedia pages while utilizing Go's concurrency features. This project helps demonstrate the power of Go's goroutines and the Colly framework for handling web scraping tasks efficiently.

## Features
- **Efficient Web Scraping**: Scrapes content from a list of Wikipedia pages using Colly.
- **Concurrency**: Utilizes Go's goroutines for concurrent scraping, enhancing performance.
- **JSON Lines Output**: Saves scraped data in JSON Lines format, with each line containing the URL and its associated text content.

## Requirements
- Go programming language installed (version 1.16 or higher recommended).
- Git: Required to clone the repository.

## Installation From Git and Setup
### Step 1: Clone the Repository
Clone this repository to your local machine:
```sh
git clone <https://github.com/Tete-Tete/web-scraper-by-go.git>
```
### Step 2: Install Dependencies
Install the Colly dependency:
```sh
go get github.com/gocolly/colly/v2
```

### Step 3: Run the Application
To compile the Go program, run the following command:
```sh
go build -o scraper main.go
```
This will create an executable file named scraper. The program will scrape the specified Wikipedia pages and save the data in output.jl.


## Testing
### Running Tests
To verify the correctness of the scraper, you can run tests by modifying the urls list in main.go and ensuring the output in output.jl matches the expected scraped content. You can also check the time taken for scraping, which is displayed after the scraping completes.

## Output
The output is saved in a file named output.jl in JSON Lines format. Each line contains the scraped data for a Wikipedia page, including: url and text. 

## Time Comparison
The Go Wikipedia Scraper Tool records the time taken for scraping all the specified pages. You can compare this runtime with a Python implementation to understand the performance differences. The python example is called WebFocusedCrawlWorkV001.zip. 

## Conclusion
The Go Wikipedia Scraper Tool effectively demonstrates the power of Go's concurrency features and the Colly framework for efficiently scraping content from multiple Wikipedia pages. By utilizing goroutines, the scraper can handle multiple requests simultaneously, making it an excellent solution for large-scale web scraping tasks. This project serves as a practical example of how to harness Go's capabilities for data collection, and it can be further extended to scrape other websites or customized to meet specific needs.

