package handlers

import (
	"os/exec"
)

type MPlayer struct {
	Volume int
}

func (MPlayer *MPlayer) Play(fileName string) error {
	mplayer := exec.Command("mplayer", "-cache", "106092", "-", fileName, "-af", "volume="+string(MPlayer.Volume))
	return mplayer.Run()
}
