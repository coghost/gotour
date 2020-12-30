module hello

go 1.15

replace github.com/coghost/greetings => ../greetings

require (
	github.com/coghost/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/tour v0.0.0-20201207214521-004403599411 // indirect
)
