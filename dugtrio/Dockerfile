FROM golang:1.13-alpine

WORKDIR "/go/bin"

# docker build -t babywaf_dugtrio:latest -f Dockerfile .

COPY "dugtrio" "/go/src/bitbucket.org/cixtor/babywaf/dugtrio"

RUN apk add --no-cache git \
	&& go get bitbucket.org/cixtor/babywaf/dugtrio

CMD ["/go/bin/dugtrio"]
