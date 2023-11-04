package database

import (
	"DEEP-backend-image/cerrors"
	"DEEP-backend-image/database/ent"
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

// initilzeDatabase는 database를 초기화 하는 함수 입니다.
// 이 함수를 호출하면 자동적으로 instance에 초기화된 데이터베이스가 삽입됩니다.
// 이 함수는 Get함수에 의해 호출됩니다.
func initailizeDatabse() {
	username := os.Getenv("DATASOURCE_USERNAME")
	password := os.Getenv("DATASOURCE_PASSWORD")
	host := os.Getenv("DATASOURCE_HOST")
	port := os.Getenv("DATASOURCE_PORT")
	dbName := os.Getenv("DATASOURCE_DB_NAME")
	maxPoolIdle, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_IDLE_CONN"))
	maxPoolOpen, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_MAX_CONN"))
	maxPollLifeTime, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_LIFE_TIME"))
	cerrors.Sniff(err)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	drv, err := sql.Open("mysql", dsn)
	cerrors.Sniff(err)

	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetMaxOpenConns(maxPoolOpen)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// load instance
	*instance = clientWrap(*client)
}

type (
	// clientWrap 함수는 ent.Client를 감싸는 타입 입니다.
	// clientWrap을 감쌈으로써, Repository를 확장할 수 있습니다.
	clientWrap ent.Client
)

// database package는 Database를 singleton 패턴으로 구현하기 합니다. 이것이 인스턴스 입니다.
var instance *clientWrap = nil

// once는 딱 한번만 실행하게 도와주는 변수
var once sync.Once

// Get 함수는 instance를 얻을 수 있는 함수 입니다.
// 이 함수를 통해 다른 패키지에서는 호출하여 사용할 수 있습니다.
func Get() *clientWrap {
	if instance == nil {
		// 한번만 호출하도록 보장 할 수 있다.
		once.Do(func() {
			initailizeDatabse()
		})
	}

	return instance
}

////////////////////////
//    repositories    //
////////////////////////

// image

func (c *clientWrap) CreateImageX(ctx context.Context, file []byte) *ent.Image {
	return c.Image.
		Create().
		SetInstance(file).
		SaveX(ctx)
}

func (c *clientWrap) Select(ctx context.Context) {

}

func (c *clientWrap) Update(ctx context.Context) {

}

func (c *clientWrap) Delete(ctx context.Context) {

}
