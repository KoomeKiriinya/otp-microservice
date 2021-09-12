
## A simple Go lang OTP microservice: WIP

Developed using the [Go kit](https://dev.to/eminetto/microservices-in-go-using-the-go-kit-jjf) a microservice architecture.

---
```plaintext



To start using ```docker-compose``` Fill in env.example with the required variables

```
---
source .env.example
docker-compose up --build

---
```markdown

To start using ```go binary``` Fill in env.example with the required variables

```
---
source .env.example
go run .

---
```

Sample ```Curl Request``` to Generate OTP 

```
curl --location --request POST 'http://127.0.0.1:8081/generateOTP' \
--header 'Content-Type: application/json' \
--data-raw '{
    "otp_key": "079758XXX"
}'
```

Sample ```Response``` to Generate OTP 

```
---

{"otp":"7N6K-2HQT"}

---
```

Sample ```Curl Request``` to Validate OTP 

```
---
curl --location --request POST 'http://127.0.0.1:8081/validateOTP' \
--header 'Content-Type: application/json' \
--data-raw '{
    "otp_key": "079758XXX",
    "user_otp_value": "7N6K-2HQT"
}'
```
Sample ```Response``` to Generate OTP 

```
---

{"otp":"OTP Matched"}

---
```