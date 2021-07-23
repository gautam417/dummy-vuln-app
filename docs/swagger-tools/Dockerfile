FROM node:13-alpine

ENV SWAGGER_CLI_VERSION 4.0.4
ENV SNIPPET_ENRICHER_CLI_VERSION 0.0.8


RUN npm install -g swagger-cli@${SWAGGER_CLI_VERSION} snippet-enricher-cli@${SNIPPET_ENRICHER_CLI_VERSION}

CMD ["swagger-cli", "-h"]
