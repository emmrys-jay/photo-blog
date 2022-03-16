# Database Details

Type: MySql <br>
Server: Amazon RDS server <br>

DB instance identifier: database-2 <br>
Master username: admin <br>
Master password: my_photo_blog-emmrys <br>
Security group name: my_blog-secgroup <br>
Connection name: my-photoblog <br>

<h3>Processes to take note of </h3>

<ul>
<li>While trying to connect to the DB, i had to set the passowrd in the connection literal to the password declared whwn defining my EC2 intance. Instead of the authentication token generated. 

<li>Modified Inbound rules of the security group to allow all traffic (IPv4 and IPv6) to be accepted.
</ul>

<h3>Database Structure</h3>

Two Tables/Schemas: 
<ul>
    <li> userspb (stands for users in the photo blog)
    <li> photoblog (stores photos and description)
</ul>

<h5>Userspb Table SQL code</h5>

CREATE TABLE userspb ( <br>
    uname VARCHAR(20) PRIMARY KEY NOT NULL, <br>
    email VARCHAR(40) UNIQUE NOT NULL, <br>
    psword VARCHAR(30) <br>
); <br>

<h5>Photob Table SQL code</h5>

CREATE TABLE photob ( <br>
    id INT AUTO_INCREMENT, <br>
    uname VARCHAR(20), <br>
    ptitle VARCHAR(20), <br>
    photo VARCHAR(100) NOT NULL, <br>
    descp VARCHAR(1024) DEFAULT NULL, <br>
    PRIMARY KEY(ptitle, uname), <br>
    FOREIGN KEY(uname) REFERENCES userspb(uname) ON DELETE CASCADE <br>
); <br>

<h3>Processes to take note of </h3>
<ul>
    <li> primary keys cannot be NULL
    <li> desc is a keyword in mysql
</ul>

# Note
<b>Bad Practice</b>: Storing pictures on a database, though i wanted to try it. But i had to store the filepath of each picture instead of storing the images with blob in the database. I couldn't find enough resources to help me with using golang to save images in mysql blob data-type. 