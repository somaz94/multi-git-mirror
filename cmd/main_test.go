package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// isolateGitConfigEnv starts the test with no GIT_CONFIG_COUNT/KEY_<n>/VALUE_<n>
// set; t.Setenv restores the originals and the cleanup drops entries run() added.
func isolateGitConfigEnv(t *testing.T) {
	t.Helper()
	for _, name := range append(gitConfigEntryNames(), "GIT_CONFIG_COUNT") {
		t.Setenv(name, "")
		os.Unsetenv(name)
	}
	t.Cleanup(func() {
		for _, name := range gitConfigEntryNames() {
			os.Unsetenv(name)
		}
	})
}

func gitConfigEntryNames() []string {
	var names []string
	for _, kv := range os.Environ() {
		name, _, _ := strings.Cut(kv, "=")
		if strings.HasPrefix(name, "GIT_CONFIG_KEY_") || strings.HasPrefix(name, "GIT_CONFIG_VALUE_") {
			names = append(names, name)
		}
	}
	return names
}

func TestRunInjectsSafeDirectoryEnv(t *testing.T) {
	tests := []struct {
		name      string
		workspace string
		want      []string
	}{
		{
			name:      "workspace set",
			workspace: "/runner/work/repo",
			want:      []string{"safe.directory=/runner/work/repo", "safe.directory=/github/workspace"},
		},
		{
			name: "workspace unset",
			want: []string{"safe.directory=/github/workspace"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateGitConfigEnv(t)
			t.Setenv("GITHUB_WORKSPACE", tt.workspace)
			// No targets: run() stops at config loading, after the env is set.
			t.Setenv("INPUT_TARGETS", "")

			if err := run(); err == nil || !strings.Contains(err.Error(), "failed to load config") {
				t.Fatalf("expected config load error, got %v", err)
			}

			n, err := strconv.Atoi(os.Getenv("GIT_CONFIG_COUNT"))
			if err != nil {
				t.Fatalf("GIT_CONFIG_COUNT: %v", err)
			}
			var got []string
			for i := range n {
				got = append(got, os.Getenv(fmt.Sprintf("GIT_CONFIG_KEY_%d", i))+"="+os.Getenv(fmt.Sprintf("GIT_CONFIG_VALUE_%d", i)))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected entries %v, got %v", tt.want, got)
			}
		})
	}
}
