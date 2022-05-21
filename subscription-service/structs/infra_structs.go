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

func NewDatastore(name string, address string)*Datastore{
	return &Datastore{Name: name, ClusterAddress: address}
}

type Cache struct{
	Name string
	Redis redis.Client
}

func NewCache(name string, redisClient redis.Client)*Cache{
	return &Cache{Name: name, Redis: redisClient}
}

type Server struct{
	Name string
	Address string
}


