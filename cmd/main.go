package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func main() {
	fmt.Println("hello")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	result := graphql.Do(graphql.Params{
		Schema:        testutil.StarWarsSchema,
		RequestString: query,
	})
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
