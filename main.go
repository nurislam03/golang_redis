package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

var myCache CacheItf
var localDB *sql.DB

func main() {
	InitDB()
	InitRedisCache() // comment if want to use app cache
	InitCache()      // comment if want to use redis cache

	r := mux.NewRouter()
	r.HandleFunc("/post", GetPost).Methods("GET")
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}

type CacheItf interface {
	Set(key string, data interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
}

type RedisCache struct {
	client *redis.Client
}

type AppCache struct {
	client *cache.Cache
}

func (r *RedisCache) Set(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.client.Set(key, b, expiration).Err()
}

func (r *RedisCache) Get(key string) ([]byte, error) {
	result, err := r.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}

	return result, err
}

func (r *AppCache) Set(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r.client.Set(key, b, expiration)
	return nil
}

func (r *AppCache) Get(key string) ([]byte, error) {
	res, exist := r.client.Get(key)
	if !exist {
		return nil, nil
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errors.New("Format is not arr of bytes")
	}

	return resByte, nil
}

type ToDo struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var result ToDo

	b, err := myCache.Get("todo")
	if err != nil {
		// error
		log.Fatal(err)
	}

	if b != nil {
		// cache exist
		err := json.Unmarshal(b, &result)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := json.Marshal(map[string]interface{}{
			"data":    result,
			"elapsed": time.Since(start).Microseconds(),
		})
		w.Write([]byte(b))
		return
	}

	// Get from DB
	err = localDB.QueryRow(`SELECT id, user_id, title, body FROM posts WHERE id = $1`, 1).Scan(&result.ID, &result.UserID, &result.Title, &result.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = myCache.Set("todo", result, 1*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.Marshal(map[string]interface{}{
		"data":    result,
		"elapsed": time.Since(start).Microseconds(),
	})

	if err != nil {
		log.Fatal(err)
	}

	w.Write(b)
}

func InitRedisCache() {
	myCache = &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}

}

func InitCache() {
	myCache = &AppCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	localDB = db
}
