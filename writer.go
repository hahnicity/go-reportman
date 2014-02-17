package reportman

import (
    "encoding/csv"
    "github.com/hahnicity/go-reportman/config"
    "os"
    "strconv"
)


func WriteToCsv(ar []*Response, dataFile string) {
    f, err := os.Create(dataFile)
    if err != nil { panic(err) }
    defer f.Close()
    w := csv.NewWriter(f)
    defer w.Flush()
    for _, resp := range ar {
        w.Write([]string{"Symbol", "Date", "Adj.Close"})
        for _, stock := range resp.Stock {
            var toWrite = []string {
                resp.Symbol,
                stock.Date,
                strconv.FormatFloat(stock.Adj, 'f', config.SigDigits, 64),
            }
            w.Write(toWrite)
        }
    }
}
