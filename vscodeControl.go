package main

import "os/exec"

func openVSCodeAndRunCommand(project Project) error {
    script := `tell application "Visual Studio Code"
        activate
        open "` + project.DirectoryPath + `"
    end tell

    tell application "System Events"
        tell process "Code"
            delay 2  -- Wait for VS Code to focus
            ` + "keystroke \"`\" using {control down, shift down}" + `  -- Open terminal
            delay 2  -- Wait for terminal to open
            keystroke "` + project.InstallCommand + " && " + project.StartCommand + `"
            keystroke return  -- Press enter
        end tell
    end tell`

    cmd := exec.Command("osascript", "-e", script)
    return cmd.Run()
}

