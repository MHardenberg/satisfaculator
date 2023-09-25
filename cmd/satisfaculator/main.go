package main

import (
    "fmt"
)


//import "github.com/MHardenberg/satisfaculator/product"


func main()  {
    fmt.Println("test");

    var prod Product;
    prod.name = "screws";
    prod.id = "p01";

    fmt.Println(prod.repr())

    var rec Recipe;
    rec.id = "p01";
    rec.name = "hallo";
    rec.output = prod;

    fmt.Println(rec.repr())
}
