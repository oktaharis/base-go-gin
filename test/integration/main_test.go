package integration_test

import (
	"base-gin/config"
	"base-gin/domain/dao"
	"base-gin/repository"
	"base-gin/rest"
	"base-gin/server"
	"base-gin/service"
	"base-gin/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	password = "Paswd123"
)

var (
	cfg config.Config
	db  *gorm.DB
	app *gin.Engine

	accountRepo *repository.AccountRepository
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	setup()

	os.Exit(m.Run())
}

func setup() {
	if err := godotenv.Load("./../../.env.test"); err != nil {
		log.Fatal(fmt.Errorf("Test.Integration: Can not find .env.test on root dir"))
	}

	cfg = config.NewConfig()

	storage.InitDB(cfg)
	db = storage.GetDB()
	teardownDB()
	setupDB()

	repository.SetupRepositories()
	accountRepo = repository.GetAccountRepo()

	createDummyAccount()

	service.SetupServices(&cfg)

	app = server.Init(&cfg, accountRepo)
	rest.SetupRestHandlers(app)
}

func teardownDB() {
	_ = db.Migrator().DropTable(
		&dao.Account{},
	)
}

func setupDB() {
	_ = db.AutoMigrate(
		&dao.Account{},
	)
}

func createDummyAccount() *dao.Account {
	account, _ := dao.NewUser("admin", password, cfg.AuthN.PasswordEncryptionSecret)
	accountRepo.Create(&account)
	return &account
}

func doTest(
	method, url string,
	body interface{},
	authAccessToken string,
) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	r, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
	if authAccessToken != "" {
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authAccessToken))
	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)
	if w.Code >= 400 {
		fmt.Printf("[REQUEST] %s \n", string(requestBody)) //nolint:forbidigo //debug
		fmt.Printf("[RESPONSE] %s \n", w.Body.String())    //nolint:forbidigo //debug
	}
	return w
}
