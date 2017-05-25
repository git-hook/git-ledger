# git-ledger

> Track your local git clones

**git-ledger** acts as a simple key-value store, remembering the
location of git-repositories on local filesystems.

## Install

WIP
```bash
go get github.com/git-hook/git-ledger
cd $GOPATH/src/github.com/git-hook/git-ledger
go build .
go get .
```

## API

### add [path-or-slug]

Start tracking an existing repository.

### find [path-or-slug]

Print the location of a tracked repository.

### ls

Print all tracked repositories.

### rm [path-or-slug]

Stop tracking an existing repository.

## Related

- [grit](https://github.com/jmalloc/grit)

## License

MIT © Eric Crosson
