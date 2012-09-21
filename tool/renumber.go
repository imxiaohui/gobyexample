// Renumber the .go content files according to the
// order in index.txt.

package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "regexp"
)

func main() {
    // read names of source files
    sourceNames := make([]string, 0)
    sourceMap := make(map[string]string)
    fileInfos, dirErr := ioutil.ReadDir("./")
    if dirErr != nil { panic(dirErr) }
    baseTrimmer, _ := regexp.Compile("([0-9x]+-)|(.go)")
    for _, fi := range fileInfos {
        baseName := baseTrimmer.ReplaceAllString(fi.Name(), "")
        if baseName != ".git" && baseName != "tool" && baseName != "README.md" {
            sourceNames = append(sourceNames, baseName)
            sourceMap[baseName] = fi.Name()
        }
    }

    // read names from index
    indexBytes, idxErr := ioutil.ReadFile("tool/index.txt")
    if idxErr != nil { panic (idxErr) }
    indexNames := strings.Split(string(indexBytes), "\n")
    
    // sanity check two lists
    if len(sourceNames) != len(indexNames) {
        panic("mismatched names")
    }
    
    // rename some stuff
    for index, indexName := range indexNames {
        oldName := sourceMap[indexName]
        newName := fmt.Sprintf("%d-%s.go", index+1, indexName)
        os.Rename(oldName, newName)
    }
}
