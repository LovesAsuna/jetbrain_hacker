FROM golang:latest
LABEL authors="LovesAsuna"

WORKDIR /usr/src/jetbrains_hacker

COPY . .

RUN go mod tidy
RUN go build -v -o jetbrains_hacker

CMD ["./jetbrains_hacker", "run-server"]