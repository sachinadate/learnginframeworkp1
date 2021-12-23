package entity

import "time"

type Person struct {
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"first_namegorm:"primary_key;auto_increment" gorm:"type:varchar(32)`
	LastName  string `json:"last_name" gorm:"type:varchar(32)`
	Age       int    `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validator:"required ,email" gorm:"type:varchar(256)`
}

type Video struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=200" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//curl "localhost:8081/videos" -u"pragmatic":"reviews" -X POST -d '{"title":"blaCool","description":"bla bla bla","url":"http://bla.com","author":{"first_name":"babu","last_name":"rao","age":58,"email":"baburao@gmail.com"}}'

//curl "localhost:8081/api/videos" -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ' -X POST -d '{"title":"blaCool","description":"bla bla bla","url":"http://bla.com","author":{"first_name":"babuCool","last_name":"rao","age":58,"email":"baburao@gmail.com"}}'

//   POST:    curl "localhost:8081/api/videos" -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ' -X POST -d '{"title":"blaCool1","description":"bla bla bla","url":"http://bla1.com","author":{"first_name":"babuCool","last_name":"rao","age":58,"email":"baburao@gmail.com"}}'

//   UPDATE:   curl "localhost:8081/api/videos/1" -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ' -X PUT -d '{"title":"blaCool2","description":"malum nahi terako","url":"http://bla2.com","author":{"first_name":"babuCool","last_name":"rao","age":58,"email":"baburao@gmail.com"}}'

//   curl "localhost:8081/api/videos/1" -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ' -X POST -d '{"title":"blaCool1","description":"bla bla bla","url":"http://bla1.com","author":{"first_name":"babuCool","last_name":"rao","age":58,"email":"baburao@gmail.com"}}'

//   GET   :  curl "localhost:8081/api/videos" -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ'

// bearer

//     curl "localhost:8081/login" -u "pragmatic":"reviews"

//	   curl "localhost:8081/login" -X POST -d'{"username":"pragmatic","password":"reviews"}'

//    curl  "localhost:8081/videos/2"  -X DELETE  -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ'

//   DELETE :   curl "localhost:8081/api/videos/2" -H 'Authorization:"bearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYzOTgyODU2NCwiaWF0IjoxNjM5NTY5MzY0LCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.xOJ0C8u55UmR9MoyLCeCsxzH9a9euifV727hFIAtJpQ' -X DELETE
