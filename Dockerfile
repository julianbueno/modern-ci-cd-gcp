# build stage
FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -o main

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /src/main /app/
ENTRYPOINT ./main