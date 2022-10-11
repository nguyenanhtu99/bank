# Bank Demo

## How to start the application
Set your enviroment in file `.env`
- For a MONGO_URL, include the replica set name and a seedlist of the members.
- For a MONGO_DATABASE, set the name of your database.

To start using command
```
go run main.go
```
The application should start on http://localhost:1234
## How to use the application
You need to create at least 2 accounts to be able to make a transfer. Username must only include letters and numbers with a length of 5.
1. To create account with username `user1` use command:
```
curl --location --request POST 'http://localhost:1234/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName": "user1"
}'
```
The result must be:
```
{
    "message": "Create account successfully!"
}
```
2. To create account with username `user2` use command:
```
curl --location --request POST 'http://localhost:1234/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName": "user2"
}'
```
The result must be:
```
{
    "message": "Create account successfully!"
}
```
3. Each user by default have 10 coins in balance. To transfer 5 coins from `user1` to `user2` use command:
```
curl --location --request POST 'http://localhost:1234/v1/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": "user1",
    "to": "user2",
    "amount": 5
}'
```
The result must be:
```
{
    "message": "Create transaction successfully!"
}
```
4. To get all information about account `user1` use command:
```
curl --location --request GET 'http://localhost:1234/v1/details?userName=user1'
```
The result must be:
```
{
    "data": {
        "id": <ObjectId>,
        "userName": "user1",
        "balance": 5,
        "history": [
            {
                "id": <ObjectId>,
                "from": "user1",
                "to": "user2",
                "amount": 5,
                "createdAt": 1665494158
            }
        ]
    },
    "message": "Get account successfully!"
}
```