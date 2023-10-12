# Pants GO Backend Endless Loop

This repository is a minimal reproduction of a bug I found in the pants build go module.

## Steps to reproduce

```bash
pants test ::
```
The execution will hang at the message
```bash
Prepare Go Test Binary
```

## Expected behavior

Test should run and finish.

```
cd src/go/test
go test -v ./...
```
Executes the test successfully.