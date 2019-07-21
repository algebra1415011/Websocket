# Websocket
-----------------------------------------------------------------------------------------------------------------------------------------------
#modules available
1. homepage : It has a file named "homepage.go". It's just a http connection . When we hit http://localhost:8080 . It will show a message "homepage".
2. routes : It is having routes. The routes includes :
    a)"/" for homepage that is http connection
    b) "/connect" for new client request. It is a websocket connection
    c) "/activeclient" GET request for getting the information of connecting client.
3. server : In this module we have server.go file which have server related stuff like handling new request from client, ping the client after every30 sec and send the message PING! , active clients information etc.
4. cronjob: In this module we have cronjob that executes after every 30sec   
steps required to run the server
5. client : In this module we have index.html which when run the sends the websocket connection request to the server (localhost).
6. github.com and golang.org are libraries that are needed for websocket connection.
-----------------------------------------------------------------------------------------------------------------------------------------------
#steps for execution:
1. just run the command "go run main.go". It will start the server at localhost port 8080 and also start the cronjob that executes after every 
30 sec.
2. open the index.html file in browser and open any number of tabs we want the websocket connection to be created. It will created that number of websocket connection.
3. IF we want to know the number of active clients than we use "http://localhost:8080/activeclient"
4. Every output is taken as logs in the terminal itself.