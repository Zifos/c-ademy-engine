# C-ADEMY-ENGINE

# Setup locally

1. Install go
2. Run `sh ./hack/install.sh` to install all the tooling dependencies
3. Run `task migrate` to apply migrations to the db
4. Run `task start-api` to start the API

> Right now the project does not have any tool for reload on change, this has to be done manually

# Setup with devpod

> Devpod requires as a minimum docker installed on the machine

1. Download [git-credential-manager](https://github.com/git-ecosystem/git-credential-manager/blob/release/docs/install.md)
2. Authenticate with github by running `git-credential-manager github login`
3. Download and install [Devpod](https://devpod.sh/docs/getting-started/install) 
4. Open Devpod, and follow the instructions to setup the repository with your prefered editor
