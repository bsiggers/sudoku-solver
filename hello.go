package main

import (
        "fmt"
        "io/ioutil"
        "path/filepath"
        "gopkg.in/yaml.v2"
)

type Sudoku struct {
  Board [][]int
}

func main() {
  filename, _ := filepath.Abs("./sudoku.yaml")
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  var sudoku Sudoku

  err = yaml.Unmarshal(yamlFile, &sudoku)
  if err != nil {
      panic(err)
  }

  fmt.Printf("Hello, test world\n")
  fmt.Printf("%#v\n", sudoku.Board)
}
