package structs

import (
	"github.com/go-redis/redis"
)
type Service struct{
	Datastore Datastore
	Cache Cache
	Server Server
}

type Datastore struct{
	Name string
	ClusterAddress string
}

type Cache struct{
	Name string
	Redis redis.Client
}

type Server struct{
	Name string
	Address string
}
