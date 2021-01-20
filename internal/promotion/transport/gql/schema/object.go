package schema

import "github.com/graphql-go/graphql"

var PromotionSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "promotion",
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type: graphql.Float,
			},
			"total_after_discount": &graphql.Field{
				Type: graphql.Float,
			},
			"discount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var CheckPromotionInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CheckPromotionInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"sku": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"qty": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)
