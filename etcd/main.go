package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/clientv3util"
	"go.etcd.io/etcd/clientv3/concurrency"

	"github.com/rs/zerolog/log"
)

const (
	etcdEndpoint   = "127.0.0.1:2379"
	dialTimeout    = 3 * time.Second
	requestTimeout = 5 * time.Second
	retryCount     = 5
	nameKey        = "name"
	nameValue      = "aakash"
)

var (
	etcdClient       *clientv3.Client
	multiKeyValueMap = map[string]string{
		"age":  "25",
		"phno": "12345",
	}
)

func main() {

	//defer cancel()
	var err error
	for i := 0; i < retryCount; i++ {
		etcdClient, err = clientv3.New(clientv3.Config{
			DialTimeout: dialTimeout,
			Endpoints:   []string{etcdEndpoint},
		})
		if err != nil {
			log.Error().Msg("Error while establishing etcd connection : " + err.Error())
			continue
		}
		break
	}
	if err != nil {
		log.Fatal().Msg("Unable to connect to etcd : " + err.Error())
		return
	}

	defer etcdClient.Close()
	log.Info().Msg("Etcd connection successful")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//putting name key with value aakash
	_, err = etcdClient.Put(ctx, nameKey, nameValue)
	if err != nil {
		log.Error().Msg("Unable to put name key")
	}

	//putting multiple key value pairs with transaction
	etcdOps := []clientv3.Op{}
	for key, val := range multiKeyValueMap {
		etcdOps = append(etcdOps, clientv3.OpPut(key, val))
	}
	_, err = etcdClient.Txn(ctx).Then(etcdOps...).Commit()
	if err != nil {
		log.Error().Msgf("Error while putting multiple keys with transaction : %s", err.Error())
	}

	//################## put, get and revision ops for key #############################
	// Insert a key value
	pr, _ := etcdClient.Put(ctx, "key", "444")
	rev := pr.Header.Revision
	log.Info().Msgf("Revision after first put : %d", rev)

	gr, _ := etcdClient.Get(ctx, "key")
	log.Info().Msgf("Value: %s, Revision on first get : %d", string(gr.Kvs[0].Value), gr.Header.Revision)

	// Modify the value of an existing key (create new revision)
	etcdClient.Put(ctx, "key", "555")

	gr, _ = etcdClient.Get(ctx, "key")
	log.Info().Msgf("Value: %s, Revision after second put on second get : %d", string(gr.Kvs[0].Value), gr.Header.Revision)

	// Get the value of the previous revision
	gr, _ = etcdClient.Get(ctx, "key", clientv3.WithRev(rev))
	log.Info().Msgf("Value for rev %d : %s", rev, string(gr.Kvs[0].Value))

	//###################################################################################

	//delete key
	_, err = etcdClient.Delete(ctx, "key")
	if err != nil {
		log.Error().Msgf("Error while deleting key : %s", err.Error())
	} else {
		log.Info().Msg("key deleted successfully")
	}

	//################## put and then get keys with options #############################
	// Insert 20 keys
	for i := 0; i < 20; i++ {
		k := fmt.Sprintf("key_%02d", i)
		etcdClient.Put(ctx, k, strconv.Itoa(i))
	}

	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(3),
	}

	//get with prefix "key", sort by keys in ascending order and get only 3 key-value pairs
	gr, _ = etcdClient.Get(ctx, "key", opts...)

	for _, item := range gr.Kvs {
		log.Info().Msgf("Key : %s, Value : %s", string(item.Key), string(item.Value))
	}

	lastKey := string(gr.Kvs[len(gr.Kvs)-1].Key)

	//getting items starting from the last fetched key
	opts = append(opts, clientv3.WithFromKey())
	gr, _ = etcdClient.Get(ctx, lastKey, opts...)

	// Skipping the first item, which was the last item from from the previous Get
	for _, item := range gr.Kvs[1:] {
		log.Info().Msgf("Key : %s, Value : %s", string(item.Key), string(item.Value))
	}
	//###################################################################################

	//delete key with prefix
	_, err = etcdClient.Delete(ctx, "key", clientv3.WithPrefix())
	if err != nil {
		log.Error().Msgf("Error while deleting keys : %s", err.Error())
	} else {
		log.Info().Msg("keys with prefix key deleted successfully")
	}

	//####################### creating key with lease ################################
	lease, _ := etcdClient.Grant(ctx, 2)

	// Insert key with a lease of 1 second TTL
	etcdClient.Put(ctx, "key", "value", clientv3.WithLease(lease.ID))

	gr, _ = etcdClient.Get(ctx, "key")
	if len(gr.Kvs) == 1 {
		log.Info().Msg("Found Key")
	}

	// Let the TTL expire
	time.Sleep(3 * time.Second)

	gr, _ = etcdClient.Get(ctx, "key")
	if len(gr.Kvs) == 0 {
		log.Info().Msg("No more key")
	}

	lease, _ = etcdClient.Grant(ctx, 2)
	// Insert key again with a lease of 2 second TTL
	etcdClient.Put(ctx, "key", "value", clientv3.WithLease(lease.ID))

	gr, _ = etcdClient.Get(ctx, "key")
	if len(gr.Kvs) == 1 {
		log.Info().Msg("Found Key")
	}

	keyChan, err := etcdClient.KeepAlive(ctx, lease.ID)
	if err != nil {
		log.Error().Msg("Error while keeping lease alive : " + err.Error())
	} else {
		go func() {
			for {
				<-keyChan
			}
		}()
	}
	time.Sleep(3 * time.Second)
	gr, _ = etcdClient.Get(ctx, "key")
	if len(gr.Kvs) == 1 {
		log.Info().Msg("Found Key")
	}

	//delete key
	_, err = etcdClient.Delete(ctx, "key")
	if err != nil {
		log.Error().Msgf("Error while deleting key : %s", err.Error())
	} else {
		log.Info().Msg("key deleted successfully")
	}

	//##############################################################################

	_, err = etcdClient.Put(ctx, "key", "value")
	if err != nil {
		log.Error().Msgf("Error while putting key : %s", err.Error())
	}
	//deleting and putting based on condition with transaction
	resp, err := etcdClient.Txn(ctx).
		If(clientv3.Cmp(clientv3util.KeyExists("key")), clientv3.Compare(clientv3.Value("key"), "=", "value")).
		Then(clientv3.OpDelete("key")).
		Else(clientv3.OpPut("key", "value1")).
		Commit()
	if err != nil {
		log.Error().Msg("Error while executing transaction : " + err.Error())
	} else {
		if resp.Succeeded {
			log.Info().Msg("Transaction Successful")
		} else {
			log.Info().Msg("Transaction Unsuccessful")
		}
	}

	//############################# watch ops on etcd keys ##############################

	//putting key1
	_, err = etcdClient.Put(ctx, "key1", "value1")
	if err != nil {
		log.Error().Msgf("Error while putting key : %s", err.Error())
	}

	watchChan := etcdClient.Watch(ctx, "key", clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithProgressNotify())
	if watchChan == nil {
		log.Error().Msg("Error while starting watch : " + err.Error())
	} else {
		for watchResp := range watchChan {
			go processWatchResp(watchResp)
		}
	}

	//################################# operation with distributed lock ##################

	// create a sessions to aqcuire a lock
	session, err := concurrency.NewSession(etcdClient)
	if err != nil {
		log.Error().Msg("Error while creating session : " + err.Error())
	}
	defer session.Close()

	/* "/distributed-lock/" is key used for distributed locking of key named "key"
	if multiple goroutines/programs try to acquire lock on the same key as shown below, only one will be allowed at a time and the rest
	will be blocked
	mutex.Lock(ctx)
	can perform any operation here and it will be executed on etcd by only goroutine/program at a time
	mutex.Unlock(ctx) */

	mutex := concurrency.NewMutex(session, "/distributed-lock/")
	// acquire lock (or wait to have it)
	if err := mutex.Lock(ctx); err != nil {
		log.Error().Msg("Error while acquiring lock : " + err.Error())
	}

	log.Info().Msg("acquired lock for " + "key")
	log.Info().Msg("Do some work in " + "key")

	time.Sleep(5 * time.Second)

	if err = mutex.Unlock(ctx); err != nil {
		log.Error().Msg("Error while releasing lock : " + err.Error())
	}
	log.Info().Msg("released lock for " + "key")
}

func processWatchResp(resp clientv3.WatchResponse) {
	if len(resp.Events) == 0 {
		log.Info().Msg("Received progress updates as no events in last 10 minutes")
	}
	for _, event := range resp.Events {
		eventType := fmt.Sprintf("%s", event.Type)
		switch eventType {
		case "PUT":
			key := string(event.Kv.Key)
			value := string(event.Kv.Value)
			if event.Kv.Version == 1 {
				log.Info().Msgf("PUT executed on new Key : %s with value : %s", key, value)
			} else {
				prevValue := string(event.PrevKv.Value)
				log.Info().Msgf("PUT executed on existing Key : %s with previous value : %s and new value : %s", key, prevValue, value)
			}
		case "DELETE":
			key := string(event.Kv.Key)
			prevValue := string(event.PrevKv.Value)
			log.Info().Msgf("DELETE executed on Key : %s with value : %s", key, prevValue)
		}
	}
}
