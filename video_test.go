package main

import "testing"

type test struct {
	name     string
	expected Video
}

var tests = []test{
	{
		"Secret.Invasion.S01E03.Betrayed.2160p.WEB-DL.DUAL.DDP5.1.Atmos.H.265-HDRMoviezZ.mkv",
		Video{
			title:           "Secret Invasion",
			category:        "episode",
			year:            "",
			season:          "01",
			episode:         "03",
			videoResolution: "2160",
			videoCodec:      "265",
			audioChannels:   "5.1",
			audioCodec:      "DDP",
		},
	},
	{
		"Ammonite (2020) 720p 10bit BluRay x265 HEVC [ Org NF Hindi HE-AAC 5.1 ~ 192Kbps + English AAC 5.1 ] ESub ~ TheAvi ~ PSA.mkv",
		Video{
			title:           "Ammonite",
			category:        "movie",
			year:            "2020",
			season:          "",
			episode:         "",
			videoResolution: "720",
			videoCodec:      "265",
			audioChannels:   "5.1",
			audioCodec:      "AAC",
		},
	},
	{
		"Ammonite (2020) 1080p 10bit BluRay x265 HEVC [ Org NF Hindi DDP 5.1 ~ 640Kbps + English AAC 5.1 ] ESub ~ TheAvi ~ PSA.mkv",
		Video{
			title:           "Ammonite",
			category:        "movie",
			year:            "2020",
			season:          "",
			episode:         "",
			videoResolution: "1080",
			videoCodec:      "265",
			audioChannels:   "5.1",
			audioCodec:      "DDP",
		},
	},
	{
		"Mosagallu (2021) UNCUT 1080p AMZN WEB-DL x264 [Org YT Hindi AAC 5.1 ~384Kbps + Telugu DDP 5.1] ESub ~ BunnyJMB.mkv",
		Video{
			title:           "Mosagallu",
			category:        "movie",
			year:            "2021",
			season:          "",
			episode:         "",
			videoResolution: "1080",
			videoCodec:      "264",
			audioChannels:   "5.1",
			audioCodec:      "DDP",
		},
	},
	{
		"Raavrambha.2023.1080p.AMZN.WEB-DL.DD+5.1.H.264-R3NGOKU.mkv",
		Video{
			title:           "Raavrambha",
			category:        "movie",
			year:            "2023",
			season:          "",
			episode:         "",
			videoResolution: "1080",
			videoCodec:      "264",
			audioChannels:   "5.1",
			audioCodec:      "DD",
		},
	},
	{
		"The.Out-Laws.2023.720p.NF.WEB-DL.AAC.2.0.H.264 -VijayRaj.mkv",
		Video{
			title:           "The Out Laws",
			category:        "movie",
			year:            "2023",
			season:          "",
			episode:         "",
			videoResolution: "720",
			videoCodec:      "264",
			audioChannels:   "2.0",
			audioCodec:      "AAC",
		},
	},
	{
		"Invincible.S01E02.1080p.AMZN.WEB-DL.DDP.5.1.Multi.H.264-Vijayraj.mkv",
		Video{
			title:           "Invincible",
			category:        "episode",
			year:            "",
			season:          "01",
			episode:         "02",
			videoResolution: "1080",
			videoCodec:      "264",
			audioChannels:   "5.1",
			audioCodec:      "DDP",
		},
	},
	{
		"Serenity (2005) 1080p BluRay REMUX VC-1 [Hindi DDP 5.1 + English DTS-HD MA 5.1] x264-[ARRoW] FraMeSToR.mkv",
		Video{
			title:           "Serenity",
			category:        "movie",
			year:            "2005",
			season:          "",
			episode:         "",
			videoResolution: "1080",
			videoCodec:      "264",
			audioChannels:   "5.1",
			audioCodec:      "DTS-HD",
		},
	},
	{
		"Knight.and.Day.2010.Extended.1080p.BluRay.REMUX.AVC.[Hindi BD 5.1 + English DTS-HD MA 5.1].ESubs-(ARRoW) FraMeSToR.mkv",
		Video{
			title:           "Knight and Day",
			category:        "movie",
			year:            "2010",
			season:          "",
			episode:         "",
			videoResolution: "1080",
			videoCodec:      "",
			audioChannels:   "5.1",
			audioCodec:      "DTS-HD",
		},
	},
	{
		"Unknown.The.Lost.Pyramid.2023.1080p.NF.WEB-DL.DUAL.DDP5.1.Atmos.HDR-DV.H.265-Telly.mkv",
		Video{
			title:           "Unknown The Lost Pyramid",
			category:        "movie",
			year:            "2023",
			season:          "",
			episode:         "",
			videoResolution: "1080",
			videoCodec:      "265",
			videoQuality:    "DV, HDR",
			audioChannels:   "5.1",
			audioCodec:      "DDP",
		},
	},
	{
		"Oblivion.2013.2160p.UHD.BluRay.HDR.x265.[Hindi DDP 5.1 + English TrueHD 7.1].ESubs-(ARRoW) CtrlHD.mkv",
		Video{
			title:           "Oblivion",
			category:        "movie",
			year:            "2013",
			season:          "",
			episode:         "",
			videoResolution: "2160",
			videoCodec:      "265",
			videoQuality:    "HDR",
			audioChannels:   "7.1",
			audioCodec:      "TrueHD",
		},
	},
	{
		"Star.Trek.Strange.New.Worlds.S02E03.Tomorrow.and.Tomorrow.and.Tomorrow.2160p.PMTP.WEB-DL.DoVi.HEVC.[Hindi DDP 2.0 + English DDP 5.1].ESubs-(ArroW) NTb.mkv",
		Video{
			title:           "Star Trek Strange New Worlds",
			category:        "episode",
			year:            "",
			season:          "02",
			episode:         "03",
			videoResolution: "2160",
			videoCodec:      "265",
			videoQuality:    "DV",
			audioChannels:   "5.1",
			audioCodec:      "DDP",
		},
	},
}

func TestVideo(t *testing.T) {

	for _, test := range tests {
		if output := NewVideo(test.name); output != test.expected {
			t.Errorf("\nTest: %q\nExpected: %q\nFound: %q", test.name, test.expected, output)
		}
	}
}
