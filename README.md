# go-chaos

This is a simple microservice to demonstrate changes in quality over time inside CI/CD pipelines.

It lets us bake into each version of the software different failure characterstics so we can use it to demonstrate failures over time with Continous Delivery tools 

## Turning failures on or off

To modify failures in this quickstart edit the [charts/go-chaos/values.yaml](charts/go-chaos/values.yaml) file:

### Crashing 

* `CRASH` set to `true` to enable crashing
* `CRASH_DURATION` specifies the time duration for the crash such as `1m` for one minute. Uses [go duration syntax](https://golang.org/pkg/time/#ParseDuration) such as `10s` or `2h`

### Failing HTTP requests

* `REQUEST_FAIL` lets you turn on returning failed http requests:
  * setting to `0` disables failing http requests
  * setting to `1` fails every http request
  * setting to `2` fails every other request
  * setting to N fails every Nth request
* `REQUEST_ERROR_CODE` the http status code to return for failures. Defaults to `404`    