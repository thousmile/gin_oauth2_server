# Golang Gin Gorm Redis OAuth 2.0 授权服务

![](static/images/3.png)

## 授权码模式

#### 1.在浏览器中打开如下地址

```
http://localhost:9018/code?response_type=code&client_id=nssbTtp5FO6NjZpUwP&redirect_uri=http://www.baidu.com/getCode&scope=base_user&state=K2KpBcF00JkRZPd
```

#### 2.输入用户名和密码， 默认都是： admin

![](static/images/2.png)



#### 3.登录成功后，就可以看到浏览器地址栏中的 code 参数

![](static/images/1.png)

#### 4.通过 code 换 access_token 

```
POST http://localhost:8098/access_token
Content-Type: application/json

body：
{
    "grant_type":"authorization_code",
    "client_id":"nssbTtp5FO6NjZpUwP",
    "client_secret":"nssbTtp5FO6NjZpUwP",
    "code":"上一步，中获取的 code "
}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "access_token": "eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOIeqPZhq-EH9RwV10hs9IHrLQ8ctEAhNbkICtEPLROP6pmyc_Q",
    "token_type": "Bearer ",
    "expires_in": 3600,
    "refresh_token": "eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiIxMThlY2NlkwMjUifQ.eNOiBa_gMNbkIChXY-Yn6IqMlzzw",
    "scope": "read,write"
  }
}
```



#### 5.通过上一步获取的 access_token 参数，获取用户信息

```
GET http://localhost:8098/loginInfo
Content-Type: application/json
Authorization: Bearer {{access_token}}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "tokenId": "ac44d5f3646f43cda94e72fcb1ddcfe5",
    "grantType": "authorization_code",
    "user": {
      "userId": 609429132107067392,
      "avatar": "https://images.xaaef.com/b9a7abacafd747bbb74cf7cb3de36c1e.png",
      "username": "admin",
      "mobile": "15071525233",
      "email": "3548794@qq.com",
      "nickname": "管理员",
      "gender": 1,
      "birthday": "1995-08-17",
      "status": 2,
      "adminFlag": 1
    },
    "client": {
      "clientId": "7KutwpFgFXv0hcvkBO",
      "name": "商户管理后台",
      "logo": "dwa",
      "description": "多商户管理后台",
      "clientType": 1,
      "grantTypes": [
        "*"
      ],
      "domainName": "www.xaaef.com",
      "scope": "read,write"
    }
  }
}
```



## 密码模式

```
POST http://localhost:8098/password
Content-Type: application/json

body：
{
  "username": "admin",
  "password": "admin",
  "grant_type": "password",
  "client_id": "7KutwpFgFXv0hcvkBO",
  "client_secret": "7KutwpFgFXv0hcvkBO"
}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "access_token": "eyJhbGciOiJIUdGkiOMNWklfskZwzUxMiJ9.eyJqdGkiOMNWklfskZw-Z6lEzJqHVOpr5-g",
    "token_type": "Bearer ",
    "expires_in": 3600,
    "refresh_token": "eyJhbGciOiJIiOiJIUzUzUxMiJ9.eyJqdGkiOiIiOiJIUz2YjYwZDM5ZGzAKZflB2z_79n-Q",
    "scope": "read,write"
  }
}
```



## 客户端模式

```
POST http://localhost:8098/client
Content-Type: application/json

body：
{
  "scope": "read",
  "grant_type": "client_credentials",
  "client_id": "VIUvXZmVXmOFh1gYWK",
  "client_secret": "VIUvXZmVXmOFh1gYWK"
}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "access_token": "eyJhbGciOiJIUdGkiOMNWklfskZwzUxMiJ9.eyJqdGkiOMNWklfskZw-Z6lEzJqHVOpr5-g",
    "token_type": "Bearer ",
    "expires_in": 3600,
    "refresh_token": "eyJhbGciOiJIiOiJIUzUzUxMiJ9.eyJqdGkiOiIiOiJIUz2YjYwZDM5ZGzAKZflB2z_79n-Q",
    "scope": "read,write"
  }
}
```



## 短信验证码模式

```
### 发送短信 验证码
POST http://localhost:8098/sms/send
Content-Type: application/json

body：
{
  "client_id": "nssbTtp5FO6NjZpUwP",
  "mobile": "15071525233"
}

response: 
{
  "status": 200,
  "message": "ok"
}


### 短信验证码 模式登录
POST http://localhost:8098/sms
Content-Type: application/json

body：
{
  "mobile": "15071525233",
  "code": "1168",
  "grant_type": "sms",
  "client_id": "nssbTtp5FO6NjZpUwP",
  "client_secret": "nssbTtp5FO6NjZpUwP"
}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "access_token": "eyJhbGciOiJIUdGkiOMNWklfskZwzUxMiJ9.eyJqdGkiOMNWklfskZw-Z6lEzJqHVOpr5-g",
    "token_type": "Bearer ",
    "expires_in": 3600,
    "refresh_token": "eyJhbGciOiJIiOiJIUzUzUxMiJ9.eyJqdGkiOiIiOiJIUz2YjYwZDM5ZGzAKZflB2z_79n-Q",
    "scope": "read,write"
  }
}
```



## 获取登录用户信息

```
### 获取登录的用户信息
GET http://localhost:8098/loginInfo
Content-Type: application/json
Authorization: Bearer {{tokenValue}}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "tokenId": "ac44d5f3646f43cda94e72fcb1ddcfe5",
    "grantType": "password",
    "user": {
      "userId": 609429132107067392,
      "avatar": "https://images.xaaef.com/b9a7abacafd747bbb74cf7cb3de36c1e.png",
      "username": "admin",
      "mobile": "15071525233",
      "email": "3548794@qq.com",
      "nickname": "管理员",
      "gender": 1,
      "birthday": "1995-08-17",
      "status": 2,
      "adminFlag": 1
    },
    "client": {
      "clientId": "7KutwpFgFXv0hcvkBO",
      "name": "商户管理后台",
      "logo": "dwa",
      "description": "多商户管理后台",
      "clientType": 1,
      "grantTypes": [
        "we_chat",
        "password",
        "tencent_qq",
        "sms"
      ],
      "domainName": "www.mhtled.com",
      "scope": "read,write"
    }
  }
}
```



## 刷新 token

```
### 刷新 token
POST http://localhost:8098/refresh
Content-Type: application/json
refresh_token: Bearer {{refreshToken}}

response: 
{
  "status": 200,
  "message": "ok",
  "data": {
    "access_token": "eyJhbGciOiJIUdGkiOMNWklfskZwzUxMiJ9.eyJqdGkiOMNWklfskZw-Z6lEzJqHVOpr5-g",
    "token_type": "Bearer ",
    "expires_in": 3600,
    "refresh_token": "eyJhbGciOiJIiOiJIUzUzUxMiJ9.eyJqdGkiOiIiOiJIUz2YjYwZDM5ZGzAKZflB2z_79n-Q",
    "scope": "read,write"
  }
}
```



## 退出登录

```
### 退出登录
POST http://localhost:8098/logout
Content-Type: application/json
Authorization: Bearer {{tokenValue}}

response: 
{
  "status": 200,
  "message": "ok"
}
```

