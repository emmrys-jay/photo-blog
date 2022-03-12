***Database Details***

Type: MySql
Server: Amazon RDS server

DB instance identifier: database-2
Master username: admin
Master password: my_photo_blog-emmrys

Security group name: my_blog-secgroup

Connection name: my-photoblog

***Processes to take note of***

While trying to connect to the DB, i had to set the passowrd in the connection literal to the password declared whwn defining my EC2 intance. Instead of the authentication token generated. 

Modified Inbound rules of the security group to allow all traffic (IPv4 and IPv6) to be accepted.

***Database Structure***

Two Tables/Schemas
    - userspb (stands for users in the photo blog)
    - photoblog (stores photos and description)

***userspb table***

uname VARCHAR(20) PRIMARY KEY
email VARCHAR(40)
psword VARCHAR(30)

CREATE TABLE userspb (
    uname VARCHAR(20) PRIMARY KEY NOT NULL,
    email VARCHAR(40) UNIQUE NOT NULL,
    psword VARCHAR(30)
);

***photoblog***

ptitle VARCHAR(20) PRIMARY KEY
uname VARCHAR(10) PRIMARY KEY REFERENCE userspb(uname)
desc VARCHAR(100)
photo VARCHAR(30)

CREATE TABLE photob (
    id INT AUTO_INCREMENT,
    uname VARCHAR(20),
    ptitle VARCHAR(20),
    photo VARCHAR(100) NOT NULL,
    descp VARCHAR(1024) DEFAULT NULL, 
    PRIMARY KEY(ptitle, uname),
    FOREIGN KEY(uname) REFERENCES userspb(uname) ON DELETE CASCADE
);

//primary keys cannot be NULL

//desc is a keyword in mysql

Bad Practice: Storing pictures on a database, i had to store it though just for this.... lol

I had to store the filepath of each picture instead of storing the images with blob in the database. I couldn't find enough resources to help me with using golang to save images in mysql blob data-type.

ON DELETE SET NULL ??
ON DELETE CASCADE ??