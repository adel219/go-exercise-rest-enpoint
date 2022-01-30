# go-exercise-rest-enpoint

## Drone Navigation Service (DNS) - AtlasCorp

### How to run the service using Docker:

```bash
$ docker build . -t atlas-drone:latest
$ docker run -p 8080:8080 -it atlas-drone
```

### How to run unit test with releasing the coverage:

```bash
$ go test --cover --coverprofile cov.out && go tool cover -html=out
```

### Q & A:

**Q: What instrumentation this service would need to ensure its observability and operational transparency?**

A: I would use CloudWatch, DataDog or Promethieus depending on the the hosting environment where we can stream logs and metrics about the performance of the service.

**Q: Why throttling is useful (if it is)? How would you implement it here?**

A: Definetly throttling is useful, it prevents the client from scrapping, collecting data from our service in unmanageble way. As it would help to prevent the client from exahsting the service resources. The implementation of it is dependent on the business logic of the service. In our case, we'll prevent one drone from posting its location multuple times for every second. It would viable for a drone to send its location one time per second for instance.

**Q: What we have to change to make DNS be able to service several sectors at the same time?**

A: I would add another endpoint that gets the sectors data in bulk for example

```json
[
   { "id": 1, "data": { "x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0" } }
   { "id": 2, "data": { "x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0" } }
   { "id": 3, "data": { "x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0" } }
   { "id": 4, "data": { "x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0" } }
]
```

**Q: Our CEO wants to establish B2B integration with Mom's Friendly Robot Company by allowing cargo ships of MomCorp to use DNS. The only issue is - MomCorp software expects loc value in location field, but math stays the same. How would you approach this? What’s would be your implementation strategy?**

A: I would create another endpoint for MomCorp to consume where the expected structure of request and response match their criteria, and since the logic of the two endpoints that we have is the same, we can wrap it in a function or package and consumed by those two endpoints.

**Q: Atlas Corp mathematicians made another breakthrough and now our navigation math is
even better and more accurate, so we started producing a new drone model, based on
new math. How would you enable scenario where DNS can serve both types of clients?**

A: I'll use versioning, by enabling another version of api that uses the new math. The new model of the drone would send the API version either in the HTTP header or the route, where our backend can check it and call the respective handler for the provided version.

**Q: In general, how would you separate technical decision to deploy something from
business decision to release something?**

A: I didn't get the question.
