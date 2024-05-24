# Penca Backend

### Levantar ambiente

configurar el archivo app.env en la raiz del proyecto

```bash
SERVER_ADDRESS=localhost
DB_URI=**********

JWT_KEY="**********"

#s3 config
AWS_S3_BUCKET=**********
AWS_S3_REGION=**********
AWS_S3_ACCESS_KEY_ID=**********
AWS_S3_SECRET_KEY=**********
```

### Levantar ambiente

```bash
go run src/main.go
```


## Endpoints

### Ping

#### Ping Endpoint

- **URL:** `/ping`
- **Method:** `GET`
- **Description:** Responds with a simple pong message.
- **Response:**
  - `200 OK` `{ "message": "pong" }`

#### Get All Match Results

- **URL:** `/match/result`
- **Method:** `GET`
- **Description:** Get all match results.
- **Response:**
  - `200 OK` `[ { "match_id": 1, "match_date": "2024-05-24T15:04:05Z", "team_local_id": 1, "team_visitor_id": 2, "goals_local": 1, "goals_visitor": 2, "championship_id": 1 } ]`

#### Get Match Result by ID

- **URL:** `/match/result/id/{match_id}`
- **Method:** `GET`
- **Description:** Get match result by match ID.
- **Response:**
  - `200 OK` `{ "match_id": 1, "match_date": "2024-05-24T15:04:05Z", "team_local_id": 1, "team_visitor_id": 2, "goals_local": 1, "goals_visitor": 2, "championship_id": 1 }`
  - `404 Not Found`

#### Insert Match

- **URL:** `/match/create`
- **Method:** `POST`
- **Description:** Insert a new match.
- **Request:**
  - `Query Parameters:`
    - `match_date` (string, required)
    - `team_local_id` (int, required)
    - `team_visitor_id` (int, required)
    - `championship_id` (int, required)
- **Response:**
  - `201 Created` `{ "data": "Match created successfully with id: 1" }`
  - `400 Bad Request`

#### Update Match

- **URL:** `/match/update`
- **Method:** `PUT`
- **Description:** Update an existing match.
- **Request:**
  - `Query Parameters:`
    - `match_id` (int, required)
    - `match_date` (string, required)
    - `team_local_id` (int, required)
    - `team_visitor_id` (int, required)
    - `goals_local` (int, required)
    - `goals_visitor` (int, required)
    - `championship_id` (int, required)
- **Response:**
  - `200 OK` `{ "data": "Match updated successfully with id: 1" }`
  - `400 Bad Request`

#### Delete Match

- **URL:** `/match/delete/{match_id}`
- **Method:** `POST`
- **Description:** Delete a match.
- **Response:**
  - `200 OK` `{ "data": "Match deleted successfully with id: 1" }`
  - `404 Not Found`

#### Insert Result

- **URL:** `/match/result/insert`
- **Method:** `POST`
- **Description:** Insert a match result.
- **Request:**
  - `Query Parameters:`
    - `match_id` (int, required)
    - `goals_local` (int, required)
    - `goals_visitor` (int, required)
- **Response:**
  - `201 Created` `{ "data": "Result inserted successfully with id: 1" }`
  - `400 Bad Request`

#### Get Matches Not Played Yet

- **URL:** `/match/notplayed`
- **Method:** `GET`
- **Description:** Get matches that have not been played yet.
- **Response:**
  - `200 OK` `[ { "match_id": 1, "match_date": "2024-05-24T15:04:05Z", "team_local_id": 1, "team_visitor_id": 2, "championship_id": 1 } ]`

## Running the Server

To run the server, use the following command:

```sh
go run main.go
```
