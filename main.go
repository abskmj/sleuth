package main

import (
	"flag"
	"fmt"

	"github.com/abskmj/sleuth/lib"
)

func main() {
	// Define a command-line flag named "name" with a default value
	name := flag.String("name", "Secret.Invasion.S01E03.Betrayed.2160p.WEB-DL.DUAL.DDP5.1.Atmos.H.265-HDRMoviezZ.mkv", "The name to greet")
	flag.Parse()

	video := lib.NewVideo(*name)

	fmt.Println("Video:", video)
	fmt.Println("Resolution:", video.Resolution())
	fmt.Println("isShow:", video.IsShow())

	// Print the greeting message
	// message := fmt.Sprintf("Hello, %s!", *name)
	// fmt.Println(message)
}
