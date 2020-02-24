package e2e

import (
	"bufio"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	registryContainerName = "github-actions-registry"
	githubActionsImage    = "github-actions-e2e"
)

type envVar struct {
	key   string
	value string
}

func parseEnvFile(envFile string) ([]envVar, error) {
	vars := []envVar{}

	wd, err := os.Getwd()
	if err != nil {
		return vars, err
	}

	file, err := os.Open(path.Join(wd, envFile))
	if err != nil {
		return vars, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		vars = append(vars, envVar{split[0], split[1]})
	}

	return vars, scanner.Err()
}

func setupEnvVars(vars []envVar) error {
	for _, v := range vars {
		if err := os.Setenv(v.key, v.value); err != nil {
			return err
		}
	}
	return nil
}

func removeEnvVars(vars []envVar) error {
	for _, v := range vars {
		if err := os.Unsetenv(v.key); err != nil {
			return err
		}
	}
	return nil
}

func setupLocalRegistry() error {
	_ = removeLocalRegistry()

	cmd := exec.Command("docker", "run", "-d", "-p", "5000:5000", "--name", registryContainerName, registryContainerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func removeLocalRegistry() error {
	return exec.Command("docker", "rm", "-f", registryContainerName).Run()
}

func runActionsCommand(command, envFile string) error {
	vars, err := parseEnvFile(envFile)
	if err != nil {
		return err
	}

	if err = setupEnvVars(vars); err != nil {
		return err
	}
	defer removeEnvVars(vars)

	bin, err := getActionsBinaryPath()
	if err != nil {
		return err
	}

	cmd := exec.Command(bin, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getActionsBinaryPath() (string, error) {
	if path := os.Getenv("GITHUB_ACTIONS_BINARY"); path != "" {
		return path, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(wd, "../bin/github-actions"), nil
}
