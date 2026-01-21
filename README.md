#Student Manage CLI (Go)

A Go-based CLI tool to manage students and grades.

## Features

- subcommands: `add`, `list`, `curve`
- Persistent JSON storage
- Table-drive unit tests

## Usage

```
bash
go run ./cmd/student-manager list
go run ./cmd/student-manager add --name=Alex --score=88
go run ./cmd/student-manager curve --points=5

```
