package db

type Database struct {
	Schema Schema
}

type Schema struct {
	Tables map[string]Table
}

type Table struct {
	Fields map[string]*Type
}

type Type string
