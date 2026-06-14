package colly

import (
	gocolly "github.com/gocolly/colly"
)

type CollyTagCollector struct {
	steamBaseUrl string
	collector    gocolly.Collector
	tags         []string
}

const TAG_CLASS_NAME = ".tag_browse_tag"
const TAG_PAGE_PATH = "/tag/browse"

func (t *CollyTagCollector) All() ([]string, error) {

	t.collector.OnHTML(TAG_CLASS_NAME, func(h *gocolly.HTMLElement) {
		t.tags = append(t.tags, h.Text)
	})

	tagPageUrl := t.steamBaseUrl + TAG_PAGE_PATH

	err := t.collector.Visit(tagPageUrl)
	if err != nil {
		return []string{}, err
	}

	return t.tags, nil
}
