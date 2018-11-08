package main

import humanize "github.com/dustin/go-humanize"
import "fmt"

func main() {
  fmt.Println("Hello World !!")
  fmt.Printf("Sample File Size in MB: %s.\n", humanize.Bytes(48234496)) // Sample File Size in MB: 46MB
}
