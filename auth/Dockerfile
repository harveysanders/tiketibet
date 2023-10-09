# https://github.com/adoublef/clear-carrot/blob/main/Dockerfile
ARG GO_VERSION=1.21
ARG DISTROLESS=static-debian11:nonroot-amd64

FROM golang:${GO_VERSION}-alpine as base

WORKDIR /usr/src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM base as build

# -ldflags "-w -s" -> https://stackoverflow.com/questions/22267189/what-does-the-w-flag-mean-when-passed-in-via-the-ldflags-option-to-the-go-comman/22276273#22276273
# -w -> omit the DWARF symbol table
# -s -> omit the symbol table and debug information
# -buildvcs=false -> omit version control information,
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  cd cmd && go build \
  -ldflags "-w -s" \
  -buildvcs=false \
  -o /service

FROM gcr.io/distroless/${DISTROLESS} AS final
WORKDIR /opt

USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /service .

EXPOSE ${PORT}
CMD [ "/opt/service" ]
