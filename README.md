# README

Build & Start Application

```
go build
./aws-api-tool
```

Request Federated Management Console URL

```
curl -XPOST -H 'Content-Type: application/json' localhost:8000/federation -d \
'{
    "username": "hogehgoe", 
    "durations": 3600, 
    "arn": "arn:aws:iam::123456789010:role/developer-role"
}'
```

