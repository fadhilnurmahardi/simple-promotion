package schema

import (
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/go-kit/kit/log"
	"github.com/graphql-go/graphql"
)

func GenerateSchema(svc service.IService, logger log.Logger) graphql.Schema {
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: GenerateQuery(svc, logger),
		},
	)
	if err != nil {
		panic(err)
	}
	return schema
}
