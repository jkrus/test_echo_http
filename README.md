# test_echo_http
POST /user/create{"id","name"} - create new user
GET /user/all - select all users from db
GET /user/get{"id"} - get user by id
PUT /user/update{"id","name"}
DELETE /user/delete{"id"}


name must compile with regular expression ([A-Z])\w+
id integer

