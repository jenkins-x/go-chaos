# go-chaos

This is a simple microservice to demonstrate changes in quality over time inside CI/CD pipelines.

It lets us bake into each version of the software different failure characterstics so we can use it to demonstrate failures over time with Continous Delivery tools 

## Turning failures on or off

To modify failures in this quickstart edit the [charts/go-chaos/values.yaml](charts/go-chaos/values.yaml) file, in particular the `FAIL` and `CRASH_DURATION` values