package main

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type test struct {
	name     string
	expected Video
}

var tests = []test{
	{
		"Secret.Invasion.S01E03.Betrayed.2160p.WEB-DL.DUAL.DDP5.1.Atmos.H.265-HDRMoviezZ.mkv",
		Video{"2160"},
	},
	{
		"Ammonite (2020) 720p 10bit BluRay x265 HEVC [ Org NF Hindi HE-AAC 5.1 ~ 192Kbps + English AAC 5.1 ] ESub ~ TheAvi ~ PSA.mkv",
		Video{"720"},
	},
	{
		"Ammonite (2020) 1080p 10bit BluRay x265 HEVC [ Org NF Hindi DDP 5.1 ~ 640Kbps + English AAC 5.1 ] ESub ~ TheAvi ~ PSA.mkv",
		Video{"1080"},
	},
}

func TestVideo(t *testing.T) {

	for _, test := range tests {
		if output := NewVideo(test.name); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
