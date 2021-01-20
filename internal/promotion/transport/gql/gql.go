package gql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fadhilnurmahardi/simple-promotion/internal/global/model"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/transport/gql/schema"
	"github.com/go-kit/kit/log"
	"github.com/graphql-go/graphql"
)

// MakeHandler ...
func MakeHandler(svc service.IService, logger log.Logger) http.HandlerFunc {
	promotionSchema := schema.GenerateSchema(svc, logger)
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Log("err", fmt.Sprintf("Error reading body: %v", err))
			http.Error(w, "can't read body", http.StatusUnprocessableEntity)
			return
		}
		var request *model.Query
		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "can't read body", http.StatusUnprocessableEntity)
		}
		logger.Log("params", request.Query)
		result := graphql.Do(graphql.Params{
			Schema:        promotionSchema,
			RequestString: request.Query,
		})
		if len(result.Errors) > 0 {
			logger.Log("err", fmt.Sprintf("%v", result.Errors))
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(result)
	}
}
