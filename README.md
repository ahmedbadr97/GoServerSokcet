# GoServerSokcet
Go server Socket and java Client

Detective Conan and Gehad the Gone Girl


Gehad is a lovely, lovely girl with a lovely, lovely smile. She likes to read and she likes to tread. One day she decided to have a walk in the woods near her home. Her mother warned her not to venture in the woods or tread, or some monster in the woods might make her bleed. She went anyway.



It has been more than three day since Gehad went to the woods, and she never came back. Her mother was worried, so she hired the detective from Japan. Conan went investigating all around, but not a clue has been found. Many neighbors have witnessed her somewhere, so Conan decided, “why don’t we ask them all something.”



Gehad is a lovely, lovely girl with a lovely, lovely smile and she is your friend, use your computer skills to find the location of the missing friend.



Write a Golang server that receives the locations where Gehad was last sighted by the neighbors, and find the most place that most neighbors agrees on, to start the search there, known that every neighbor will send a location as a single string.

-----------------------------
### server Side
the server will start waiting for clients to connect  
it is written in go  
you have two options while the server running
1. press one to print the most place that neighbors agree on now "as it changes while clients vote for a location"
2. press two to print all locations and each loacation with no of neighbours ageee on this location
3. press '0' to exit and turn off the server
###
### client side 
written in java 
you just run the client side on local host it will ask you for the location to send it to the server and the program will shutdown
