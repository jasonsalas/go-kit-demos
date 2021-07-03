# Go Kit demos

## Some simple tech talk demos for building microservices with [Go Kit](https://gokit.io/), based off the project's excellent [docs](https://gokit.io/examples). 


### Let's build the [string manpulation microservice](https://gokit.io/examples/stringsvc.html)


*HOW TO RUN THE DEMO*

- Compile the binary: 

_go build ./..._
- Run the server: 

_./gokit-demo_
- Tap the microservice's endpoints: 

_curl -X POST localhost:8882/uppercase -d '{"s":"hi there!"}'_ 

_curl -X POST localhost:8882/count -d '{"s":"hi there!"}'_
