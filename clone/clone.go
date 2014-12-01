package clone

import (
    "log"
    "errors"
    "os"
)

func Exec(args []string) (bool, error){

    lenArgs := len(args)

    // accepts only the github url and the path to clone
    if lenArgs < 1 || lenArgs > 2 {
        help()
    }

    // the github repository url
    repository := args[0]

    // preparing the path to clone the project
    clonePath, err := getClonePath(args)

    return true, nil

}

func getClonePath(args []string) (string, error){
    return "Hello", nil
}

func help() {
    errors.New("Just showing a new error")
    log.Println("On erreor should present error message")
}
