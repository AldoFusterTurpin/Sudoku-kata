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

*IMPORTANT:*

If you execute the 
```sh
podman build -t sudoku .
```
and in the step where the test are run you see something like

[2/2] STEP 3/6: RUN go test ./...
--> Using cache c59c88df1a42190f04b899b16e510e24b1046d5b564d5ec65b03d57fc95c458c
--> Pushing cache []:b72b53679f98fcd75e49f51eb04c1afe247026a94d87a6955bdaf5afd032e600

it means that the image layer is cached and the tests are NOT executed because nothing has changed in your code. 

If you still want to execute the tests even you code have not changed, modify with a simple random comment (or blank line) the test file.
(You probably don't need to do that, but just to inform you)
