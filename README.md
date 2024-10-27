# web-scraper-by-go
Wikipedia Scraper using Go and Colly

Overview

This project is a simple web scraper written in Go using the Colly framework. The scraper is used to gather information from a list of Wikipedia pages related to intelligent systems and robotics. The collected data is saved in a JSON Lines file (output.jl). The main goal is to demonstrate concurrent scraping using Go's goroutines and Colly's asynchronous capabilities, making the scraper efficient and fast.

Features

Scrapes content from a list of Wikipedia pages.

Extracts text data from paragraph tags.

Utilizes Go's concurrency with goroutines for efficient scraping.

Saves the scraped data as JSON Lines (.jl), where each line is a separate JSON object containing the URL and its text content.

Prerequisites

Go installed on your system (version 1.14 or above).

Git installed to clone the repository.

Installation

Clone the repository:

git clone <repository_url>
cd <repository_folder>

Install the Colly dependency:

go get github.com/gocolly/colly/v2

How to Run

Compile the Go program:

go build -o scraper main.go

This will create an executable file named scraper (scraper.exe on Windows).

Run the executable:

./scraper

The program will scrape the Wikipedia pages and save the data in output.jl.

Output

The output is stored in a file named output.jl in JSON Lines format, where each line contains information about one scraped Wikipedia page. Each JSON object contains:

url: The URL of the page.

text: The scraped text content from the page.

Example of output.jl:

{"url": "https://en.wikipedia.org/wiki/Robotics", "text": "Robotics is..."}
{"url": "https://en.wikipedia.org/wiki/Robot", "text": "A robot is..."}

Code Explanation

URLs to Scrape: The urls variable contains a list of Wikipedia pages to scrape.

Concurrency: Uses WaitGroup and Mutex to manage concurrent scraping and safe file writes.

Colly Collector: Configured with AllowedDomains to restrict scraping to Wikipedia.

OnHTML Callback: Collects the text from paragraph tags (<p>) on each page and writes the data to output.jl.

Error Handling

The program includes basic error handling to print errors related to:

File creation (output.jl).

JSON marshalling of page data.

Writing data to the output file.

Visiting URLs.

How to Test

To verify the scraper's output, you can open the output.jl file after running the program and ensure the scraped content is correctly stored.

Modify the urls list to include other Wikipedia pages to test scraping on different topics.

Compare the processing time by reviewing the output at the end, which displays the total time taken to scrape all the URLs.

Notes

The scraper is designed specifically for Wikipedia pages and may need adjustments for other websites.

The output is in JSON Lines format, which is commonly used for data ingestion into databases or other applications.

License

This project is licensed under the MIT License. See the LICENSE file for details.

