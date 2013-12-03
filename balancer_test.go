package reportman

import (
    "fmt"
    "testing"
)

func makeTestRequest() (Request, chan Request) {
    work := make(chan Request)
    return *NewRequest(make(chan *Response), "AAPL", nil), work
}

func TestWorker(t *testing.T) {
    r, c := makeTestRequest()
    w := &Worker{c, 0}
    done := make(chan *Worker)
    go w.work(done)
    w.requests <- r
    resp := <- r.Response
    for i := 0; i < 3; i++ {
        fmt.Println(resp.Stock[i])    
    }
}
