//go:build darwin

package initialize

func RunInitCmd() error {
	runInitCommon()

	// err := runInitMacOsCommon()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return runInitMacOsArch()
}
