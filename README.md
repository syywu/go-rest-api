### Go RESTful API

Recommended Test project order
Build a Go REST API to mimic one endpoint from https://jsonplaceholder.typicode.com/


See here for some best practise https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html


<br>

## Routes

| Method | Path           | Additional Info | Result                      |
| ------ | -------------- | --------------- | --------------------------- |
| GET    | /posts    |                 | all posts                   |
| GET    | /posts/{id} |                 | get post by a user ID |
| POST   | /posts     |  { body }        | create a new post           |
| PUT    | /posts/{id} |  { body }        | update post by ID           |
| DELETE | /posts/{id} |                 | delete post by ID           |

<br>