# simple-server

## Introduction

simple-server is a go package designed to allow easily creating a server with template files, include javascript and bootstrap.

## Usage

1. Clone package with git clone.
2. `go run web/server/server.go`
3. Navigate to `localhost:9000`

## Making Changes

### Easy Add

1. Add a template file to `views`, e.g. `filename.html`.
2. Navigate to `localhost:9000/d/filename` or `localhost:9000/d/filename.html`.

### Alternate

1. Follow practice for `Homepage`
2. Execute template and pass in variables

### Headers

Headers may be edited in `partials/_header.html`

### Scripts

Scripts may be added by creating `{{define "scripts"}}{{end}}`
