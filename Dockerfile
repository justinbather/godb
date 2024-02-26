FROM golang:1.21.4 AS build
WORKDIR /src

# dependencies
COPY go.mod /src
COPY go.sum /src
RUN go mod tidy

# build
COPY . /src
RUN make server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /src/bin/example /example
CMD ["/example"]
