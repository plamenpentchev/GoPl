package seven

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32
type database struct {
	db  map[string]dollars
	mux sync.RWMutex
}

//DB ...
var Db database

//Mux is aMultiplexer ...
var Mux *http.ServeMux

var mux = &sync.Mutex{}

func init() {
	Db = database{db: map[string]dollars{"shoes": 50, "socks": 5, "shirt": 25, "hat": 12}}

	Mux = http.NewServeMux()
	//... using our ServeMux(Mux) instance
	Mux.HandleFunc("/list", Db.list)
<<<<<<< HEAD
	Mux.HandleFunc("/", Db.list)

	//... using the global ServeMux(from net/http) instance.
	http.HandleFunc("/price", Db.price)
	http.HandleFunc("/priceProd", Db.price)
=======
	Mux.HandleFunc("/price", Db.price)
	Mux.HandleFunc("/create", Db.credate)
	Mux.HandleFunc("/update", Db.credate)
	Mux.HandleFunc("/read", Db.read)
	Mux.HandleFunc("/delete", Db.delete)
}

func (db *database) credate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if price, err := strconv.ParseFloat(query.Get("price"), 64); err == nil {
		dlrs := dollars(price)
		db.mux.Lock()
		db.db[query.Get("item")] = dlrs
		db.mux.Unlock()
		w.WriteHeader(http.StatusOK)
		db.PrintDatabase(w)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unacceptable price [%s]", query.Get("price"))
	}
}

func (db *database) read(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	db.mux.RLock()
	defer db.mux.RUnlock()
	price, ok := db.db[query.Get("item")]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s: %s", query.Get("item"), price)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no [%s] in the data base", query.Get("item"))
	}
}

func (db *database) delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_, ok := db.db[query.Get("item")]
	if ok {

		db.mux.Lock()
		delete(db.db, query.Get("item"))
		db.mux.Unlock()
		db.PrintDatabase(w)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no product [%s] in data base", query.Get("item"))
	}
>>>>>>> 32f0b8f1d0ea9940eb01d7bd63115e61d96213d7

}

func (db *database) list(w http.ResponseWriter, r *http.Request) {
	db.PrintDatabase(w)
}

func (db *database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db.db[item]
	if !ok {
		msg := fmt.Sprintf("product '%s' was not found", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s: %s", item, price)
}

//ServeHTTP satisfies http.Handler Interface
func (db *database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/list":
		db.PrintDatabase(w)
		break
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db.db[item]
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

func (db *database) PrintDatabase(w io.Writer) {
	for item, price := range db.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("€ %.2f", float32(d))
}
