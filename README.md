# stringfish

## POST /subscriptions
```
Content-Type: application/json

{
  "Type": string, // currently supported: "hackernews"
  "Source": string // see "Source" reference on what this should be for different source types
}
```

## GET /subscriptions
```
Content-Type: application/json

{
  "Type": string, // currently supported: "hackernews"
  "Source": string // see "Source" reference on what this should be for different source types
}[]
```

## GET /rss?source=&type=
```
Content-Type: application/rss+xml
```



### Source reference

Type: `hackernews`
