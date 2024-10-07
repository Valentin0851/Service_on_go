# rest-api-tutorial

# user-service

# REST API

GET /users -- list of users -- 200, 404, 500   // codes of requests
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx Header location: url
PUT /user/:id -- flutty update user -- 204/200
PATCH /users/:id -- partically update user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 400