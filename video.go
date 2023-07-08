package main

import (
	"regexp"
	"strings"
)

var reSeasonEpisode = regexp.MustCompile(`(?i)s([0-9]{1,2})e([0-9]{1,2})`)
var re2160p = regexp.MustCompile(`(?i) (2160p) `)
var re1080p = regexp.MustCompile(`(?i) (1080p) `)
var re720p = regexp.MustCompile(`(?i) (720p) `)

type Video struct {
	resolution string
}

func is2160p(name string) bool {
	return find(name, *re2160p) != ""
}

func is1080p(name string) bool {
	return find(name, *re1080p) != ""
}

func is720p(name string) bool {
	return find(name, *re720p) != ""
}

func getResolution(name string) string {
	var resolution = ""

	if is2160p(name) {
		resolution = "2160"
	} else if is1080p(name) {
		resolution = "1080"
	} else if is720p(name) {
		resolution = "720"
	}

	return resolution
}

func IsShow(name string) bool {
	return reSeasonEpisode.FindString(name) != ""
}

func Season(name string) string {
	return find(name, *reSeasonEpisode)
}

func find(text string, re regexp.Regexp) string {
	var match = ""
	var matches = re.FindStringSubmatch(text)

	if len(matches) > 1 {
		match = matches[0]
	}

	return match
}

func sanitize(name string) string {
	var replacer = strings.NewReplacer(
		".", " ",
		"-", " ",
	)

	return replacer.Replace(name)
}

func NewVideo(name string) Video {
	var v = Video{}

	var sanitized = sanitize(name)

	v.resolution = getResolution(sanitized)

	return v
}
