package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const defaultErrorCode = 1

func formEnvString(env Environment) []string {
	envVarsCmd := make([]string, 0)
	envVars := os.Environ()

	for _, envVar := range envVars {
		envVarKey := strings.Split(envVar, "=")[0]

		v, ok := env[envVarKey]
		if !ok {
			envVarsCmd = append(envVarsCmd, envVar)
		} else {
			if !v.NeedRemove {
				envVarsCmd = append(envVarsCmd, fmt.Sprintf("%s=%s", envVarKey, v.Value))
			}
			delete(env, envVarKey)
		}
	}

	for k, v := range env {
		if !v.NeedRemove {
			envVarsCmd = append(envVarsCmd, fmt.Sprintf("%s=%s", k, v.Value))
		}
	}

	return envVarsCmd
}

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	//nolint:gosec
	command := exec.Command(cmd[0], cmd[1:]...)

	stdout := os.Stdout
	stderr := os.Stderr

	command.Stdout = stdout
	command.Stderr = stderr

	command.Env = formEnvString(env)

	if err := command.Run(); err != nil {
		//nolint:errorlint // Errors.As requires non-nil pointer which is not possible here
		if exitError, ok := err.(*exec.ExitError); ok {
			returnCode = exitError.ExitCode()
		} else {
			returnCode = defaultErrorCode
		}
	}

	return
}
