package schema

import (
	"encoding/json"

	"github.com/fadhilnurmahardi/simple-promotion/internal/global/helper"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/endpoint"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/go-kit/kit/log"
	"github.com/graphql-go/graphql"
)

func GenerateQuery(svc service.IService, logger log.Logger) *graphql.Object {
	end := endpoint.Calculate(svc)

	end = helper.LogRequest(logger)(end)
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"check_promotion": &graphql.Field{
					Type:        PromotionSchema,
					Description: "For check promotion",
					Args: graphql.FieldConfigArgument{
						"cart": &graphql.ArgumentConfig{
							Type: graphql.NewList(CheckPromotionInput),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						var request *model.RawPayload
						bytesCart, err := json.Marshal(p.Args)
						if err != nil {
							return nil, err
						}
						err = json.Unmarshal([]byte(bytesCart), &request)
						if err != nil {
							return nil, err
						}
						return end(p.Context, request.Cart)
					},
				},
			},
		},
	)
}
