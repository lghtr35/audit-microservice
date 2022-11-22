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
GET /api/v1/audit/filter
```

This is the query endpoint. It has pagination built in, so page and size are mandatory query variables. For querying the database, you can use MongoDB queries over the queryParameters.
For example:
```
/api/v1/audit/filter?page=2&size=10&action=BILL_USER&variant_fields.billed_amount=$gte&variant_fields.billed_amount=100&variant_fields.currency="USD"
```
This will return documents with BILL_USER action and greater than or equals to 100 USD with offset 20 and limit 10.

For producing events, there is a kafka-rest built-in. So, for producing events to Kafka and testing the system, you can curl there. An example curl woudl look like:
```
curl -X POST -H "Content-Type: application/vnd.kafka.json.v2+json" --data '{"records":[ {"value":  {"event_status": true,"event_id": 2,"action": "UPDATE_USER","description": "Update an user.","variant_fields": {"name": "Test Test1","user_id": "10101010","username": "test35","call_card": "new field"}}}]}' "http://localhost:8082/topics/events"
```
The data structure lookslike:
```
{
  records:[
    { 
      "key":"nullable",
        "value": {
        "action": "string",
        "description": "string",
        "event_id": int,
        "event_status": bool,
        "id": "string",
        "variant_fields": {
          object
        }
      }
    }
  ]
}
```
If you curl with this format you can see that these events will be published to Kafka broker and then consumed by the backend. 
You can put multiple events inside records array and they will be produced sequentially.
## Kafka

Kafka config:
```
auto.offset.reset: earliest
bootstrap.servers: kafka:29092
group.id: event_audits
topic: events
```
