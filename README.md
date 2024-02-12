# Goral

Go tool for generating CRUD back-ends services using a unified manifest file format.

![Chinese Goral](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8d/Nemorhaeduscaudatusarnouxianus.JPG/1280px-Nemorhaeduscaudatusarnouxianus.JPG)

## Repository Rules

- `main` branch is reserved for production facing code.
- `beta` branch holds release-candidate next versions.
- `develop` is for current development usage.
- Use slash-elimited prefixes for specific tasks based on their type:
  - `issue/*` using the issue ID for tasks in the Issues page
  - `feature/*` with a name for large feature branches
  - `task/*` for mundane tasks like linting or small refactors

Work in the develop and prefixed branches only. When we are ready, we will merge into `beta` and then `main` with their respective PRs and tags.

go get github.com/golang/mock/mockgen/model
