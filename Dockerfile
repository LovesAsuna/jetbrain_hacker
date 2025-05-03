FROM golang:latest
LABEL authors="LovesAsuna"

WORKDIR /usr/src/jetbrains_hacker

COPY . .

RUN apt-get update \
  && apt-get install -y wget curl unzip

RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin
CMD ["just"]