package dbconnections

import "socialnetwork/repo/dbconnections/connectors"

type Db struct {
	MongoDB *connectors.MongoDbConnector
}
