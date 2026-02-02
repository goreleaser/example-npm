# example-npm

GoReleaser example using GitHub and publishing to NPM.

```console
$ npx @goreleaser/example
$ npm i -g @goreleaser/example
$ example
```

## OIDC

This uses OIDC to publish, so we don't need to have a NPM token.

For that to work, we need:

- `id-token: write` permissions in our workflow
- `setup-node` with registry and node version 24 (older versions don't really work)
- proper setup in the package settings (owner, repository, workflow name)

And that's it :)
