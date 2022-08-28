package models

import (
	"time"
)

type User struct {
	//ID        primitive.ObjectID `bson:"_id,omitempy" json:"id,omitempy"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempy"`
}

type Users []*User

//curl http://localhost:8080/users

/*
curl http://localhost:8080/users \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"Name": "evy222", "Email": "em2@gg.com"}'
*/

/*
curl http://localhost:8080/users/63063fc5f8db717c87dc6cee --include --header "Content-Type: application/json" --request "PUT" --data '{"name": "evyPut", "email": "evy.post@gmail.com"}'
*/
/*
curl http://localhost:8080/users/63063fc5f8db717c87dc6cee --include --header "Content-Type: application/json" --request "DELETE"
*/

/*
curl http://localhost:8080/login --include --header "Content-Type: application/json" --request "POST" --data '{"Username": "blablabla@gmail.com", "Password": "123456"}'
*/

/*
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYmxhYmxhYmxhQGdtYWlsLmNvbSIsInVzZXIiOnRydWUsImV4cCI6MTY2MTY0NzA5MiwiaWF0IjoxNjYxNDc0MjkyLCJpc3MiOiJSb2hhbiJ9.YmxO9kFC4ySjzyhcrpBlq4VcIOhi0isdHXOsGyg2qSM" http://localhost:8080/users
*/
