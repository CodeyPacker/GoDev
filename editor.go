package main

import (
	"fmt"
	"os"
	"os/exec"
  "github.com/fatih/color"
)

func openEditor(project Project) error {
  var cmd *exec.Cmd

  fmt.Print("🚀 Launching ")
  yellow := color.New(color.FgYellow).PrintfFunc()
  yellow(project.Name)
  println()

  switch project.Editor {
  case "code":
    cmd = exec.Command("code", project.DirectoryPath)
  case "vim":
    cmd = exec.Command("vim", project.DirectoryPath)
  case "fleet":
    cmd = exec.Command("fleet", project.DirectoryPath)
  case "goland":
    cmd = exec.Command("goland", project.DirectoryPath)
  default:
    return fmt.Errorf("Unknown editor: %s", project.Editor)
  }

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}
