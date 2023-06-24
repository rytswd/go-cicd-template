# Go CI/CD Template

This repository is to test Go CI/CD setup, with essentially no third party
dependencies (other than those from GitHub).

You can go head to [`/.github/workflows`](/.github/workflows) to find more on
how the CI/CD workflows are configured.

> **Note**: This repository uses the [Reusable Workflow](https://docs.github.com/en/actions/using-workflows/reusing-workflows)
> (to be precise, [`/.github/workflows/reusable-go-cicd-for-pr.yaml`](/.github/workflows/reusable-go-cicd-for-pr.yaml)),
> which you _could_ use it as is. However, it is not designed to cover general
> use cases, and originally meant for me to test. While you can use this as is
> and manage the dependency using Git ref, I would advise copying the setup over
> to your repository for best extensibility and flexibility. I cannot promise
> that the state of the reusable workflow here will be always stable.
