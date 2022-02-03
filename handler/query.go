package handler

import (
	"encoding/json"
	"github.com/graph-gophers/graphql-go"
	"github.com/hasura_auth/util"
	"net/http"
)

func Handle(schema *graphql.Schema) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(request.Body).Decode(&params); err != nil {
			util.ErrJsonResponse(writer, err)
			return
		}
		response := schema.Exec(request.Context(), params.Query, params.OperationName, params.Variables)
		util.JsonResponse(writer, response)
	})
}
