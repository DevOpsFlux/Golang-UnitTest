package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
// CASE 0
// http://127.0.0.1:8080/handler

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UnitTest Handler - Hello %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/handler", handler)
	http.ListenAndServe(":8080", nil)
}

*/

/*
// CASE 1
// 1. No request parsing
// http://127.0.0.1:8080/
// http://127.0.0.1:8080/abc
// http://127.0.0.1:8080/xyz

type pounds float32

func (p pounds) String() string {
	return fmt.Sprintf("£%.2f", p)
}

type database map[string]pounds

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{
		"foo": 1,
		"bar": 2,
	}

	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
*/

/*
//2. Manual request parsing
// http://127.0.0.1:8080/
// http://127.0.0.1:8080/abc
// http://127.0.0.1:8080/xyz
type pounds float32

func (p pounds) String() string {
	return fmt.Sprintf("£%.2f", p)
}

type database map[string]pounds

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/foo":
		fmt.Fprintf(w, "foo: %s\n", d["foo"])
	case "/bar":
		fmt.Fprintf(w, "bar: %s\n", d["bar"])
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No page found for: %s\n", r.URL)
	}
}

func main() {
	db := database{
		"foo": 1,
		"bar": 2,
	}

	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
*/

/*
// 3. Multiplexer
// http://127.0.0.1:8080/
// http://127.0.0.1:8080/foo
// http://127.0.0.1:8080/bar
type pounds float32

func (p pounds) String() string {
	return fmt.Sprintf("£%.2f", p)
}

type database map[string]pounds

func (d database) foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo: %s\n", d["foo"])
}

func (d database) bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar: %s\n", d["bar"])
}

func (d database) baz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "baz: %s\n", d["baz"])
}

func main() {
	db := database{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	mux := http.NewServeMux()

	mux.Handle("/foo", http.HandlerFunc(db.foo))
	mux.Handle("/bar", http.HandlerFunc(db.bar))

	// Convenience method for longer form mux.Handle
	mux.HandleFunc("/baz", db.baz)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
*/

// 4. Global multiplexer
type pounds float32

func (p pounds) String() string {
	return fmt.Sprintf("£%.2f", p)
}

type database map[string]pounds

func (d database) foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo: %s\n", d["foo"])
}

func (d database) bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar: %s\n", d["bar"])
}

func (d database) baz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "baz: %s\n", d["baz"])
}

func main() {
	db := database{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	http.HandleFunc("/foo", db.foo)
	http.HandleFunc("/bar", db.bar)
	http.HandleFunc("/baz", db.baz)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// https://www.integralist.co.uk/posts/understanding-golangs-func-type/
// 번역 : https://dejavuqa.tistory.com/314?category=318996
