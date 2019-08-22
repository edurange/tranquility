auth route:
step -1: set correct postgres and redis connection variables!!!!!!!!!!

when creating scenario, write scenerio ID and secret key to database(add to datatype)
then, write that information to the bashrc file of every scenario

then, write the curl command to the bashrc as well, and include the information in the headers

add ability for ruby inst to access active queue!!


todo:
test all on device dialing in to another server with exposed port 80 
-how to serve this in production????

schema:
{
	user: $(whoami),
	time: $(date),
	command: $(echo $BASH_HISTORY)
}


step1: write this line to bashrc
export PROMPT_COMMAND='history -a; histJSON=$(echo {\"user\": \"$(whoami)\", \"time\": $(date +%s), \"command\": \"$(tail -n 1 ~/.bash_history)\"}); curl -s -d "$histJSON" -H "Content-Type: application/json" -H "uuid: deb9b30a-9c41-42e2-8fc1-41fe1ecc7084" -H "secret: foo" -X POST http://localhost:8080/logger >> /dev/null'


forV2:
export PROMPT_COMMAND='history -a; histJSON=$(echo {\"user\": \"$(whoami)\", \"time\": $(date +%s), \"command\": \"$(tail -n 1 ~/.bash_history)\"}); curl -s -d "$histJSON" -H "Content-Type: application/json" -H "uuid: deb9b30a-9c41-42e2-8fc1-41fe1ecc7084" -X POST http://localhost:8080/logger >> /dev/null'


forV3:
export PROMPT_COMMAND='history -a; histJSON=$(echo {\"user\": \"$(whoami)\", \"time\": $(date +%s), \"command\": \"$(history 1 | cut -c 8-)\"}); curl -s -d "$histJSON" -H "Content-Type: application/json" -H "uuid: deb9b30a-9c41-42e2-8fc1-41fe1ecc7084" -X POST http://localhost:8080/logger >> /dev/null'

step2: expose port 8080(for now)

step3: set correct redis and postgres connection vars

step 4: run server on same box as edurange_server 

big bug: literally can't do quotes!!! single or double!!! need to fix this!!!!!!!!

2 options:
-possibly move to microservece model and host all this crap elsewhere
-bake into edu-server

MUST SWITCH OVER TARGET URL ON EDURANGE-SERVER!!!!!!!
also, it's not going to work, because localhost != my box. How can we make it so? 
-will have to take hit for development(server only move)
-will have to open port 8080 on local machine, and use literal URL....
or just make it a microservice! lol.  

still have to worry about exposing port 80.............

GOOS=linux GOARCH=amd64 go build -o my_app .

https://github.com/edurange/instant-history/raw/master/tranquility/main


echo -e "$?\t$(date --utc +%FT%T.%3NZ)\t$(whoami)\t$(pwd)\t$(history 1 | cut -c 8-)" >> ~/.bash_history.log

<% if instance.os != "nat" %>

open port 80 on nat......
and make it accessable to everything!!!

NAT might also not be writable without changing chef scripts on AWS...
also check that UUID is being set!

