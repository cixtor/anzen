FROM golang:1.13-alpine

WORKDIR "/go/bin"

# docker build -t babywaf_ariados:latest -f ariados/Dockerfile .

COPY "ariados" "/go/src/bitbucket.org/cixtor/babywaf/ariados"

RUN apk add --no-cache git \
	&& go get bitbucket.org/cixtor/babywaf/ariados

CMD ["/go/bin/ariados"]
