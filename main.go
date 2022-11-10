package main

// #cgo CXXFLAGS: -I./
// #cgo LDFLAGS: -L./ -lstdc++
// #include "libreria.hpp"
import "C"

func main() {
    C.funcion()
}