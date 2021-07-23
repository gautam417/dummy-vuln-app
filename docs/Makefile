SWAGGER_DOCS_PORT ?= 21210
DOCKER := docker
GO := go

PROJECT ?= api-docs
IMAGE_REPO ?= quay.io/sysdig
IMAGE_TAG ?= dev
GO_VERSION ?= 1.15
BASE_DIR=$(CURDIR)
DOCS_IMAGE_NAME := $(IMAGE_REPO)/$(PROJECT):$(IMAGE_TAG)

all: swagger-html-docs

.PHONY: swagger-docs
## preview the API documentation
swagger-docs: swagger-html-docs
	@echo "API docs preview will be running at http://localhost:$(SWAGGER_DOCS_PORT)"
	@docker run --rm -it --init -v "$(PWD):/data/" -p $(SWAGGER_DOCS_PORT):8080 badpacketsllc/redoc-cli serve -w swagger-with-snippets.json

.PHONY: swagger-docs-internal
## preview the bleeding edge API documentation
swagger-docs-internal: swagger-html-docs
	@echo "API docs preview will be running at http://localhost:$(SWAGGER_DOCS_PORT)"
	@docker run --rm -it --init -v "$(PWD):/data/" -p $(SWAGGER_DOCS_PORT):8080 badpacketsllc/redoc-cli serve -w swagger-internal-with-snippets.json

.PHONY: swagger-html-docs
## Builds a static HTML page from swagger-with-snippets.json
swagger-html-docs: swagger-snippets
	@docker run --rm -v "$(PWD):/data/" badpacketsllc/redoc-cli bundle -o index.html swagger-with-snippets.json
	@docker run --rm -v "$(PWD):/data/" badpacketsllc/redoc-cli bundle -o internal.html swagger-internal-with-snippets.json
	@echo "Cleaning up code snippets"
	sed -i -e 's/Go + Native/Go/g; s/Java + Unirest/Java/g; s/Python + Python3/Python/g; s/Node + Request/Node/g; s/Shell + Curl/Curl/g' index.html
	sed -i -e 's/Go + Native/Go/g; s/Java + Unirest/Java/g; s/Python + Python3/Python/g; s/Node + Request/Node/g; s/Shell + Curl/Curl/g' internal.html

.PHONY: swagger-html-docs
## Adds code snippets to the swagger YAML file
swagger-snippets:
	@echo "Building swagger-cli and snippet-enricher-cli docker for swagger files"
	@docker build ./swagger-tools/ -t tools/swagger-tools:latest
	@echo "Converting swagger yaml to json"
	@docker run --rm -v "$(PWD):/host/docs" tools/swagger-tools:latest swagger-cli bundle /host/docs/swagger.yaml -r -o /host/docs/swagger.json
	@docker run --rm -v "$(PWD):/host/docs" tools/swagger-tools:latest swagger-cli bundle /host/docs/swagger-internal.yaml -r -o /host/docs/swagger-internal.json
	@echo "Adding code snippets swagger files"
	@docker run --rm -v "$(PWD):/host/docs" tools/swagger-tools:latest snippet-enricher-cli --targets="go,shell_curl,node_request,java,python"  --input=/host/docs/swagger.json > swagger-with-snippets.json
	@docker run --rm -v "$(PWD):/host/docs" tools/swagger-tools:latest snippet-enricher-cli --targets="go,shell_curl,node_request,java,python"  --input=/host/docs/swagger-internal.json > swagger-internal-with-snippets.json
	@docker rmi -f tools/swagger-tools:latest

test:
	$(GO) test -cover ./...

images:
	@echo $(DOCS_IMAGE_NAME)

build-images: swagger-html-docs
	$(DOCKER) build -t $(DOCS_IMAGE_NAME) --build-arg GITHUB_API_KEY=$(GITHUB_API_KEY) --build-arg GITHUB_API_USER=$(GITHUB_API_USER) -f ./Dockerfile ..

push-images:
	$(DOCKER) push $(DOCS_IMAGE_NAME)

delete-images:
	$(DOCKER) rmi $(DOCS_IMAGE_NAME)

do-deploy:
	echo "Deploying api-docs project api-docs using Harness passing Environment: ${DEPLOY_TO} HelmChartVersion: ${IMAGE_TAG}"
	curl -X POST -H 'content-type: application/json' --url "https://app.harness.io/gateway/api/webhooks/cjSjYQmNuxWYfDTvx5IgKdGkbO34mRh7ggh90h2y?accountId=yDPcAzPrSaOCkY6B2U8mCQ" -d '{"application":"5oXI0z9FQGiCv6DqdNmrRA","parameters":{"Environment":"$(DEPLOY_TO)","Service":"api-docs","InfraDefinition_KUBERNETES":"k8s","HelmChartVersion":"$(IMAGE_TAG)"}}'

deploy-print-info:
	echo "Deploying tag $(IMAGE_TAG) to $(DEPLOY_TO)"

helm-publish:
	$(DOCKER) run -v `pwd`/.k8s:/charts -e ARTIFACTORY_USER=$(ARTIFACTORY_CREDENTIALS_USR) -e ARTIFACTORY_PASSWORD=$(ARTIFACTORY_CREDENTIALS_PSW) -e CHART_VERSION=$(CHART_VERSION) docker.internal.sysdig.com/helm-push-artifactory:1.0.0

helm-chart-version:
	@echo `docker run --rm -v "${PWD}":/workdir mikefarah/yq:4 e '.version' .k8s/helm/Chart.yaml`

.PHONY: all
