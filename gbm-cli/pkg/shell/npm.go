package shell

import (
	"os"
	"os/exec"
)

type NpmCmds interface {
	Install(...string) error
	Ci() error
	Run(...string) error
	RunIn(string, ...string) error
	Version(string) error
	VersionIn(string, string) error
}

func (c *client) Ci() error {

	// If NVM is installed, use it to install the correct version of node
	// This should be added under a $CI flag or something. I'm just noticing
	// that Github workflows are using a different process for each exec.Command
	// So even though we switch to the correct version of node earlier on, when we call
	// npm ci, it's using the default version of node.
	if os.Getenv("NVM_DIR") != "" {
		cmd := exec.Command("bash", "-l", "-c", "nvm use && npm ci")
		cmd.Dir = c.dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return c.cmd("ci")
}

func (c *client) Run(args ...string) error {
	if os.Getenv("NVM_DIR") != "" {
		run := append([]string{"-l", "-c","nvm use && npm run "}, args...)
		cmd := exec.Command("bash", run...)
		cmd.Dir = c.dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	run := append([]string{"run"}, args...)
	return c.cmd(run...)
}

func (c *client) RunIn(path string, args ...string) error {
	run := append([]string{"run"}, args...)
	return c.cmdInPath(path, run...)
}

func (c *client) Version(version string) error {
	// Let's not add the tag by default.
	// If we need it we should consider a different function.
	versionCmd := []string{"version", version, "--no-git-tag=false"}
	return c.cmd(versionCmd...)
}

func (c *client) VersionIn(packagePath, version string) error {
	versionCmd := []string{"version", version, "--no-git-tag=false"}
	return c.cmdInPath(packagePath, versionCmd...)
}
