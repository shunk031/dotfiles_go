//go:build linux

package app

import (
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallTmux(t *testing.T) {
	InstallTmux()

	if !common.CmdExists("tmux") {
		t.Fatal("Command not found: tmux")
	}

	if !common.CmdExists("xsel") {
		t.Fatal("Command not found: xsel")
	}

	if !common.CmdExists("cmake") {
		t.Fatal("Command not found: cmake")
	}
}
