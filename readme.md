# WebCalc
Simple project based on REST API
<br>Can summarize, diffrence, multiply, divide two values and count factorial. You can also check liveness and readiness of service.</br>
<br>URL Paths:</br>
|Path|Port|Operation|
|----|----|---------|
| `/sum/{a}/{b} ` | 8080 | summarize `a` and `b`|
| `/dif/{a}/{b} ` | 8080 | diffrence `a` and `b`|
| `/mul/{a}/{b} ` | 8080 | multiply `a` and `b`|
| `/div/{a}/{b} ` | 8080 | division `a` and `b`|
| `/factorial/{a} ` | 8080 | factorial of `a` |
| `/live` | 8081 | check liveness of service |
| `/ready` | 8081 | check readiness of service |
<br>Links to additional packages used in project:</br>
* [kuberprobes](https://pkg.go.dev/github.com/Icikowski/kubeprobes)
* [zerrologs](https://github.com/rs/zerolog)
* [gorilla mux](https://pkg.go.dev/github.com/gorilla/mux)