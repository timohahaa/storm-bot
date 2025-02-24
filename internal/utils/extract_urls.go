package utils

import "mvdan.cc/xurls/v2"

var extractor = xurls.Relaxed()

func ExtractURLs(text string) []string {
	return extractor.FindAllString(text, -1)
}
