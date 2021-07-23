FROM golang:1.16 as build-env

ARG GITHUB_API_USER
ARG GITHUB_API_KEY

RUN echo -e "machine github.com\nlogin ${GITHUB_API_USER}\npassword ${GITHUB_API_KEY}" > ~/.netrc
COPY . /secure-backend
WORKDIR /secure-backend/docs

# build docs-api
RUN GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o bin/docs-api . && \
    ls -lha /secure-backend/docs/bin/docs-api

FROM quay.io/sysdig/sysdig-mini-ubi:1.1.7

COPY --from=build-env /secure-backend/docs/bin/docs-api /docs-api
COPY --from=build-env /secure-backend/docs/index.html /static/index.html
COPY --from=build-env /secure-backend/docs/internal.html /static/internal.html
CMD ["/docs-api"]
