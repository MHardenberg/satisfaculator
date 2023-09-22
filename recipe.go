package main
import (
    "fmt"
)


type Recipe struct {
    id string
    name string
    
    // production output
    output Product
    output_amount int32

    // production input
    input []Product
    input_amounts map[string]int32
}


func (r Recipe) repr() string {
    //string input_str = '';
    //for _, e := range r.input {
    //    input_str += e + "\n"
    //}

    return fmt.Sprintf("Recipe: %s: %s out: %s",
        r.id, r.name, r.output.repr())
}
