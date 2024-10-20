# Movie Festival

to start the application

make sure port 5432, 8080 is free. and copy env.example into .env. and run this into terminal

```shell
docker compose up -d --build
```

make sure the application with this ping API

```curl
curl --location 'http://localhost:8080/ping'
```

if te response is `pong`. The Server is running

## Requirements

## Basic Requirements

### Admin APIs

- API to create and upload movies. Required information related with a movies are at
  least title, description, duration, artists, genres, watch URL (which points to the
  uploaded movie file)
- API to update movie
- API to see most viewed movie and most viewed genre

### All Users APIs

- API to list all movies with pagination
- API to search movie by title/description/artists/genres
- API to track movie viewership

### Bonus Requirements

- Vote system:

1. API to login as an authenticated user
2. API to vote a movie as an authenticated user
3. API to unvote a movie as an authenticated user
4. API to list all of the user’s voted movie
5. API to see most voted movie and most viewed genre, as an admin

- User registration and authentication system:

1. API to register
2. API to login and logout

- Trace viewership based on watching duration
