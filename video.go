package main

import (
	"regexp"
	"strings"
)

// utilities
func find(text string, re regexp.Regexp) string {
	var match = ""
	var matches = re.FindStringSubmatch(text)

	if len(matches) > 1 {
		match = matches[1]
	}

	return match
}

func has(text string, re regexp.Regexp) bool {
	return find(text, re) != ""
}

var reSanitize = regexp.MustCompile(`(?i)[()\[\]~\.\-\+\s]+`)

func sanitize(name string) string {
	return reSanitize.ReplaceAllString(name, " ")
}

// Video Resolutions
var re2160p = regexp.MustCompile(`(?i) (2160p) `)
var re1080p = regexp.MustCompile(`(?i) (1080p) `)
var re720p = regexp.MustCompile(`(?i) (720p) `)

func is2160p(name string) bool {
	return find(name, *re2160p) != ""
}

func is1080p(name string) bool {
	return find(name, *re1080p) != ""
}

func is720p(name string) bool {
	return find(name, *re720p) != ""
}

func getVideoResolution(name string) string {
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

// Video Codec
var re265 = regexp.MustCompile(`(?i) (HEVC|x265|H 265) `)
var re264 = regexp.MustCompile(`(?i) (x264|H 264) `)

func getVideoCodec(name string) string {
	var codec = ""

	if has(name, *re265) {
		codec = "265"
	} else if has(name, *re264) {
		codec = "264"
	}

	return codec
}

// video quality
var reDV = regexp.MustCompile(`(?i) (DV) `)
var reHDR = regexp.MustCompile(`(?i) (HDR) `)

func getVideoQuality(name string) string {
	var qualities []string

	if has(name, *reDV) {
		qualities = append(qualities, "DV")
	}
	if has(name, *reHDR) {
		qualities = append(qualities, "HDR")
	}

	return strings.Join(qualities, ", ")
}

// audio channels
var re51Channels = regexp.MustCompile(`(?i) (5 1|DDP5 1) `)
var re20Channels = regexp.MustCompile(`(?i) (2 0) `)

func getAudioChannels(name string) string {
	var channels = ""

	if has(name, *re51Channels) {
		channels = "5.1"
	} else if has(name, *re20Channels) {
		channels = "2.0"
	}

	return channels
}

// audio codec
var reAudioDTSHD = regexp.MustCompile(`(?i) (DTS HD) `)
var reAudioDDP = regexp.MustCompile(`(?i) (DDP|DDP5) `)
var reAudioDD = regexp.MustCompile(`(?i) (DD) `)
var reAudioAAC = regexp.MustCompile(`(?i) (AAC) `)

func getAudioCodec(name string) string {
	var codec = ""

	if has(name, *reAudioDTSHD) {
		codec = "DTS-HD"
	} else if has(name, *reAudioDDP) {
		codec = "DDP"
	} else if has(name, *reAudioDD) {
		codec = "DD"
	} else if has(name, *reAudioAAC) {
		codec = "AAC"
	}

	return codec
}

// episode
var reSeason = regexp.MustCompile(`(?i) s([0-9]{1,2})e[0-9]{1,2} `)
var reEpisode = regexp.MustCompile(`(?i) s[0-9]{1,2}e([0-9]{1,2}) `)

func getSeason(name string) string {
	return find(name, *reSeason)
}

func getEpisode(name string) string {
	return find(name, *reEpisode)
}

// year
var reYear = regexp.MustCompile(`(?i) ((19|20)[0-9]{2}) `)

func getYear(name string) string {
	return find(name, *reYear)
}

// title
var reTitle = regexp.MustCompile(`(?i)(.+?) ((19|20)[0-9]{2}|s[0-9]{1,2}e[0-9]{1,2}) `)

func getTitle(name string) string {
	return find(name, *reTitle)
}

// Video
type Video struct {
	title           string
	year            string
	season          string
	episode         string
	videoResolution string
	videoCodec      string
	videoQuality    string
	audioChannels   string
	audioCodec      string
}

func NewVideo(name string) Video {
	var sanitized = sanitize(name)

	var v = Video{
		title:           getTitle(sanitized),
		year:            getYear(sanitized),
		season:          getSeason(sanitized),
		episode:         getEpisode(sanitized),
		videoResolution: getVideoResolution(sanitized),
		videoCodec:      getVideoCodec(sanitized),
		videoQuality:    getVideoQuality(sanitized),
		audioChannels:   getAudioChannels(sanitized),
		audioCodec:      getAudioCodec(sanitized),
	}

	return v
}
