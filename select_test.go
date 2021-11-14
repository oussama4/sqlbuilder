package sqlbuilder

import "testing"

func TestSelect(t *testing.T) {
	cases := []struct {
		query     *SelectBuilder
		wantQuery string
		args      []interface{}
	}{
		{
			query:     Select("*").From("products"),
			wantQuery: "SELECT * FROM products",
		},
		{
			query:     Select("id", "title", "description").From("products").Where(Eq("title", "airpod")),
			wantQuery: "SELECT id, title, description FROM products WHERE title = $1",
			args:      []interface{}{"airpod"},
		},
		{
			query:     Select("title").From("products").OrderBy("title").Limit(10).Offset(20),
			wantQuery: "SELECT title FROM products ORDER BY title LIMIT 10 OFFSET 20",
		},
		{
			query: Select("*").From("product_variants").Where(
				And(
					Or(
						Gte("price", 10.00),
						Eq("weight", 0),
					),
					Lt("quantity", 10),
				),
			),
			wantQuery: "SELECT * FROM product_variants WHERE ((price >= $1) OR (weight = $2)) AND (quantity < $3)",
			args:      []interface{}{10.00, 0, 10},
		},
		{
			query: Select("product_id", "SUM(price)").From("product_variants").
				GroupBy("product_id").
				Having(Neq("weight", 0)),
			wantQuery: "SELECT product_id, SUM(price) FROM product_variants GROUP BY product_id HAVING weight <> $1",
			args:      []interface{}{0},
		},
		{
			query: Select("p.id", "v.sku", "v.price").From("products p").
				Join("product_variants v", Expr("p.id = v.product_id")),
			wantQuery: "SELECT p.id, v.sku, v.price FROM products p INNER JOIN product_variants v ON p.id = v.product_id",
		},
	}

	for _, c := range cases {
		q, args := c.query.Query()
		if q != c.wantQuery || len(c.args) != len(args) {
			t.Errorf("want:\n %s\n got:\n %s\n", c.wantQuery, q)
		}
	}
}
