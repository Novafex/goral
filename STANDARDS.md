# Development Standards

This document outlines some of the coding and development standards of this
repository and project. It needs to be adhered to for all branches and
modifications.

## Repository Rules

- `main` branch is reserved for production facing code.
- `beta` branch holds release-candidate next versions.
- `develop` is for current development usage.
- Use slash-elimited prefixes for specific tasks based on their type:
  - `issue/*` using the issue ID for tasks in the Issues page
  - `feature/*` with a name for large feature branches
  - `task/*` for mundane tasks like linting or small refactors

Work in the develop and prefixed branches only. When we are ready, we will merge
into `beta` and then `main` with their respective PRs and tags.

## Formatting

This is a Go based project, as such standard `gofmt` formatting will be applied
to those standard files.

Otherwise some general standards of code style are:

- Space indentation on commiting, using 4 spaces per tab.
- UTF-8 character set encoding
- LF line breaks
- Braces on same line as expression
