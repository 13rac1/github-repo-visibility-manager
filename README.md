# Github Repository Visibility Manager

Control Github Repository public/private visibility using a YAML file and your CI (Continuous Integration) system.

## Why

Github Enterprise Organizations can restrict members [from creating Public
repositories][github-restrict-create] or [making repositories
Public][github-restrict-update]. This project enables Infrastructure/DevOps
teams to enable Engineers to make repositories public using git (as the
historical record) and a CI system.

[github-restrict-create]: https://help.github.com/en/articles/restricting-repository-creation-in-your-organization
[github-restrict-update]: https://help.github.com/en/articles/restricting-repository-visibility-changes-in-your-organization

## Usage

### Manual

1. Set a `GITHUB_TOKEN` environment variable
2. Create a `repos.yaml` containing the following

    ```yaml
    organization: ORGANIZATION_NAME
    repos:
      - name: REPOSITORY_NAME
        private: false
    ```

3. Run `grvm repos.yaml`

## CI

`//TODO`

## Why not Terraform

Terraform can maintain the public/private setting of a repository, but it owns
the _entire_ repository. The [Terraform Github provider][terraform-github] is
useful in many cases, but overkill in others.

[terraform-github]: https://www.terraform.io/docs/providers/github/index.html
