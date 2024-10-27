package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// List of Wikipedia pages to scrape // 要爬取的 Wikipedia 页面列表
var urls = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

// PageData structure to store page information // 用于存储页面信息的 PageData 结构
type PageData struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func main() {
	// Create a new collector // 创建一个新的爬虫收集器
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true), // Enable asynchronous requests for concurrency
	)

	// Record start time for Go scraper // 记录 Go 爬虫开始时间
	startTime := time.Now()

	// Use WaitGroup to wait for all scraping tasks to complete // 使用 WaitGroup 等待所有爬取任务完成
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Create a file to save JSON output // 创建文件用于保存 JSON 输出
	file, err := os.Create("output.jl")
	if err != nil {
		fmt.Println("Sry, Plz check ur file there is an error to create file:", err)
		return
	}
	defer file.Close()

	// Write text to the file for each paragraph found // 将每个找到的段落文本写入文件
	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.ChildText("p")
		pageData := PageData{
			URL:  e.Request.URL.String(),
			Text: text,
		}

		jsonData, err := json.Marshal(pageData)
		if err != nil {
			fmt.Println("Sry, there is an error to marshal JSON:", err)
			return
		}

		// Use a mutex to prevent race conditions when writing to the file // 使用互斥锁防止写文件时的竞态条件
		mutex.Lock()
		defer mutex.Unlock()
		_, err = file.WriteString(string(jsonData) + "\n")

		if err != nil {
			fmt.Println("Sry, there is an error to write to file:", err)
		}
	})

	// Start scraping task for each URL // 为每个 URL 启动爬取任务
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println("Scraping:", url)
			err := c.Visit(url)
			if err != nil {
				fmt.Println("Error visiting URL:", err)
			}
		}(url)
	}

	// Wait for all scraping tasks to complete // 等待所有爬取任务完成
	wg.Wait()

	// Allow the collector to finish all ongoing requests // 允许爬虫完成所有正在进行的请求
	c.Wait()

	// Record end time for Go scraper // 记录 Go 爬虫结束时间
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Time for this file to proceed by using go : %s\n", elapsedTime)

}
