package reportman

import (
    "fmt"
    "testing"
)

func TestExecute(t *testing.T) {
    r := NewRequest(make(chan *Response), "GOOG", nil)
    resp := r.Execute()
    for i := 0; i < 3; i++ {
        fmt.Println(resp.Stock[i])    
    }
}
