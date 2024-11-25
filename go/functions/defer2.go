package main

import (
    "fmt"
    "os"
)

func doClose(handle *os.File, filename string) {
    fmt.Println("Closing file", filename)
    if handle != nil {
        handle.Close()
    }
}

func createFile(filename, content string) {
    // open or create a file
    handle, err := os.Create(filename)
    
    // ensure the file will close no matter what happens next
    defer doClose(handle, filename)

    if err != nil {
        fmt.Println(err)
        return
    }

    numWritten, writeErr := handle.WriteString(content)
    if writeErr != nil {
        fmt.Println(writeErr)
        return
    }

    // beware len() on strings in general (UTF encoding)...
    if numWritten != len(content) {
        fmt.Println("Incorrect number of bytes written!")
        return
    }
}

func main() {
    createFile("test.txt", "TESTING!")

    // here, nope.txt already exists and we don't have write permissions
    createFile("nope.txt", "THIS WON'T WORK!")
}
