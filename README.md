<h1>Premier League API</h1>
	<p>This is an unofficial Premier League API client for pulling standings, and fixtures data from the Premier League. The API is built using Go and Fiber, and the data is scraped from the <a href="https://onefootball.com/en/competition/premier-league-9/table">Premier League website</a>.</p>


<h2>API Endpoints</h2>

<p>The application provides the following API endpoints:</p>

- `GET /api/standings`
- `GET /api/standings/position/:position` Example: /api/standings/position/1
- `GET /api/standings/team/:team` Example: /api/standings/team/liverpool
- `GET /api/standings/zone/:zone` Example: /api/standings/zone/champions (champions, europa, conference, relegation)
<p>The JSON object contains an array of strings, where each string represents a team's position, name, number of games played, wins, draws, losses, goal difference, and total points.</p>
<p>The API returns a JSON object with the following structure:</p>
<pre><code>[
      "position",
      "team",
      "played",
      "wins",
      "draws",
      "losses",
      "goal_difference",
      "points"
    ]</code></pre>


<h2>Setup Details</h2>
1. Clone the repository:

```sh
git clone https://github.com/ahmadammarm/premier-league-api.git
```

2. Navigate to the project directory:

```sh
cd premier-league-api
```

3. Install the project dependencies:

```sh
go mod download
```

4. Run the project:

```sh
go run cmd/main.go
```

<h2>Getting Started with Docker</h2>

1. Clone the repository:
```sh
git clone https://github.com/ahmadammarm/go-rest-api-template.git
```

2. Navigate to the project directory:

```sh
cd go-rest-api-template
```

3. Build the docker image:

```sh
docker build -t premier-league-api:your_tag .
```


4. Start the container with port forwarding from the image:
```sh
docker run -p 8080:8080 premier-league-api:your_tag
```


The project will be available at:

`http://localhost:8080`

5. To stop the container:
```sh
docker container stop your_container_name
```

### This project still needs more improvements and also available for your contributions, thank you :)


<H2>Disclaimer</H2>
This project is created solely for learning and educational purposes. It is not intended for production-level use or commercial applications.
