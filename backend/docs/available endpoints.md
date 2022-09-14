# Available API endpoints
Within this doc all the available endpoints are displayed and explained.
<br />
<br />

# Movies model endpoints
**To get up and running quickly with Insomnia for testing an insomnia export is available within docs/Insomnia Movies endpoint export.json**

### Get all movies
HTTP request: ```GET localhost:4000/movies```<br>
Returns an array with objects of all the movie entries that are stored in the database.
<br><br>

### Get movie by id
HTTP request: ```GET localhost:4000/movies/{id}```<br>
Returns the databse entry with the specified id
<br><br>

### Create a new movie entry
HTTP request: ```POST localhost:4000/movies``` <br>
Create a new entry within the dabase with the following payload

HTTP Payload example
```json
{
	"imdb_id":"new imdb_id"     //Type string
	"title":"new title",        //Type string
	"rating":9.2,               //Type flaot32
	"year":1999                 //Type Int
}
```
<br>

### Delete a movie entry by id
HTTP request: ```DEL localhost:4000/movies/{id}```<br>
Delete the db entry of the given value
<br><br>

### update a movie entry by id
HTTP request: ```DEL localhost:4000/movies/{id}```<br>
Update the db entry of the given value with the provided payload.

HTTP Payload example
```json
{
	"imdb_id":"updated imdb_id"     //Type string
	"title":"updated title",        //Type string
	"rating":99,               		//Type flaot32
	"year":9999                 	//Type Int
}
```
