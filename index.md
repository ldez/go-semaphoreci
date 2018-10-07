## Go-SemaphoreCI

go-semaphoreci is a Go client library for accessing the [Semaphore CI](https://semaphoreci.com/) API.

* [x] [API v1](https://semaphoreci.com/docs/branches-and-builds-api.html)
* [~] [API v2](http://semaphoreci.com/docs/api-v2-overview.html)

## Examples

### API v1

```go
import (
	"log"

	"github.com/ldez/go-semaphoreci/v1"
)

func main() {
	authToken := "your-token"
	client := v1.NewClient(nil, authToken)

	projects, _, err := client.Projects.Get()
	if err != nil {
		log.Fatal(err)
	}

	for _, project := range projects {
		log.Println(project)
	}
}
```

### API v2

```go
import (
	"log"

	"github.com/ldez/go-semaphoreci/v2"
)

func main() {
	authToken := v2.TokenTransport{
		Token: "your-token",
	}

	client := v2.NewClient(authToken.Client())

	projects, resp, err := client.Projects.GetByOrg("your-organization")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP response: ", resp)

	for _, project := range projects {
		log.Println(project)
	}
}
```

### Documentation

https://godoc.org/github.com/ldez/go-semaphoreci
