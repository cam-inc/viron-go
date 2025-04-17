# Contributing to Viron GO

---

We welcome and support contributions to this project.

Tasks are managed in the [GitHub Project](https://github.com/orgs/cam-inc/projects/2), where you can check ongoing, completed, and high-priority issues.

## How can you contribute?

### Reporting Bugs

To report a bug, please submit a [GitHub issue](https://github.com/cam-inc/viron-go/issues/new?assignees=&labels=bug&template=bug_report.md&title=). Before submitting, make sure a similar issue does not already exist.

### Proposing Feature Enhancements

To propose a feature enhancement, please submit a [GitHub issue](https://github.com/cam-inc/viron-go/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=).

### Contributing Code

#### Getting the Code

```bash
git clone git@github.com:cam-inc/viron-go.git
```

#### Setting Up Tools

##### Installing asdf

```bash
# Hint: Other installation methods are also available.
# See https://asdf-vm.com/guide/getting-started.html for details.
$ brew install asdf

# Configuration
$ echo 'export PATH="${ASDF_DATA_DIR:-$HOME/.asdf}/shims:$PATH"' >> ~/.zshrc
$ source ~/.zshrc
```

##### Installing asdf Plugins

```bash
asdf plugin add golang
asdf plugin add task https://github.com/paulvollmer/asdf-task.git
asdf plugin add lefthook https://github.com/jtzero/asdf-lefthook.git
```

##### Installing Dependencies with asdf

```bash
asdf install
```

#### Setting Up the Application

##### Installing Dependencies with task

```bash
task install
```

##### Running go mod tidy

```bash
task tidy
```

#### Running Tests

```bash
task test
```

Here is a simple guide to contribute code:

1. Fork the repository and clone it to your local machine.
2. Create a new branch with a meaningful name suitable for the task from the `main` branch.
3. Run the following command for setup: `go mod tidy`.
4. Push the branch.
5. Submit a pull request to the upstream repository.

#### Using the Library in Other Projects

To use this library in other Go projects, use `go get`.

```bash
Example: go get github.com/cam-inc/viron-go@v1.0.0

Import it as usual:

import "github.com/cam-inc/viron-go/lib/domains"
```

#### Release Workflow

[Module Release and Versioning Workflow](https://go.dev/doc/modules/release-workflow)

The release workflow for custom Go (Golang) libraries (assuming GitHub integration) can be designed as follows. This ensures that versioned code is correctly released and available for external projects using Go Modules.

**🔁 General Release Workflow**

1. Implement and test features on a development branch.
2. Merge into the `main` branch.
3. A new version is assigned automatically upon merging into `main`.
4. Use `go get` in other projects to import it.

**📌 Detailed Steps**

- ① Development and Testing
  - When adding features or fixing bugs, work on a branch (e.g., feature/xxx, fix/yyy), and validate with unit tests or CI.

- ② Merging into the `main` Branch
  - After review (e.g., via pull request), merge into the `main` branch to complete the milestone.

- ③ Versioning
  - A new version is automatically assigned upon merging into `main`.
  - Determined from the PR commit message:
    - If it includes `BREAKING CHANGE` → Major update (v1.x.x → v2.0.0)
    - If it includes `feat` → Minor update (v1.1.0 → v1.2.0)
    - If it includes `fix` → Patch update (v1.2.3 → v1.2.4)
    - Otherwise → Patch update (v1.2.3 → v1.2.4)

## Code of Conduct

Please adhere to [this document](./CODE_OF_CONDUCT.md).

## License

By contributing to this project, you agree that your contributions will be licensed under the [MIT License](./LICENSE).

## Contributors

We appreciate everyone who has contributed to this project.

<table>
  <tr>
    <td align="center"><a href="https://github.com/takoring"><img src="https://avatars.githubusercontent.com/u/24517668?v=4" width="100px;" alt=""/><br /><sub><b>takoring</b></sub></a><br />💻</td>
  </tr>
</table>
