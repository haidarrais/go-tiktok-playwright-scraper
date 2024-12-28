package services

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type Video struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Thumb string `json:"thumb"`
}

type VideoService struct{}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (vs *VideoService) FetchVideos(keyword string) ([]Video, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true), // Run headless
	})
	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %v", err)
	}

	searchURL := fmt.Sprintf("https://www.tiktok.com/search?q=%s", keyword)
	_, err = page.Goto(searchURL, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	if err != nil {
		return nil, fmt.Errorf("could not navigate to search URL: %v", err)
	}

	// Locate all matching elements
	videoLocator := page.Locator(".css-1soki6-DivItemContainerForSearch") // Adjust as per TikTok's structure
	videoCount, err := videoLocator.Count()
	fmt.Println(videoCount)
	if err != nil {
		return nil, fmt.Errorf("could not count video elements: %v", err)
	}

	videos := []Video{}
	for i := 0; i < videoCount; i++ {
		element := videoLocator.Nth(i)

		// Extract title
		title, err := element.Locator("h1").InnerText()
		if err != nil {
			return nil, fmt.Errorf("could not get video title: %v", err)
		}

		// Extract URL
		url, err := element.Locator("a").First().GetAttribute("href")
		fmt.Println(url)
		if err != nil || url == "" {
			continue
		}

		// Extract thumbnail
		thumb, err := element.Locator("img").First().GetAttribute("src")
		fmt.Println(thumb)
		if err != nil || thumb == "" {
			continue
		}

		videos = append(videos, Video{
			Title: title,
			URL:   url,
			Thumb: thumb,
		})
	}

	return videos, nil
}
