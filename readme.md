# Movie Festival

to start the application

make sure port 5432, 8080 is free. and copy env.example into .env. and run this into terminal

```shell
docker compose up -d --build
```

and dont forget check the Postman docs in this file

`./MovieFestival.postman_collection.json`

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
  1. **DONE** API upload video `POST {{API_ENDPOINT}}/api/movies/upload`
  2. **DONE** API create movie `POST {{API_ENDPOINT}}/api/movies`
- API to update movie
  1. **DONE** API update movie `PUT {{API_ENDPOINT}}/api/movies/15`
- API to see most viewed movie and most viewed genre
  1. **DONE** API get most viewed movie `GET {{API_ENDPOINT}}/api/movies/popular`
  2. **DONE** get most viewed genre `GET {{API_ENDPOINT}}/api/genres/popular`

### All Users APIs

- API to list all movies with pagination
  1. **DONE** API get movies `GET {{API_ENDPOINT}}/api/movies` with additional page anda pageSize query
- API to search movie by title/description/artists/genres
  1. **DONE** API get movies `GET {{API_ENDPOINT}}/api/movies` with additional search query
- API to track movie viewership
  1. **DONE** API get most viewed movie `GET {{API_ENDPOINT}}/api/movies/popular`

### Bonus Requirements

- Vote system:

  - API to login as an authenticated user
  - API to vote a movie as an authenticated user
  - API to unvote a movie as an authenticated user
  - API to list all of the user’s voted movie
  - API to see most voted movie and most viewed genre, as an admin

- User registration and authentication system:

  - API to register

  1. **DONE** API register `POST {{API_ENDPOINT}}/api/auth/register`

  - API to login and logout

  2. **DONE** API login `POST {{API_ENDPOINT}}/api/auth/login`

- Trace viewership based on watching duration
