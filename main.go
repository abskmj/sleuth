package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "", "name of the file")
	flag.Parse()

	fmt.Println("Name:", *name)

	var video = NewVideo(*name)

	fmt.Println("Title:", video.title)
	fmt.Println("Category:", video.category)
	fmt.Println("Year:", video.year)
	fmt.Println("Season:", video.season)
	fmt.Println("Episode:", video.episode)
	fmt.Println("Video Resolution:", video.videoResolution)
	fmt.Println("Video Codec:", video.videoCodec)
	fmt.Println("Video Quality:", video.videoQuality)
	fmt.Println("Audio Channels:", video.audioChannels)
	fmt.Println("Audio Codec:", video.audioCodec)
}
