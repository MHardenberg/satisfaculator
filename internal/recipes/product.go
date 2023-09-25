package recipes

import (
    "fmt"
)


type Product struct {
    Name string
    ID string
//    Recipes []Recipe

}


func (p Product) Repr() string {
    return fmt.Sprintf("Prod: %s: %s", p.ID, p.Name)
}
