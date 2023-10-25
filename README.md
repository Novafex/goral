# Goral

Go tool for generating CRUD back-end services using a unified manifest file format.

![Gary the Goral](./docs/goral.png)

Goral is useful for quickly making, and maintaining simple CRUD style REST-is
API servers. These are services that generally Create-Retrieve-Update-Delete
resources connected to a database. Instead of the monotony of repeatedly making
struct-types and endpoints for the same patterns, let Goral do it instead.

Using a simple manifest declaring the data outline, it can generate all the
useful code for you.

* Automatic struct generation for resources
* Endpoint binding for web servers
* SQL generation scripts for building the schema
* Exporting TypeScript declarations
* And much more...

## Installation

Install using Go...

```bash
go install github.com/novafex/goral
```

## Usage

Goral is a CLI (command-line interface) which handles a lot of the work. Use the
command `goral` without any arguments, or `goral help` to see the guide.

To start using it in your Go project, navigate to your projects root folder,
where your `go.mod` file would be and execute the following:

```bash
goral init
# goral init -D
```

This will start the initalization process which will scaffold the needed files
and folders.

It will ask if you want to add a `.gitignore` rule to ignore generated code files,
it is your choice if you do or do not.

## See Also

- [CONTRIBUTING.md](./CONTRIBUTING.md) for community guidelines
- [STANDARDS.md](./STANDARDS.md) for coding and development standards
- [LICENSE](./LICENSE) for MIT license declaration
