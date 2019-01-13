# backend

### <a name="execution_flags"></a>Execution Flags

|Flag|Env. variable|Default value|Description|
|:----|:----|:---|:---|
|api-port|API_PORT|30303|HTTP API port number|
|initial-amount|INITIAL_AMOUNT|12000|Initial amount of money for each player|
|salary|SALARY|2000|Salary being issued every loop|

### Running tests

`go test -cover ./...`

### API

* [POST /player](#post_player) - Register a new player.
* [POST /auth](#post_auth) - Validate player authentication and get player info.

#### <a name="post_player"></a>POST /player

*Register a new player.*

Sample request:

```
curl 'http://localhost:30303/player' -d '{
  "name": "John"
}'
```

Success response:

```json
{
    "msg": "ok",
    "player": {
        "token": "9f4c9ed9440d34f4a5bec7b901401429"
    },
    "status": 200
}
```

#### <a name="post_auth"></a>POST /auth

*Validate player authentication and get player info.*

Sample request:

```
curl 'http://localhost:30303/auth' -d '{
  "token": "9f4c9ed9440d34f4a5bec7b901401429"
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
      "cashAmount": 12000,
      "isAdmin": true
    }
  },
  "msg": "ok",
  "status": 200
}
```