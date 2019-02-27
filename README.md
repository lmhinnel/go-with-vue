# go-with-vue
Simple web app build with Golang, PostgreSQL, Docker and Vue.js

# Getting started
First, we have to create Docker images of components by executing `docker-compose build`. Afterward, we execute `docker-compose up` and this involves 3 stages: 

1. It spins up a PostgreSQL instance container that runs on port `5432`
2. After the database instance is ready, a small Golang program named newsfeeder gets executed. It fetches RSS feeds from NewYork Times and insert the entries into the database. 
3. Lastly, a Golang web server gets spun up after the database and newsfeeder are ready.

### Main front page
By accessing to `http://localhost:8080/`, we see a list of news articles that were fetched from NewYork Times RSS feed along with Rating for each article which is implemented for this app.

### Top 5 best rated articles
By accessing to `http://localhost:8080/bestnews`, we see a JSON object that contains an array of the five best rated news articles. This object is updated every 5 minutes, so accessing the endpoint does not make the back-end hit the database, therefore reducing the database load.

# Architectural Decision
One of the common design problems when developing a web application with Golang is how to manage a database connection such that the application runs efficiently and is unit-testable. In order to solve these issues, we use a context object  and Golang's `interface` feature for our database type.
In `api/main.go` we can see a struct:
```go

type Context struct {
	DB       database.Database
	BestNews *models.BestNews
}

```
that holds a database object that implements `Database` interface and a bestnews object that gets updated every 5 minutes. The `Context` object gets passed to the handlers for endpoints and each handler can have an access to the database object.
This brings us the following benefits:
1. Eliminating the use of global variable for a database connection object.
2. Each handler does not have to instantiate a new database object, which is computationally expensive.
3. Making the handlers unit-testable by implementing a mock database object that implements the database interface.
