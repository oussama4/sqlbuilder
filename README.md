# A simple package for building sql queries in Go.


[![Go Reference](https://pkg.go.dev/badge/github.com/oussama4/sqlbuilder.svg)](https://pkg.go.dev/github.com/oussama4/sqlbuilder)


- [Usage](#usage)
   - [Basic queries](#basic-queries)
	 - [Select](#select)
	 - [Update](#update)
	 - [Insert](#insert)
	 - [Delete](#delete)


## Usage
### Basic queries
#### Select
```go
// query: SELECT * FROM products
query, _ := sqlbuilder.Select("*").From("products").
		Query()

// query: SELECT id, title FROM products
query, _ := sqlbuilder.Select("id", "title").From("products").
		Query()
```

#### Update
```go
// query: UPDATE products SET title = $1
// args: []interface{}{"new title"}
query, args := sqlbuilder.Update("products").
		Set("title", "new title").
		Query()
```

#### Insert
```go
// query: INSERT INTO products (title, description)
// 		  VALUES ($1, $2)
// args: []interface{}{"new title", "new description"}
query, args := sqlbuilder.Insert("products").
		Columns("title", "description").
		Values("new title", "new desscription").
		Query()
```

#### Delete
```go
// query: DELETE FROM products
query, _ := sqlbuilder.DeleteFrom("products").Query()
```