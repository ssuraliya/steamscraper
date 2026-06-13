package colly

import (
	collectors "steamscraper/internal/collectors"

	gocolly "github.com/gocolly/colly"
)

type CollyTagCollector struct {
	tags []string
}

const TAG_CLASS_NAME = ".tag_browse_tag"
const TAG_PAGE_URL = "https://store.steampowered.com/tag/browse"

func (t *CollyTagCollector) All() ([]string, error) {

	collector := collectors.NewCollector()

	collector.OnHTML(TAG_CLASS_NAME, func(h *gocolly.HTMLElement) {
		t.tags = append(t.tags, h.Text)
	})

	err := collector.Visit(TAG_PAGE_URL)
	if err != nil {
		return []string{}, err
	}

	return t.tags, nil
}
