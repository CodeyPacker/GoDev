package main

import (
	"os/exec"

	"github.com/fatih/color"
)

func openProjectURLs(project Project) error {
    openURL := func(url string) error {
        color.Green("Opening %s", url)
        cmd := exec.Command("open", url)
        return cmd.Run()
    }

    if err := openURL(project.LocalhostUrl); err != nil {
        // return err
    }
    if err := openURL(project.ContentfulUrl); err != nil {
        // return err
    }
    if err := openURL(project.RepoUrl); err != nil {
        // return err
    }

    return nil
}

func closeBrowserTabs() error {
    arcScript := `
    tell application "Arc"
        tell front window
            repeat with _tab in tabs
                set _url to get URL of _tab
                if _url contains "bitbucket" or _url contains "contentful" or _url contains "localhost" then
                    set _tab_id to get id of _tab
                    tell tab id _tab_id to close
                end if
            end repeat
        end tell
    end tell
    `

    err := exec.Command("osascript", "-e", arcScript).Run()
    if err != nil {
      color.Green("Closing old browser tabs")
        // return err
    }

    return nil
}

