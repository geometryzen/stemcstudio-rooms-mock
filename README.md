# stemcstudio-rooms-mock

stemcstudio-rooms mock service

```
curl -XPOST 'localhost:8082/rooms' -d '{"owner":"geometryzen","description":"Blue","public":true}'
```

```
curl -XGET 'localhost:8082/rooms/12345'
```

```
curl -XDELETE 'localhost:8082/rooms/12345'
```
