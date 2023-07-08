package main

import (
	"fmt"
	"regexp"
)

var reSeasonEpisode = regexp.MustCompile(`(?i)s([0-9]{1,2})e([0-9]{1,2})`)

// utilities
func find(text string, re regexp.Regexp) string {
	var match = ""
	var matches = re.FindStringSubmatch(text)

	if len(matches) > 1 {
		match = matches[0]
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

func IsShow(name string) bool {
	return reSeasonEpisode.FindString(name) != ""
}

func Season(name string) string {
	return find(name, *reSeasonEpisode)
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

var reAudioDDP = regexp.MustCompile(`(?i) (DDP|DDP5) `)
var reAudioDD = regexp.MustCompile(`(?i) (DD) `)
var reAudioAAC = regexp.MustCompile(`(?i) (AAC) `)

func getAudioCodec(name string) string {
	var codec = ""

	if has(name, *reAudioDDP) {
		codec = "DDP"
	} else if has(name, *reAudioDD) {
		codec = "DD"
	} else if has(name, *reAudioAAC) {
		codec = "AAC"
	}

	return codec
}

// Video

type Video struct {
	videoResolution string
	videoCodec      string
	audioChannels   string
	audioCodec      string
}

func NewVideo(name string) Video {
	var v = Video{}

	var sanitized = sanitize(name)

	fmt.Println(sanitized)

	v.videoResolution = getVideoResolution(sanitized)
	v.videoCodec = getVideoCodec(sanitized)
	v.audioChannels = getAudioChannels(sanitized)
	v.audioCodec = getAudioCodec(sanitized)

	return v
}
