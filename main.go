/*
Multiplexer/Router
------------------
Multiplexing is a technique used to combine and send the multiple data
streams over a single medium. The process of combining the data streams
is known as multiplexing and hardware used for multiplexing is known
as a multiplexer. Muliplexing is achieved by using a device called
Multiplexer (MUX) that combines n input lines to generate a single
output line. Multiplexing follows many-to-one, i.e., n input lines and
one output line.

0. start http web server and listen
1. route request to handler
2. read request
3. prepare response
4. return json encoded http body response

Go modules
----------
module named normally github.com/guergabo/banking
import [MODULE_NAME]/[PACKAGE_NAME]
no need to have package in the path like python
no longer need to include other package in go run main.go ...
the dependency manager takes care of that
go modules tells go that the route of all my packages will be
the module name

Hexagonal Architecture
----------------------
Created in 2005, also known as "Ports & Adapters"

Domain Business Logic in the center, allows you to isolate core business of architecture
that way behavior can be tested independently from anything else.

3 core principles:
1. Explicitly separate the code into three areas (
	left side: drive application logic, user or external program interacts with application
	business logic: implements all the business logic
	right side: infrastructure details that interacts to databse, file system, mocks, etc.
)
2. Dependencies go inside, the one who is triggering has to know the dependency
3. The boundaries are isolated by interfaces

Built for portability and ease of replacing parts -
Ports (External) - entry point into the system
Adapter (middle) - get information in and out of a port (middleman)
Application Services (thin layer) - the conductors, talk directly to adapters, should not house business logic
Domain (logic) -

Adapters never communicate directly to other adapters this allows very loose
coupling that will allow replacement etc.
*/
package main

import "banking-app/app"

// HTTP web server
func main() {
	app.Start()
}
