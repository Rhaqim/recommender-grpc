# Recommendation Micro-service

This micro-service is responsible for providing recommendations to the user based on the user's preferences and the user's history.

## Architecture

The micro-service is built using the micro-service architecture. The micro-service is built using Python as the recommendation engine and Golang as the API layer. The micro-service is deployed using Docker.

## API

The micro-service provides the following API:

1. `GET /recommendations` - This API is used to get the recommendations for the user. The API
2. `POST /recommendations` - This API is used to update the recommendations for the user. The API
3. `GET /recommendations/{id}` - This API is used to get the recommendations for the user based on the user's id. The API
4. `POST /recommendations/{id}` - This API is used to update the recommendations for the user based on the user's id. The API
5. `GET /recommendations/{id}/history` - This API is used to get the history of the user based on the user's id. The API
6. `POST /recommendations/{id}/history` - This API is used to update the history of the user based on the user's id. The API

## Deployment

The micro-service is deployed using Docker. The micro-service is deployed on AWS using ECS. The micro-service is deployed using the following command:

```bash
docker-compose up
```

## Usage

The micro-service can be used to get the recommendations for the user. The micro-service can be used to update the recommendations for the user. The micro-service can be used to get the recommendations for the user based on the user's id. The micro-service can be used to update the recommendations for the user based on the user's id. The micro-service can be used to get the history of the user based on the user's id. The micro-service can be used to update the history of the user based on the user's id.
