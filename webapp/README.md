# Winners Tracker

You have been asked to create a web server where users can track how many games players have won.

- `GET /players/{name}` should return a number indicating the total number of wins
- `POST /players/{name}` should record a win for that name, incrementing for every subsequent POST

## Requirements

- **HTTP Server**: for exposing an API with the aforementioned endpoints;
- **Command Line**: enable the application to read input from CLI;
- **WebSockets**: the server should make use of WebSockets.
- **Standard Library**: prioritize using the standard library as much as possible.
