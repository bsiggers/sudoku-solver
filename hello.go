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

var VALIDSUM = 1+2+3+4+5+6+7+8+9

func isvalid(s *Sudoku) bool {
    for i := 0; i < 9; i++ {
      colsum := 0
      rowsum := 0
      for j := 0; j < 9; j++ {
        rowsum += s.Board[i][j]
        colsum += s.Board[j][i]
      }
      fmt.Printf("Rowsum is %d, Colsum is %d\n",rowsum,colsum)
      if rowsum != VALIDSUM || colsum != VALIDSUM {
          return false
      }
    }
    return true
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
  if isvalid(&sudoku) {
    fmt.Printf("Sudoku is valid")
  } else {
    fmt.Printf("Sudoku is not valid")
  }
}
