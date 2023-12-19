package main

import (
    "github.com/ktr0731/go-fuzzyfinder"
)

func selectProject(projects []Project) (Project, error) {
    index, err := fuzzyfinder.Find(
        projects,
        func(i int) string {
            return projects[i].Name  // Assuming each project has a Name field
        },
        fuzzyfinder.WithPromptString("Select project:"),
    )
    if err != nil {
        return Project{}, err
    }
    return projects[index], nil
}

