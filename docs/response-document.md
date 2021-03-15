#Response

### sign/up
```cassandraql
{
    "token": "bb",
    "loginID": "eb882228-195c-4536-bedd-4afb721b7436"
}
```

### sign/in
```cassandraql
{
    "loginID": "920450eb-f3eb-40bb-9efb-646cc94c9760"
}
```

### read/all
```
{
    "cards": [
        {
            "cardID": "b",
            "userName": "",
            "faceImage": ""
        },
        {
            "cardID": "c",
            "userName": "",
            "faceImage": ""
        },
        {
            "cardID": "d",
            "userName": "",
            "faceImage": ""
        }
    ]
}
```

### read/qr
```cassandraql
{
    "qrImage": "https://qrコードだよー"
}
```

### read?cardID=a
```cassandraql
{
    "nameImage": "http",
    "faceImage": "http",
    "tagImage": "http",
    "StatusImage": "http",
    "freeImage": "http"
}
```

### read/myCard
```cassandraql
{
    "nameImage": "http",
    "faceImage": "http",
    "tagImage": "http",
    "StatusImage": "http",
    "freeImage": "http"
}
```