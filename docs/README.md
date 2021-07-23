# Sysdig API Documentation

## API Organization

APIs are defined using OpenAPI specification. All the APIs are hosted in the `docs` directory.
* `api` directory contains all the swagger definitions for the API.
* Under `api`, we have one sub-directory per service. For example,
```
docs
|
├── api
│   ├── benchmarks
│   │   └── v2
│   │       ├── internal
│   │       │   └── swagger.yaml
│   │       └── swagger.yaml
│   └── compliance
│       └── v2
│           ├── internal
│           │   └── swagger.yaml
│           └── swagger.yaml

```

## Documenting APIs

### Customer visible or External APIs
* Add a new sub-directory under `api` for the service
* To maintain versions, create a sub-directory for the version. For example,
  `api/benchmarks/v2`
* Add the required swagger spec file (eg: `api/benchmarks/v2/swagger.yaml`)
* Connect this to the root `docs/swagger.yaml` file using `$ref`. 
  For example, to add `status` endpoint for compliance service, 
  update the root `docs/swagger.yaml` file using 
  ```
  /api/compliance/v1/status:
    $ref: "api/compliance/v1/swagger.yaml#/paths/~1api~1compliance~1v1~1status"  
  ```

### Internal APIs
* All the internal or work in progress APIs should be under `api/<service>/<version>/internal/<swagger>.yaml`
* The root swagger combining all definitions for internal APIs is `docs/swagger-internal.yaml`

## Developing APIs

Here are some helpful tools to aid in developing API swagger specs.

### Live Preview

Run the following command in docs directory to generate live preview of the API docs

External or customer visible APIs: `make swagger-docs`
Internal APIs: `make swagger-docs-internal`

### Static HTML

Run the following command in docs directory to generate static HTML API doc file

`make swagger-html-docs`

## CI/CD

* API documentation is served from docs service.
* It is built using the following Jenkins Pipeline: https://sysdig-jenkins.internal.sysdig.com/job/secure-docs/
* Note that internal API spec is only visible on staging/dev env. Consult dev leads and follow PR process to expose
  APIs externally.
