package sqlbuilder

import "testing"

func TestInsert(t *testing.T) {
	cases := []struct {
		query     *InsertBuilder
		wantQuery string
		args      []interface{}
	}{
		{
			query:     Insert("products").Columns("title, description").Values("airpod", "airpod description"),
			wantQuery: "INSERT INTO products (title, description) VALUES ($1, $2)",
			args:      []interface{}{"airpod", "airpod description"},
		},
		{
			query: Insert("products").Columns("title, description").
				Values("airpod", "airpod description").
				Values("airpod 2", "airpod 2 description").
				Values("airpod 3", "airpod 3 description"),
			wantQuery: "INSERT INTO products (title, description) VALUES ($1, $2), ($3, $4), ($5, $6)",
			args:      []interface{}{"airpod", "airpod description", "airpod 2", "airpod 2 description", "airpod 3", "airpod 3 description"},
		},
	}

	for _, c := range cases {
		q, args := c.query.Query()
		if q != c.wantQuery || len(c.args) != len(args) {
			t.Errorf("want:\n %s\n got:\n %s\n", c.wantQuery, q)
		}
	}
}
