package handlers

import (
	"os/exec"
)

type MPlayer struct {}

func (MPlayer *MPlayer) Play(fileName string) error {
	mplayer := exec.Command("mplayer", "-cache", "106092", "-", fileName)
	return mplayer.Run()
}
