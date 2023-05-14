# reference: https://docs.docker.com/language/golang/build-images/
FROM golang:1.17 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /sudoku

FROM build-stage AS run-test-stage
RUN go build 
# RUN go test ./...
RUN go test -run TestSolveSudoku ./...

WORKDIR /
COPY --from=build-stage /sudoku /sudoku

ENTRYPOINT ["sudoku"]
