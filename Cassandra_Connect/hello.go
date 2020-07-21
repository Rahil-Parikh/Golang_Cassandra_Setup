package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "testing"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a tweet
	primes := [3]int{111, 222, 333}
	if err := session.Query(`INSERT INTO test_table (id, name,phone) VALUES (?, ?, ?)`,
		gocql.TimeUUID(), "ZILLION", primes).Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var name string
	var phone []int

	/*Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search*/

	if err := session.Query(
		"SELECT id, name,phone FROM test_table").Scan(
		&id, &name, &phone); err != nil {
		if err != gocql.ErrNotFound {
			log.Fatalf("Query failed: %v", err)
		}
	}
	log.Printf("ID: %v", id)
	log.Printf("Name: %v", name)
	log.Printf("Name: %v", phone)
}
