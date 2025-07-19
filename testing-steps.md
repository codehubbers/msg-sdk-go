Step one - Run on a separate terminal . 

nc -l 8080
// this will run a local temp. TCP server

and then 
go build
./msg-sdk.go

and then if you write anything on lc terminal and enter, It must be received by our executable.



**YOU must see **
Noise CLI SDK starting up...
Connected to localhost:8080
