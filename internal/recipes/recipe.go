package recipes

import (
    "fmt"
)


type Recipe struct {
    ID string
    Name string
    
    // production output
    Output Product
    Output_amount int

    // production input
//    Input []Product
//    Input_amounts map[string]int32
}


func (r Recipe) Repr() string {
    //string input_str = '';
    //for _, e := range r.input {
    //    input_str += e + "\n"
    //}

    return fmt.Sprintf("Recipe: %s: %s out: %s amount: %d",
        r.ID, r.Name, r.Output.Repr(), r.Output_amount)

}
