package crawler

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		url := page.URL
		count := page.Count
		fmt.Printf("Found %d internal links to %s\n", count, url)
	}
}

func sortPages(pages map[string]int) []Page {
	pagesSlice := []Page{}
	for url, count := range pages {
		pagesSlice = append(pagesSlice, Page{URL: url, Count: count})
	}

	// Sort by count descending, then url ascending
	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].Count == pagesSlice[j].Count {
			return pagesSlice[i].URL < pagesSlice[j].URL
		}
		return pagesSlice[i].Count > pagesSlice[j].Count
	})
	return pagesSlice
}

func SaveReportCSV(pages map[string]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"URL", "Count"}); err != nil {
		return err
	}

	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		record := []string{page.URL, fmt.Sprintf("%d", page.Count)}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
