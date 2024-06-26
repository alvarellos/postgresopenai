package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/tmc/langchaingo/tools/sqldatabase/postgresql"
)

func TestMakeSample(t *testing.T) {
	// Create a temporary database connection string
	dsn := "postgres://postgres:postgres@localhost:5432/postgres"

	// Call the makeSample function
	makeSample(dsn)

	// Check that the tables were created and populated correctly
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Errorf("Error opening database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM foo")
	if err != nil {
		t.Errorf("Error querying table 'foo': %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			t.Errorf("Error scanning row: %v", err)
		}
		if id < 0 || id >= 100 {
			t.Errorf("Invalid ID: %d", id)
		}
		if name != fmt.Sprintf("Foo %03d", id) {
			t.Errorf("Invalid name: %s", name)
		}
	}

	rows, err = db.Query("SELECT * FROM foo1")
	if err != nil {
		t.Errorf("Error querying table 'foo1': %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			t.Errorf("Error scanning row: %v", err)
		}
		if id < 0 || id >= 200 {
			t.Errorf("Invalid ID: %d", id)
		}
		if name != fmt.Sprintf("Foo1 %03d", id) {
			t.Errorf("Invalid name: %s", name)
		}
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_makeSample(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeSample(tt.args.dsn)
		})
	}
}

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
