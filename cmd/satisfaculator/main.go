package main

import (
    "fmt"
    "satisfaculator/internal/recipes"
)


func main()  {
    fmt.Println("test");

    var prod = recipes.Product{"p01", "screws"};
    fmt.Println(prod.Repr())

    var rec = recipes.Recipe{"r01", "screws recipe", prod, 10}
    fmt.Println(rec.Repr())
}
