1. make docker_up
   1. Now our kafka UI is available at: http://localhost:8086/topics
   2. Database and kafka are up
   3. Postgres table "companies" created in the database
   4. The topic is created
2. to check functionality of the service
    1. Authorize:
   ```
   curl --location 'http://localhost:11124/api/jwt/signin' \
    --header 'Content-Type: application/json' \
    --data '{
    "username": "user1",
    "password": "password1"
    }'
   ```
   I've hardcoded two users for this exercise:  user1:password1, user2:password2
   2. Refresh token:
   ```
   curl --location 'http://localhost:11124/api/jwt/refresh' \
   --header 'Content-Type: application/json' \
   --data '{
   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNjgyMjAzMTAyfQ.ceYoqjtUH8B0GBPqTIHpQk6_DhvpieBhFDhyIJgFUH4"
   }'
   ```
   We have to wait 4 minutes and 30 seconds to refresh the token.
   3. Get a company:
   ```
   curl --location 'http://localhost:11124/api/companies/{company uuid}' \
   --data ''
   ```
   4. Create new one:
   ```
   curl --location 'http://localhost:11124/api/companies' \
   --header 'token: {JWT token here}' \
   --header 'Content-Type: application/json' \
   --data '{
   "uuid": "{uuid}",
   "name": "Name of the company",
   "description": "Description of the company",
   "employeesAmount": 4,
   "registered": false,
   "type": "Sole Proprietorship"
   }'
   ```
   5. Update a company:
   ```
   curl --location --request PATCH 'http://localhost:11124/api/companies/{uuid}' \
   --header 'Content-Type: application/json' \
   --data '{
   "name": "whaaaaaaaaaaassss",
   "description": "desc",
   "employeesAmount": 333,
   "registered": true,
   "type": "fdsadsadas"
   }'
   ```
   6. Delete:
   ```
   curl --location --request DELETE 'http://localhost:11124/api/companies/{uuid}' \
   --data ''
   ```
3. Auxiliary
    1. For linters golangci-lint must be installed https://golangci-lint.run/usage/install/
    2. To run linters: make lint

Backlog:
- logging system
- tests
- storage for users, registration etc (but it depends on requirements)