# my-photo-blog
A Photo Blog which you can seamlessly deploy locally

## Project Description
This a photo blog which allows a user to login and add pictures via a profile. The server path of each picture is saved to the postgres database along with the picture's title, description and username of its author. Picture paths are queried from the database and used to get the served pictures on the server. 

## How to Set Up this Project Locally
- Ensure docker and docker-compose is installed. Run this command to confirm its installation:
```shell
docker -v
docker-compose -v
```
- Clone this project from git repository using 
```shell
git clone https://github.com/Emmrys-Jay/photo-blog.git
```
- Run this command in project directory
```shell
docker-compose up --build
```

## Features
<ul>
  <li> User Signup
  <li> User Login
  <li> Stateful connection using cookies after user sign in
  <li> Add a picture at a time with a title and description
  <li> View all pictures on a general home-page
  <li> Save pictures on the server
  <li> Picture path on the server saved to a postgresql database
  <li> Token authentication with JWT
  <li> Search database with regular expressions
  <li> Create, Read, update and delete (CRUD) pictures from database
  <li> Password base64-encoded before transit to database
</ul>


## Home Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/home.png" alt="home-page">

## Pictures on Home Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/pictures.png" alt="pictures">

## Search Result Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/search.png" alt="search">

## Sign In Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/signin.png" alt="sign in">

## Sign Up Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/signup.png" alt="sign up">

## Add Picture Page

<img src="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/views/screenshots/addpics.png" alt="add pictures">

## Backend Features

<ul>
  <li> Created HTTP server using net/http package
  <li> Validate request method
  <li> Authentication system using JWT
  <li> Error handling
  <li> CRUD operations
  <li> Read database with regex
  <li> Database server running on Amazon RDS
</ul>

<p>Home Page Background Photo by <a href="https://www.pexels.com/photo/brown-hummingbird-selective-focus-photography-1133957/">Philippe Donn</a> </p>
<p>HTML and CSS used for this project was gotten from <a href="https://bootstrapmade.com/mentor-free-education-bootstrap-theme/">bootstrapmade.com</a> </p>
<b>NB:</b> Instruction on setting up an Amazon RDS instance can be found <a href="https://github.com/Emmrys-Jay/my-photo-blog/blob/main/models/README.md">here</a>