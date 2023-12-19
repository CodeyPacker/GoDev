package main

import (
    "bufio"
    "os/exec"
    "strings"
)

func killProcessesOnPort(port string) error {
    // Finding process IDs using lsof
    cmd := exec.Command("lsof", "-t", "-i", "tcp:"+port)
    output, err := cmd.Output()
    if err != nil {
        return err
    }

    // Create a scanner to read the process IDs
    scanner := bufio.NewScanner(strings.NewReader(string(output)))
    for scanner.Scan() {
        pid := scanner.Text()

        // Killing each process found
        killCmd := exec.Command("kill", pid)
        if err := killCmd.Run(); err != nil {
            return err
        }
    }

    return scanner.Err()
}

