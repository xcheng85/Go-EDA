# 

## Design

app: composition root

dep injection happens

dep graph construction

driven adpaters: 
application
driver adapters


## Run
cd cmd/tennis-shop

go run .

swagger: http://localhost:8080/

### Get player id
curl -X 'GET' \
  'http://localhost:8080/api/players/1' \
  -H 'accept: application/json'