//go:build darwin

package app

import (
	"testing"
)

func TestInstallMecab(t *testing.T) {

	// // firstly, uninstall mecab-ipadic. Then, uninstall mecab
	// if common.CmdExists("mecab") {
	// 	common.BrewUninstall("mecab-ipadic", "mecab-ipadic")
	// }
	// if common.BrewList("mecab-ipadic") {
	// 	common.BrewUninstall("mecab", "mecab")
	// }

	// InstallMecab()

	// if !common.CmdExists("mecab") {
	// 	t.Fatal("Package `mecab` is not installed")
	// }
	// if !common.BrewList("mecab-ipadic") {
	// 	t.Fatal("Package `mecab-ipadic` is not installed")
	// }
}
