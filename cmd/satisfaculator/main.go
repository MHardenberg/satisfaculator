package main

import (
    "fmt"
    "github.com/MHardenberg/satisfaculator/internal/recipes"
)


func main()  {
    fmt.Println("test");

    var prod recipes.Product;
    prod.name = "screws";
    prod.id = "p01";

    fmt.Println(prod.repr())

    var rec recipes.Recipe;
    rec.id = "p01";
    rec.name = "hallo";
    rec.output = prod;

    fmt.Println(rec.repr())
}
