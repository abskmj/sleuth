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
		Video{"2160", "265"},
	},
	{
		"Ammonite (2020) 720p 10bit BluRay x265 HEVC [ Org NF Hindi HE-AAC 5.1 ~ 192Kbps + English AAC 5.1 ] ESub ~ TheAvi ~ PSA.mkv",
		Video{"720", "265"},
	},
	{
		"Ammonite (2020) 1080p 10bit BluRay x265 HEVC [ Org NF Hindi DDP 5.1 ~ 640Kbps + English AAC 5.1 ] ESub ~ TheAvi ~ PSA.mkv",
		Video{"1080", "265"},
	},
	{
		"Mosagallu (2021) UNCUT 1080p AMZN WEB-DL x264 [Org YT Hindi AAC 5.1 ~384Kbps + Telugu DDP 5.1] ESub ~ BunnyJMB.mkv",
		Video{"1080", "264"},
	},
	{
		"Raavrambha.2023.1080p.AMZN.WEB-DL.DD+5.1.H.264-R3NGOKU.mkv",
		Video{"1080", "264"},
	},
	{
		"The.Out-Laws.2023.720p.NF.WEB-DL.AAC.2.0.H.264 -VijayRaj.mkv",
		Video{"720", "264"},
	},
	{
		"Invincible.S01E02.1080p.AMZN.WEB-DL.DDP.5.1.Multi.H.264-Vijayraj.mkv",
		Video{"1080", "264"},
	},
}

func TestVideo(t *testing.T) {

	for _, test := range tests {
		if output := NewVideo(test.name); output != test.expected {
			t.Errorf("\nTest: %q\nExpected: %q\nFound: %q", test.name, test.expected, output)
		}
	}
}
