package e2e

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

const (
	registryContainerName = "github-actions-registry"
	githubActionsImage    = "github-actions-e2e"
)

func readEnvFile(envFile string) ([]string, error) {
	var vars []string

	file, err := os.Open(envFile)
	if err != nil {
		return vars, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vars = append(vars, scanner.Text())
	}

	path := os.Getenv("PATH")
	vars = append(vars, fmt.Sprintf("PATH=%s", path))
	return vars, scanner.Err()
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
	vars, err := readEnvFile(envFile)
	if err != nil {
		return err
	}

	bin, err := getActionsBinaryPath()
	if err != nil {
		return err
	}

	cmd := exec.Command(bin, command)
	cmd.Env = vars
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getActionsBinaryPath() (string, error) {
	if path := os.Getenv("GITHUB_ACTIONS_BINARY"); path != "" {
		return path, nil
	}

	return "../bin/github-actions", nil
}

func inspectImage(image string) (inspectResult, error) {
	out, err := exec.Command("docker", "inspect", image).Output()
	if err != nil {
		return inspectResult{}, err
	}
	var result []inspectResult
	if err = json.Unmarshal(out, &result); err != nil {
		return inspectResult{}, err
	}
	return result[0], nil
}

type inspectResult struct {
	RepoTags []string            `json:"RepoTags"`
	Config   inspectResultConfig `json:"Config"`
}

type inspectResultConfig struct {
	Labels map[string]string `json:"Labels"`
}
