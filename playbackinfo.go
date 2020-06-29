package main

import (
	"fmt"
)

type PlaybackInfo struct {
	Status string
	Artist string
	Title  string
}

func (pbi PlaybackInfo) String() string {
	if len(pbi.Artist) == 0 && len(pbi.Title) == 0 {
		return ""
	}
	if len(pbi.Artist) == 0 {
		pbi.Artist = "Unknown"
	}
	if len(pbi.Title) == 0 {
		pbi.Title = "Unknown"
	}
	return fmt.Sprintf("%s  %s - %s", icons[pbi.Status], pbi.Artist, pbi.Title)
}

var icons = map[string]string{
	"Playing": "",
	"Paused":  "",
	"Stopped": "",
}
