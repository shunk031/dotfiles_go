package common

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/ini.v1"
)

const SHELL = "/bin/bash"

func isDebug() bool {
	flag := os.Getenv("DEBUG_DOTFILES")
	if flag != "" {
		return true
	} else {
		return false
	}
}

func GetDotPath() (string, error) {
	dotPath := os.Getenv("DOTPATH")
	if dotPath != "" {
		return dotPath, nil
	} else {
		return "", fmt.Errorf("Need to specify DOTPATH")
	}
}

func Mkd(p string) {
	if len(p) > 0 {
		fInfo, err := os.Stat(p)
		if err != nil {
			log.Fatal(err)
		} else if errors.Is(err, os.ErrNotExist) {
			msg := p
			cmd := fmt.Sprintf("mkdir -p %s", p)
			err := Execute(msg, cmd)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if !!fInfo.IsDir() {
				printError(fmt.Errorf("%s - a file with the same name already exists", p))
			}
		}
	}
}

func ExecuteCmd(cmd string) error {

	c := exec.Command(SHELL, "-c", cmd)

	var out bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr

	err := c.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("command %s; %s", cmd, fmt.Sprint(err)))
	} else {
		return nil
	}
}

func Execute(msg string, cmd string) error {
	// ref. https://gist.github.com/bamoo456/7e21773e8ef742a726c041f5f0019d2e

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "  ["
	s.Suffix = fmt.Sprintf("] %s", cmd)

	s.Start()

	c := exec.Command(SHELL, "-c", cmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	err = c.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s.Suffix = fmt.Sprintf("] %s", scanner.Text())
	}
	c.Wait()

	s.Stop()

	printResult(msg, err)

	return err
}

func CmdExists(c string) bool {
	cmd := exec.Command("/bin/bash", "-c", "command", "-v", c)
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func extract(archive string, outputDir string) error {

	if CmdExists("tar") {
		msg := fmt.Sprintf("Extract from %s", archive)
		cmd := fmt.Sprintf("tar -zxf %s --strip-components 1 -C %s", archive, outputDir)
		return Execute(msg, cmd)
	} else {
		return fmt.Errorf("Command not found: tar")
	}

}

func readOsRelease(cfgFile string) map[string]string {
	cfg, err := ini.Load(cfgFile)
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	cfgParams := make(map[string]string)
	cfgParams["ID"] = cfg.Section("").Key("ID").String()
	return cfgParams

}

func getLinuxDistribution() string {
	osInfo := readOsRelease("/etc/os-release")
	return osInfo["ID"]
}

func GetOs() (string, error) {
	kernelName := runtime.GOOS
	if kernelName == "darwin" {
		return "macos", nil
	} else if kernelName == "linux" {
		return getLinuxDistribution(), nil
	} else {
		return kernelName, fmt.Errorf("Invalid OS: %s", kernelName)
	}
}

func GetCpuArch() string {
	out, err := exec.Command("uname", "-m").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
