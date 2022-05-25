# Database Details

Type: PostgreSql <br>
Server: Amazon RDS server <br>

DB instance identifier: database-1 <br>
Master username: postgres <br>
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
    uname VARCHAR PRIMARY KEY NOT NULL, <br>
    email VARCHAR UNIQUE NOT NULL, <br>
    psword VARCHAR <br>
); <br>

<h5>Photob Table SQL code</h5>

CREATE TABLE photob ( <br>
    id BIGSERIAL, <br>
    uname VARCHAR, <br>
    ptitle VARCHAR, <br>
    photo VARCHAR NOT NULL, <br>
    descp VARCHAR DEFAULT NULL, <br>
    PRIMARY KEY(id, uname), <br>
    FOREIGN KEY(uname) REFERENCES <br> userspb(uname) ON DELETE CASCADE <br>
); <br>

<h3>Processes to take note of </h3>
<ul>
    <li> primary keys cannot be NULL
    <li> desc is a keyword in mysql
</ul>

# Note
<b>Bad Practice</b>: Storing pictures on a database, though i wanted to try it. But i had to store the filepath of each picture instead of storing the images with blob in the database. I couldn't find enough resources to help me with using golang to save images in mysql blob data-type. 

<h3>To connect to DB instance from psql:</h3>
```bash
    psql \
    --host=[host-DNS-name] \
    --port=[port] \
    --username=[username] \
    --password 
```

<b>NB:</b> The password flag on the CLI does not need any values. The CLI will prompt for a password after running the command. Specify "--no-password" if the DB does not need authentication.

