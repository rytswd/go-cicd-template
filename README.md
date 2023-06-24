# Go CI/CD Template

This repository is to test Go CI/CD setup, with essentially no third party
dependencies (other than those from GitHub).

You can go head to [`/.github/workflows`](/.github/workflows) to find more on
how the CI/CD workflows are configured.

> **Note**: This repository uses the [Reusable Workflow](https://docs.github.com/en/actions/using-workflows/reusing-workflows),
> which you _could_ use it as is. However, it is not designed to be used for all
> cases, and is mainly meant for me to test the setup. While you can manage the
> dependency using Git ref, I would advise not to use it directly. I cannot
> promise that the state of reusable workflows will be always stable.
