# Audit Backend
This is an microservice for auditing events of other microservices. It is written in Golang and uses MongoDB and Kafka.

## Installation and Setup

You will need to have docker for running the project.

If the project is clonned and docker is installed, you can use this command:
```
docker compose up
```
It will build and deploy the project in a docker container.
You can edit the ```docker-compose.yaml``` for different environment variables and configuration.

After running the command Http server, Kafka and MongoDB will intialize in container.

There is an swagger endpoint on the server. You can connect to it using ```http://localhost:8080/swagger/index.html```

## Usage

There is 2 endpoints which are:
```
POST /api/v1/audit/audit
GET /api/v1/audit/filter
```

First one, is an alternative way of creating event audits.
Second one, is the query endpoint. It has pagination built in, so page and size are mandatory query variables. For querying the database, you can use MongoDB queries over the queryParameters.
For example:
```
/api/v1/audit/filter?page=2&size=10&action=BILL_USER&variant_fields.billed_amount=$gte&variant_fields.billed_amount=100&variant_fields.currency="USD"
```
This will return documents with BILL_USER action and greater than or equals to 100 USD with offset 20 and limit 10

## Kafka

Kafka config:
```
auto.offset.reset: earliest
bootstrap.servers: kafka:29092
group.id: event_audits
topic: events
```
