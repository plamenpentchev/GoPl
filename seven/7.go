package seven

import (
	"fmt"
	"io"
	"net/http"
)

type dollars float32
type database map[string]dollars

//DB ...
var Db database

func init() {
	Db = database{"shoes": 50, "socks": 5, "shirt": 25}
}

//ServeHTTP satisfies http.Handler Interface
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/list":
		db.PrintDatabase(w)
		break
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if ok {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s: %s\n", item, price)
		} else {
			// w.WriteHeader(http.StatusNotFound)
			msg := fmt.Sprintf("product [%s] not found", item)
			http.Error(w, msg, http.StatusNotFound)

		}
		break
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "undefined path [%s]", r.URL.Path)
	}

}

func (db database) PrintDatabase(w io.Writer) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("€ %.2f", float32(d))
}
