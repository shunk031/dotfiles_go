//go:build darwin && amd64

package initialize

import (
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize/app"
)

func runInitMacOsArch() error {

	if err := app.InstallIterm2(); err != nil {
		return err
	}
	if err := app.InstallTmux(); err != nil {
		return err
	}
	if err := app.InstallMecab(); err != nil {
		return err
	}
	if err := app.InstallMecabIpadicNeologd(); err != nil {
		return err
	}
	if err := app.InstallPowerlevel10kRequirements(); err != nil {
		return err
	}
	return nil
}
