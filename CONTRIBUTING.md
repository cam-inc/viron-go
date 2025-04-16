# Contributing to Viron GO

---

We encourage and support everyone to contribute to this project.

Tasks are managed in the [GitHub Project](https://github.com/orgs/cam-inc/projects/2), where you can check ongoing, completed, and high-priority issues.

## How can you contribute?

### Reporting Bugs

To report a bug, please submit a [GitHub issue](https://github.com/cam-inc/viron-go/issues/new?assignees=&labels=bug&template=bug_report.md&title=). Before submitting, make sure a similar issue does not already exist.

### Proposing Feature Enhancements

To propose a feature enhancement, please submit a [GitHub issue](https://github.com/cam-inc/viron-go/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=).

### Contributing Code

#### Getting the Code
```
$ git clone git@github.com:cam-inc/fensi-go.git
```

#### Setting Up Tools
##### Installing taskfile
```
# Tips: Other installation methods are also acceptable.
# For more details, refer to https://taskfile.dev/installation/.

$ brew install go-task/tap/go-task
```

##### Installing asdf
```
# Tips: Other installation methods are also acceptable.
# For more details, refer to https://asdf-vm.com/guide/getting-started.html
$ brew install asdf

# Configuration
$ echo 'export PATH="${ASDF_DATA_DIR:-$HOME/.asdf}/shims:$PATH"' >> ~/.zshrc
$ source ~/.zshrc
```

##### Installing asdf plugins
```
$ asdf plugin add golang
$ asdf plugin add ko https://github.com/zasdaym/asdf-ko.git
$ asdf plugin add lefthook https://github.com/jtzero/asdf-lefthook.git
$ asdf plugin add helm https://github.com/Antiarchitect/asdf-helm.git
```

##### Installing dependencies with asdf
```
$ asdf install
```

##### Installing lefthook
```
$ brew install lefthook
```

##### Installing dependencies with asdf
```
$ npx lefthook install
```

#### Setting Up the Application
##### Installing dependencies with task
```
$ task install
```

##### Running go mod tidy
```
$ task tidy
```

#### Running Tests
```
$ task test
```

Below is a simple guide for contributing code:

1. Fork the repository and clone it to your local machine.
2. Create a new branch with a meaningful name suitable for the task from the `main` branch.
3. Run the following command to set up: `go mod tidy`.
4. Push the branch.
5. Submit a pull request to the upstream repository.

#### Managing Versions
To make the version recognizable by Go Modules, you need to tag it in Git. Always use the format vX.Y.Z ([Semantic Versioning](https://semver.org)).

[Module version numbering](https://go.dev/doc/modules/version-numbers)
```
Example:

To tag v1.0.0:

git tag v1.0.0
git push origin v1.0.0
```

📌 Note

The tag name must include the `v` prefix (e.g., v1.2.3). Go Modules detects this and uses the version as a module.

#### Using the Library in Other Projects
To use this library in another Go project, use `go get`.
```
Example: go get github.com/cam-inc/viron-go@v1.0.0

Import it as usual:

import "github.com/cam-inc/viron-go/lib/domains"
```

#### Versions v2 and Beyond
For major versions v2 and later, Go Modules requires special handling. The directory name and module path must include the version.

[Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules)
```
Example:

Module name: github.com/cam-inc/viron-go/v2
Directory structure: github.com/cam-inc/viron-go/v2/domains
Command to execute:

go mod init github.com/cam-inc/viron-go/v2

git tag v2.0.0

go get github.com/cam-inc/viron-go/v2@v2.0.0
```

#### Release Workflow
[Module release and versioning workflow](https://go.dev/doc/modules/release-workflow)

The release workflow for custom Go (Golang) libraries (assuming GitHub integration) can be designed as follows. This ensures that version-controlled code is correctly released and usable by external projects with Go Modules.

**🔁 General Release Workflow**

1. Implement and test features on a development branch.
2. Merge into the `main` branch.
3. Decide on the version number (Semantic Versioning).
4. Add a Git tag (Go Modules compatible).
5. Create a GitHub release (optional).
6. Import and use with `go get` in other projects.

**📌 Detailed Steps**

- ① Development and Testing
  - Add features or fix bugs on branches (e.g., feature/xxx, fix/yyy) and validate with unit tests or CI.

- ② Merge into the `main` Branch
  - After review (e.g., via Pull Requests), merge into the `main` branch to complete the milestone.

- ③ Decide the Version
  - Feature addition → Minor update (v1.1.0 → v1.2.0)
  - Bug fixes only → Patch update (v1.2.3 → v1.2.4)
  - Breaking changes → Major update (v1.x.x → v2.0.0)

- ④ Tag and Release
  - Create a tag via CLI (always include `v`):
    ```
    git tag v1.2.0
    git push origin v1.2.0

    This allows Go Modules to recognize the tag as a version.
    ```

- ⑤ Write Release Notes on GitHub (Optional)
  - Use GitHub UI or CLI (`gh release`) to add release notes, making changes clear to users.
    ```
    Example:

    gh release create v1.2.0 --title "v1.2.0 Release" --notes "This release includes new features and bug fixes."
    ```

## Code of Conduct

Please follow [this document](./CODE_OF_CONDUCT.md).

## License

By contributing to this project, you agree to license your contributions under the [MIT License](./LICENSE).

## Contributors

We thank everyone who has contributed to this project.

<table>
  <tr>
    <td align="center"><a href="https://github.com/takoring"><img src="https://avatars.githubusercontent.com/u/24517668?v=4" width="100px;" alt=""/><br /><sub><b>takoring</b></sub></a><br />💻</td>
  </tr>
</table>
