package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// define a SnippetModel tpye which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// values for the placeholder parameters: title, content and expiry in
	// that order. This method returns a sql.Result type, which contains some
	// basic information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		fmt.Println("Error executing insert statement:", err)
		return 0, err
	}
	// Use the LastInsertId() method on the result to get the ID of our // newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		return 0, err
	}
	// The ID returned has the type int64, so we convert it to an int type // before returning.
	fmt.Println("Insert successful, last insert ID:", id)
	return int(id), nil
}

// get snippet by id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// get latest 10 snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
