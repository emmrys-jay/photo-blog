# my-photo-blog
My first full web project with a PostgreSql database runnning on Amazon RDS.

# Project Description
This a photo-blog which allows a user to login and add pictures via a profile. The server path of each picture is save to the database along with the picture title, description and username. Picture paths are queried from the database and used to get the served pictures on the server. 

# Features
<ul>
  <li> User Signup
  <li> User Login
  <li> Stateful connection using cookies after user sign in
  <li> Add a picture at a time with a title and description
  <li> View all pictures on a general home-page
  <li> Save pictures on the server
  <li> Picture path on the server saved to a postgresql database
  <li> PostgreSQL database server running on AWS RDS service
  <li> Token authentication with JWT
  <li> Search database with regular enpressions
  <li> Create, Read, update and delete (CRUD) from database
  <li> Password base64-encoded before transit to database
</ul>

<b>Home Page Background Photo by Philippe Donn</b>: https://www.pexels.com/photo/brown-hummingbird-selective-focus-photography-1133957/

# Home Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/home.png" alt="home-page">