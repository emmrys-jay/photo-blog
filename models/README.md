# Database Details

<b>Type:</b> PostgreSql <br>
<b>Server:</b> Amazon RDS server <br>

<p>
<a href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html">How to create an Amazon RDS instance?</a></p>

<p>
<a href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_CommonTasks.Connect.html">Connecting to an Amazon RDS instance</a></p>

<b>DB instance identifier:</b> database-1 <br>
<b>Master username:</b> postgres <br>
<b>Master password:</b> my_photo_blog-emmrys <br>
<b>Security group name:</b> my_blog-secgroup <br>
<b>Connection name:</b> my-photoblog <br>

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

<h3>Userspb Table SQL code</h3>

CREATE TABLE userspb ( <br>
    uname VARCHAR PRIMARY KEY NOT NULL, <br>
    email VARCHAR UNIQUE NOT NULL, <br>
    psword VARCHAR <br>
); <br>

<h3>Photob Table SQL code</h3>

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

<h3>To connect to DB instance from psql:</h3>

```bash
    psql \
    --host=[host-DNS-name] \
    --port=[port] \
    --username=[username] \
    --password 
```

The password flag on the CLI does not need any values. The CLI will prompt for a password after running the command. Specify "--no-password" if the DB does not need authentication.
