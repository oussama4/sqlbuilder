package sqlbuilder

import "testing"

func TestUpdate(t *testing.T) {
	cases := []struct {
		query     *UpdateBuilder
		wantQuery string
		args      []interface{}
	}{
		{
			query:     Update("products").Set("title", "airpod up").Set("description", "descr"),
			wantQuery: "UPDATE products SET title = $1, description = $2",
			args:      []interface{}{"airpod up", "descr"},
		},
		{
			query: Update("product_variants").Set("price", 10.00).
				Where(
					And(
						Or(
							Gte("price", 10.00),
							In("product_id", 1, 2, 3),
						),
						Lt("quantity", 10),
					),
				),
			wantQuery: "UPDATE product_variants SET price = $1 WHERE ((price >= $2) OR (product_id IN ($3, $4, $5))) AND (quantity < $6)",
			args:      []interface{}{10.00, 10.00, 1, 2, 3, 10},
		},
	}

	for _, c := range cases {
		q, args := c.query.Query()
		if q != c.wantQuery || len(c.args) != len(args) {
			t.Errorf("want:\n %s\n got:\n %s\n", c.wantQuery, q)
		}
	}
}
