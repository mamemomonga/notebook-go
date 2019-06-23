package main

import ()

func main() {
	dc := NewMyDockerClient()
	dc.ServerVersion()
	dc.LaunchAlpine()
}
