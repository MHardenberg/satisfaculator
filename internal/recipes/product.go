package main
import (
    "fmt"
)


type Product struct {
    name string
    id string
    recipes []Recipe
}


func (p Product) repr() string {
    return fmt.Sprintf("Prod: %s: %s", p.id, p.name)
}
