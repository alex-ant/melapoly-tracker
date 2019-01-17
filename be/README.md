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
* [GET /player/:token](#get_player) - Get player info.
* [GET /players](#get_players) - Get all players.
* [DELETE /player](#delete_player) - Remove player.
* [GET /lp](#get_update_players) - Longpolling endpoint responding once any user's data has been updated.
* [GET /game/info](#get_game_info) - Get game info.
* [POST /cash/add](#post_cash_add) - Add cash for the player.
* [POST /salary/add](#post_salary_add) - Add salary for the player.
* [POST /cash/deduct](#post_cash_deduct) - Deduct cash from the player.
* [POST /cash/transfer](#post_cash_transfer) - Transfer cash between players.

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

#### <a name="get_player"></a>GET /player/:token

*Get player info.*

Sample request:

```
curl 'http://localhost:30303/player/9f4c9ed9440d34f4a5bec7b901401429'
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

#### <a name="get_players"></a>GET /players

*Get all players.*

Sample request:

```
curl -H 'X-Token: 9f4c9ed9440d34f4a5bec7b901401429' 'http://localhost:30303/players'
```

Success response:

```json
{
    "msg": "ok",
    "players": [
        {
            "id": "b6d984aff34a602f967a06479e820524",
            "name": "John",
            "cashAmount": 12000,
            "isAdmin": false,
            "you": true
        },
        {
            "id": "e760bd6e0feadd8f0d45a0ca27a26a38",
            "name": "Sarah",
            "cashAmount": 11334,
            "isAdmin": true,
            "you": false
        }
    ],
    "status": 200
}
```

#### <a name="delete_player"></a>DELETE /player

*Remove player.*

The endpoint assigns a new random admin if the removed player was one.

Sample request:

```
curl -X DELETE -H 'X-Token: 9f4c9ed9440d34f4a5bec7b901401429' 'http://localhost:30303/player'
```

Success response:

```json
{
    "msg": "ok",
    "status": 200
}
```

#### <a name="get_update_players"></a>GET /lp

*Longpolling endpoint responding once any user's data has been updated.*

`timeout <= 120` and `category = update-players` parameters are required.

Sample request:

```
curl 'http://localhost:30303/lp?timeout=30&category=update-players'
```

Response in case of an update event:

```json
{
    "events": [
        {
            "timestamp": 1547460400644,
            "category": "update-players",
            "data": {
                "updated": 1547460400644011000
            }
        }
    ]
}
```

Response if no updates have been made before the timeout:

```json
{
    "timeout": "no events before timeout",
    "timestamp": 1547460599662
}
```

#### <a name="get_game_info"></a>GET /game/info

*Get game info.*

Sample request:

```
curl 'http://localhost:30303/game/info'
```

Success response:

```json
{
    "gameInfo": {
        "initialAmount": 12000,
        "salary": 2000
    },
    "msg": "ok",
    "status": 200
}
```

#### <a name="post_cash_add"></a>POST /cash/add

*Add cash for the player.*

Only admin player can perform this action.

Sample request:

```
curl 'http://localhost:30303/cash/add' -d '{
  "token": "fab0acf89168b3f7a883fc4e63ae8918",
  "id": "da54ebf406e323d7a5948846dd57357b",
  "amount": 100
}'
```

Success response:

```json
{
    "msg": "ok",
    "status": 200
}
```

#### <a name="post_salary_add"></a>POST /salary/add

*Add salary for the player.*

Only admin player can perform this action.

Sample request:

```
curl 'http://localhost:30303/salary/add' -d '{
  "token": "fab0acf89168b3f7a883fc4e63ae8918",
  "id": "da54ebf406e323d7a5948846dd57357b"
}'
```

Success response:

```json
{
    "msg": "ok",
    "status": 200
}
```

#### <a name="post_cash_deduct"></a>POST /cash/deduct

*Deduct cash from the player.*

Only admin player can perform this action. Returns an error if insufficient amount of cash is available.

Sample request:

```
curl 'http://localhost:30303/cash/deduct' -d '{
  "token": "fab0acf89168b3f7a883fc4e63ae8918",
  "id": "da54ebf406e323d7a5948846dd57357b",
  "amount": 100
}'
```

Success response:

```json
{
    "msg": "ok",
    "status": 200
}
```

#### <a name="post_cash_transfer"></a>POST /cash/transfer

*Transfer cash between players.*

Only player *sending* cash can perform this action. Returns an error if the sender has an insufficient amount of cash available.

Sample request:

```
curl 'http://localhost:30303/cash/transfer' -d '{
  "token": "f177bf59f6a7d79b33c29dd0bb145d4b",
  "idTo": "996e1d8a3e0700e55ad981dcdfb40c9d",
  "idFrom": "fb05d539204023d3cd765e7a6ff06729",
  "amount": 200
}'
```

Success response:

```json
{
    "msg": "ok",
    "status": 200
}
```