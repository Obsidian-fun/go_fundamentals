/***

Program reads from multiple databases, and the fastest query is used.
conns is a slice of database connections
query is the query string
Result can be an empty interface as type Result interface{}

***/

package main

import (
	"fmt"
)

func Query(conns []Conn, query string) Result {
	ch := make(chan Result, 1);
	for _, conn := range conns {
		go func(c conn) {
			select {
				case ch<- c.DoQuery(query): // DoQuery queries the DBs
				default:
			}
		}(conn)
	}
}


