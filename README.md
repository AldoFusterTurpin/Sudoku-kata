# Sudoku-kata

## Kata instructions 
[Instructions](https://www.codurance.com/katalyst/sudoku-kata)

## How to run the test inside a container 
We have containerised the app using a container file.
You can run the test in the container by performing the below command (technically speaking the tests get executed while trying to build the image, not after the image is built, so we are executing them before the container is actually created/run)
```sh
podman build -t sudoku .
```

If any test do not pass, the image will not be constructed and we will see the test failing inside the container. 

Someone could say that we can just build the image without executing the go test command inside the container file and then execute the container interactively with something like

```sh
podman run sudoku:latest -it /bin/bash
```
and then when we are inside the container execute the go test ./... comand. As of today this is not needed because the different layers of the container image are cached and it works pretty fast.
