package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Float("price").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(10,2)",
				dialect.Postgres: "numeric",
			}),
		field.Int("stock").
			Default(0).
			Positive(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
