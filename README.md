
## A simple Go lang OTP microservice: WIP

Developed using the [Go kit](https://dev.to/eminetto/microservices-in-go-using-the-go-kit-jjf) a microservice architecture.




while using ```docker-compose``` Fill in env.example with the required variables

``` markdown
source .env.example
docker-compose up --build

```

while using ```go binary``` Fill in env.example with the required variables

```markdown

source .env.example
go run .

```

Sample ```Curl Request``` to Generate OTP 

```markdown
curl --location --request POST 'http://127.0.0.1:8081/generateOTP' \
--header 'Content-Type: application/json' \
--data-raw '{
    "otp_key": "079758XXX"
}'
```

Sample ```Response``` to Generate OTP 

```markdown

{"otp":"7N6K-2HQT"}

```

Sample ```Curl Request``` to Validate OTP 

```markdown

curl --location --request POST 'http://127.0.0.1:8081/validateOTP' \
--header 'Content-Type: application/json' \
--data-raw '{
    "otp_key": "079758XXX",
    "user_otp_value": "7N6K-2HQT"
}'
```
Sample ```Response``` to Validate OTP 

```markdown

{"otp_match_status":"OTP Matched"}

```

For ```tests``` and ```test coverage``` 

```markdown
go test ./... -v cover

```

```CICD```
CI/CD workflows for test and build set up on .github/workflows folder

