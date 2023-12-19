package main

import (
  "encoding/json"
  "fmt" 
  "io/ioutil"
)

type Project struct {
    Name           string `json:"name"`
    DirectoryPath  string `json:"directoryPath"`
    InstallCommand string `json:"installCommand"`
    StartCommand   string `json:"startCommand"`
    LocalhostUrl   string `json:"localhostUrl"`
    ContentfulUrl  string `json:"contentfulUrl"`
    RepoUrl        string `json:"repoUrl"`
    Editor         string `json:"editor"`
}

type Configuration struct {
    Projects []Project `json:"projects"`
}

func main() {
  configFile, err := ioutil.ReadFile("config.json")
  if err != nil {
    fmt.Println("Error reading config file:", err)
    return
  }

  var config Configuration
  err = json.Unmarshal(configFile, &config)
  if err != nil {
    fmt.Println("Error parsing config file:", err)
    return
  }

  selectedProject, err := selectProject(config.Projects)
  if err != nil {
    fmt.Println("Project selection canceled:", err)
    return
  }

  err = openVSCodeAndRunCommand(selectedProject)
  if err != nil {
    fmt.Println("Error opening VS Code:", err)
  }

  err = openEditor(selectedProject)
  if err != nil {
    fmt.Println("Error opening editor:", err)
  }

  err = closeBrowserTabs()
  // if err != nil {
  //   fmt.Println("Error closing browser tabs:", err)
  // }

  err = openProjectURLs(selectedProject)
  // if err != nil {
  //   fmt.Println("Error opening project URLs:", err)
  

  if err := killProcessesOnPort("8000"); err != nil {
    // fmt.Println("Error killing processes on port 8000:", err)
  }
}
