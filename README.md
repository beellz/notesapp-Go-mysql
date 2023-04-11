# notesapp-Go-mysql


## notes app using GO lang and myqsl


Go lang with argv mysql 


To start with mysql we are going to use the docker mysql container 

`docker run -p 13306:3306 --name mysql-local -e MYSQL_ROOT_PASSWORD=notes -d mysql:latest `


Port : 13306 for outside and inside 3306  
name : mysql-local  
-e   : enviornmnet variable  
-d   : run the file on background  
mysql : name of the image  
latest : tag of the image  
MYSQL_ROOT_PASSWORD : root password   
  

Once you run the above command we can check the status by  

` docker ps -a `

if you want to connect to ur local client you can install mysql client locally
on ubuntu 

` sudo apt install mysql-client`

Once the client is installed you can connect it using 

` mysql --host=127.0.0.1 --port=13306 -u root -p`


to start with the code

`go mod init noteapp`

`go mod tidy `

`go run main.go`

