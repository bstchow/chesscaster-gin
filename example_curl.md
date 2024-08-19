## POST
```
curl -H "Content-Type: application/json" \
 --request POST \
 --data '{"game_state": "active","activity_state": 1, "white_player_fid": 1,"black_player_fid": 2}' \
 -u admin:123 \
 http://localhost:8080/games/
```

## PATCH
```
curl -H "Content-Type: application/json" \
--request PATCH \
--data '{"game_state": "inactive","activity_state": 2}' \
-u admin:123 \
http://localhost:8080/games/1

```

 ## GET /active?fid=1
 ```
curl -H "Content-Type: application/json" \
 -u admin:123 \
 http://localhost:8080/games/active?fid=1
```