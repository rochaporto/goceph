package rados

import ( 
  "fmt"
  "testing"
)

func TestVersion(t *testing.T) {
  fmt.Println("version: " + rados.Version())
}  
