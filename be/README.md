# backend

### <a name="execution_flags"></a>Execution Flags

|Flag|Env. variable|Default value|Description|
|:----|:----|:---|:---|
|api-port|API_PORT|30303|HTTP API port number|
|initial-amount|INITIAL_AMOUNT|12000|Initial amount of money for each player|

### Running tests

`go test -cover ./...`

### API

* [POST /auth](#post_auth) - Validate player authentication and get player info.

#### <a name="post_auth"></a>POST /auth

*Validate player authentication and get player info.*

Sample request:

```
curl 'http://localhost:30303/auth' -d '{
  "token": "5aa50b703a729598d3669616ecf08254"
}'
```

Success response:

```json
{
  "auth": {
    "authenticated": true,
    "playerData": {
      "id": "c8c66a8acc1153edb9635a26d5940510",
      "name": "John",
      "cashAmount": 12000
    }
  },
  "msg": "ok",
  "status": 200
}
```