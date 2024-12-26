package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Checkout switches to the specified branch.
func Checkout(repoPath string, branchName string) error {
	// Ensure the repository exists
	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return fmt.Errorf("not a Git repository: %s", repoPath)
	}

	// Check if the branch exists
	branchPath := filepath.Join(gitDir, "refs", "heads", branchName)
	if _, err := os.Stat(branchPath); os.IsNotExist(err) {
		return fmt.Errorf("branch %s does not exist", branchName)
	}

	// Read the commit hash from the branch reference
	commitHash, err := os.ReadFile(branchPath)
	if err != nil {
		return fmt.Errorf("failed to read branch reference %s: %w", branchName, err)
	}

	// Update the HEAD to point to the new branch
	headPath := filepath.Join(gitDir, "HEAD")
	err = os.WriteFile(headPath, []byte(fmt.Sprintf("ref: refs/heads/%s", branchName)), 0644)
	if err != nil {
		return fmt.Errorf("failed to update HEAD: %w", err)
	}

	// Optionally, print the commit hash to confirm
	fmt.Printf("Switched to branch %s (commit: %s)\n", branchName, strings.TrimSpace(string(commitHash)))
	return nil
}
