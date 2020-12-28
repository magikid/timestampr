# Compile stage
FROM golang:1.15 as compiler
ENV CGO_ENABLED 0
RUN apt update && apt install -y git
WORKDIR /go/src/timestamper

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

COPY ./app ./app
COPY ./conf ./conf
COPY ./public ./public
RUN revel build --debug -m prod . output

FROM golang:1.15-alpine as main
WORKDIR /app
COPY --from=compiler /go/src/timestamper/output .
ENTRYPOINT [ "sh" ]
CMD [ "/app/run.sh" ]