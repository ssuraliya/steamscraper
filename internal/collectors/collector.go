package collectors

import (
	gocolly "github.com/gocolly/colly"
)

const STEAM_BASE_URL = "store.steampowered.com"

func NewCollector() gocolly.Collector {
	collector := gocolly.NewCollector(
		gocolly.AllowedDomains(STEAM_BASE_URL),
	)
	return *collector
}
