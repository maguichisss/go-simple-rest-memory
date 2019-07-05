# to get dependencies
go get -u github.com/gorilla/mux

# GET PEOPLE
curl http://localhost:3000/people

# GET PERSON
curl http://localhost:3000/people/2

# SET PERSON
curl -k http://localhost:3000/people -H "Content-Type: application/json" -d '{"firstname":"Maria","lastname":"Conchita"}'

# DELETE PERSON
curl -X "DELETE" http://localhost:3000/people/1
