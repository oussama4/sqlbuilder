package sqlbuilder

import "testing"

func TestDelete(t *testing.T) {
	cases := []struct {
		query     *DeleteBuilder
		wantQuery string
		args      []interface{}
	}{
		{
			query:     DeleteFrom("products"),
			wantQuery: "DELETE FROM products",
		},
		{
			query:     DeleteFrom("products").Where(Like("title", "%airpod%")),
			wantQuery: "DELETE FROM products WHERE title LIKE $1",
			args:      []interface{}{"airpod"},
		},
	}

	for _, c := range cases {
		q, args := c.query.Query()
		if q != c.wantQuery || len(c.args) != len(args) {
			t.Errorf("want:\n %s\n got:\n %s\n", c.wantQuery, q)
		}
	}
}
