package lib

import (
	"regexp"
	"strings"
)

var reSeasonEpisode = regexp.MustCompile(`(?i)s[0-9]{1,2}e[0-9]{1,2}`)
var reVideoResolution = regexp.MustCompile(`(?i)[0-9]{3,4}p`)

type Video struct {
	name         string
	santizedName string
}

func (v Video) Resolution() string {
	return reVideoResolution.FindString(v.santizedName)
}

func (v Video) IsShow() bool {
	return reSeasonEpisode.FindString(v.santizedName) != ""
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

	v.name = name
	v.santizedName = sanitize(name)

	return v
}
