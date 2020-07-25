package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func StarWars(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	result := graphql.Do(graphql.Params{
		Schema:        testutil.StarWarsSchema,
		RequestString: query,
	})
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)

}
