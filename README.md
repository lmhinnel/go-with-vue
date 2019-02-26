# go-with-vue
Simple web app build with Golang, PostgreSQL, Docker and Vue.js

# Getting started
First, we have to create Docker images of components by executing `docker-compose build`. Afterward, we execute `docker-compose up` and this involves 3 stages: 

1. It spins up a PostgreSQL instance container that runs on port `5432`
2. After the database instance is ready, a small Golang program named newsfeeder gets executed. It fetches RSS feeds from NewYork Times and insert the entries into the database. 
3. Lastly, a Golang web server gets spun up after the database AND newsfeeder are either ready or finished.

### Main front page
By accessing to `http://localhost:8080/`, we see a list of news articles that were fetched from NewYork Times RSS feed along with Rating for each article which is implemented for this app.

### Top 5 best rated articles
By accessing to `http://localhost:8080/bestnews`, we see a JSON object that contains an array of the five best rated news articles. This object is updated every 5 minutes, so accessing the endpoint does not make the back-end hit the database, therefore reducing the database load.


