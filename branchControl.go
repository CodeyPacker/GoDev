package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// createGitBranch creates a new Git branch for the specified project
func createGitBranch(project Project, branchName string) error {
    // Change to the project's directory
    err := os.Chdir(project.DirectoryPath)
    if err != nil {
        return fmt.Errorf("error changing to project directory: %v", err)
    }

    // Switch to the master branch
    if err := runGitCommand("checkout", "master"); err != nil {
        return err
    }

    // Pull the latest changes
    if err := runGitCommand("pull"); err != nil {
        return err
    }

    // Check if there are uncommitted changes
    if err := checkForUncommittedChanges(); err != nil {
        return err
    }

    // Create and checkout the new branch
    return runGitCommand("checkout", "-b", branchName)
}

func runGitCommand(args ...string) error {
    cmd := exec.Command("git", args...)
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("git %s failed: %s, %v", strings.Join(args, " "), string(output), err)
    }
    return nil
}

func checkForUncommittedChanges() error {
    cmd := exec.Command("git", "status", "--porcelain")
    output, err := cmd.Output()
    if err != nil {
        return err
    }

    if strings.TrimSpace(string(output)) != "" {
        fmt.Print("Uncommitted changes found. Do you want to stash them? (y/n): ")
        reader := bufio.NewReader(os.Stdin)
        response, _ := reader.ReadString('\n')
        if strings.TrimSpace(response) == "y" {
            return runGitCommand("stash")
        }
    }
    return nil
}
