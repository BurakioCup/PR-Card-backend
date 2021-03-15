# Request&Response

### sign/up
- request  
header  
```cassandraql
userID:""
pass:""
```
- response
```cassandraql
{
    "token": "bb",
    "loginID": "eb882228-195c-4536-bedd-4afb721b7436"
}
```

### sign/in
- request  
header  
```cassandraql
loginID:""
```
- response
```cassandraql
{
    "loginID": "920450eb-f3eb-40bb-9efb-646cc94c9760"
}
```

### read/all
- request  
header  
```cassandraql
token:""
```
- response
```cassandraql
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
- request  
header  
```cassandraql
token:""
```
- response
```cassandraql
{
    "qrImage": "https://qrコードだよー"
}
```

### read?cardID=a
- request  
params  
```cassandraql
cardID:""
```
header  
```cassandraql
token:""
```
- response
```cassandraql
{
    "nameImage": "http",
    "faceImage": "http",
    "tagImage": "http",
    "statusImage": "http",
    "freeImage": "http"
}
```

### read/myCard
- request  
header  
```cassandraql
token:""
```
- response
```cassandraql
{
    "nameImage": "http",
    "faceImage": "http",
    "tagImage": "http",
    "statusImage": "http",
    "freeImage": "http"
}
```