package main

import (
	"fmt"
	"os"

	"github.com/somaz94/multi-git-mirror/internal/config"
	"github.com/somaz94/multi-git-mirror/internal/mirror"
	"github.com/somaz94/multi-git-mirror/internal/output"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "::error::%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Trust the GitHub Actions workspace directory inside Docker, via env so no config file is written.
	if workspace := os.Getenv("GITHUB_WORKSPACE"); workspace != "" {
		// Best effort: if this fails, the git commands below report it themselves.
		_ = mirror.AddConfigEnv("safe.directory", workspace)
	}
	_ = mirror.AddConfigEnv("safe.directory", "/github/workspace")

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	m := mirror.New(cfg)
	results := m.Run()

	return output.Write(results)
}
