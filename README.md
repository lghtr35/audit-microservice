# Audit Backend
This is an microservice for auditing events of other microservices. It is written in Golang and uses MongoDB and Kafka.

## Installation and Setup

You will need docker for running the project.
After cloning the project, run the command at the top of the directory 
```
docker compose up
```

Http server, Kafka and MongoDB will intialize in container. Http server has a example event producer built in. So, at the start it will fire up 8 example events to kafka and save to db.

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
