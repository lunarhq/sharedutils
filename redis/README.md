How to use
=============

	import (
		"github.com/lunarhq/sharedutils/redis"
		"github.com/lunarhq/sharedutils/redis/client"
	)

	red := client.New()

	//Get key data from secretToken
	red.Keys.Get("secretToken")

	//Store key data in redis
	red.Keys.Store(redis.Key{})



	//Get cached response for a request
	resp := red.Request.Get(redis.Request{})

	//Store the response in the redis cache
	red.Request.Store(redis.Request{}, redis.Response{})
