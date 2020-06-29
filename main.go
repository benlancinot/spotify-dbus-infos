package main

import (
	"log"
	"os"
	"strings"

	"github.com/godbus/dbus"
)

const (
	busName         = "org.mpris.MediaPlayer2.spotify"
	objectPath      = "/org/mpris/MediaPlayer2"
	playerInterface = "org.mpris.MediaPlayer2.Player"
)

type Spotify struct {
	*dbus.Object
}

func SpotifyBus(conn *dbus.Conn) *Spotify {
	return &Spotify{
		conn.Object(busName, objectPath).(*dbus.Object),
	}
}

func (spotify *Spotify) get(propName string) (interface{}, error) {
	variant, err := spotify.GetProperty(playerInterface + "." + propName)
	if err != nil {
		return nil, err
	}
	return variant.Value(), nil
}

func (spotify *Spotify) Infos() (pbi PlaybackInfo) {
	if pbs, err := spotify.get("PlaybackStatus"); err == nil {
		pbi.Status = pbs.(string)
	}

	m, err := spotify.get("Metadata")
	if err != nil {
		return
	}

	metadata := m.(map[string]dbus.Variant)
	pbi.Artist = strings.Join(metadata["xesam:artist"].Value().([]string), ", ")
	pbi.Title = metadata["xesam:title"].Value().(string)
	return
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	conn, err := dbus.SessionBus()
	if err != nil {
		logger.Println(err)
		return
	}
	bus := SpotifyBus(conn)
	logger.Println(bus.Infos())
}
