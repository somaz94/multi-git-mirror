package mirror

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultSSHDir = "/root/.ssh"
	sshKeyFile    = "mirror_key"
	sshConfigFile = "config"
	knownHosts    = "known_hosts"
)

// setupSSH configures SSH key authentication for git operations.
func (m *Mirror) setupSSH() error {
	if m.cfg.SSHPrivateKey == "" {
		return nil
	}

	m.logInfo("Configuring SSH key authentication...")

	// Create .ssh directory
	if err := os.MkdirAll(m.sshDir, 0700); err != nil {
		return fmt.Errorf("failed to create .ssh directory: %w", err)
	}

	// Write private key
	keyPath := filepath.Join(m.sshDir, sshKeyFile)
	if err := os.WriteFile(keyPath, []byte(m.cfg.SSHPrivateKey+"\n"), 0600); err != nil {
		return fmt.Errorf("failed to write SSH key: %w", err)
	}

	// Write SSH config to use the key and disable strict host checking
	sshConfig := fmt.Sprintf(`Host *
  IdentityFile %s
  StrictHostKeyChecking no
  UserKnownHostsFile %s
`, keyPath, filepath.Join(m.sshDir, knownHosts))

	configPath := filepath.Join(m.sshDir, sshConfigFile)
	if err := os.WriteFile(configPath, []byte(sshConfig), 0600); err != nil {
		return fmt.Errorf("failed to write SSH config: %w", err)
	}

	// Create empty known_hosts
	knownHostsPath := filepath.Join(m.sshDir, knownHosts)
	if err := os.WriteFile(knownHostsPath, []byte{}, 0600); err != nil {
		return fmt.Errorf("failed to create known_hosts: %w", err)
	}

	// Set GIT_SSH_COMMAND to use our config
	// Without this git would fall back to the ambient SSH credentials and mirror
	// under the wrong identity, so a failure has to stop the setup.
	if err := os.Setenv("GIT_SSH_COMMAND", fmt.Sprintf("ssh -F %s -o BatchMode=yes", configPath)); err != nil {
		return fmt.Errorf("failed to set GIT_SSH_COMMAND: %w", err)
	}

	m.logDebug("SSH key configured at %s", keyPath)
	return nil
}

// cleanupSSH removes SSH key files.
func (m *Mirror) cleanupSSH() {
	if m.cfg.SSHPrivateKey == "" {
		return
	}

	// Best-effort cleanup; the process is on its way out either way.
	_ = os.Remove(filepath.Join(m.sshDir, sshKeyFile))
	_ = os.Remove(filepath.Join(m.sshDir, sshConfigFile))
	_ = os.Remove(filepath.Join(m.sshDir, knownHosts))
	_ = os.Unsetenv("GIT_SSH_COMMAND")

	m.logDebug("SSH key files cleaned up")
}
