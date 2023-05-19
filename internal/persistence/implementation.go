package persistence

import (
	"log"
	"sync"

	"github.com/isophtalic/License/internal/configs"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence/database"
	"github.com/isophtalic/License/internal/persistence/postgres"
	"github.com/isophtalic/License/internal/persistence/redis"
	"github.com/isophtalic/License/internal/repository"
)

var (
	postgresDB *postgres.Postgres
	redisDB    *redis.UserRedisRepository

	user                   repository.UserRepository
	account                repository.AccountRepository
	product                repository.ProductRepository
	productOption          repository.ProductOptionRepository
	productOptionDetail    repository.ProductOptionDetailRepository
	key                    repository.KeyRepository
	customer               repository.CustomerRepository
	license                repository.LicenseRepository
	licenseConfig          repository.LicenseConfigRepository
	loadUserRepositoryOnce sync.Once
)

func loadRepositoryProvider(config *configs.Configure) {
	loadUserRepositoryOnce.Do(func() {
		account = redis.NewUserAccessIDRedisRepository(config) // not change engine yet
		user = database.NewPostgresUserProvider("user", postgresDB)
		product = database.NewPostgresProductProvider("product", postgresDB)
		key = database.NewPostgresKeyProvider("keys", postgresDB)
		customer = database.NewPostgresCustomerProvider("customer", postgresDB)
		license = database.NewPostgresLicenseProvider("license", postgresDB)
		licenseConfig = database.NewPostgresLicenseConfigProvider("licenseConfig", postgresDB)
	})
}

func ConnectDatabase(config *configs.Configure) {
	postgresDB = postgres.NewPostgresQL(config)
	redisDB = redis.NewUserAccessIDRedisRepository(config)
	loadRepositoryProvider(config)
}

func MigrateDatabase() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	err := postgresDB.GetDB().AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductOption{},
		&models.Key{},
		&models.OptionDetail{},
		&models.Customer{},
		&models.License{},
		&models.LicenseConfig{},
	)
	if err != nil {
		panic(err)
	}
}

func User() repository.UserRepository {
	if user == nil {
		log.Fatalln("persistence: user not initiated")
	}
	return user
}

func Account() repository.AccountRepository {
	if account == nil {
		log.Fatalln("persistence: account not initiated")
	}
	return account
}

func Product() repository.ProductRepository {
	if product == nil {
		log.Fatalln("persistence: product not initiated")
	}
	return product
}

func ProductOption() repository.ProductOptionRepository {
	if productOption == nil {
		log.Fatalln("persistence: product-option not initiated")
	}
	return productOption
}

func ProductOptionDetail() repository.ProductOptionDetailRepository {
	if productOptionDetail == nil {
		log.Fatalln("persistence: product-option-detail not initiated")
	}
	return productOptionDetail
}

func Key() repository.KeyRepository {
	if key == nil {
		log.Fatalln("persistence: key not initiated")
	}
	return key
}

func Customer() repository.CustomerRepository {
	if customer == nil {
		log.Fatalln("persistence: customer not initiated")
	}
	return customer
}
func License() repository.LicenseRepository {
	if license == nil {
		log.Fatalln("persistence: license not initiated")
	}
	return license
}

func LicenseConfig() repository.LicenseConfigRepository {
	if licenseConfig == nil {
		log.Fatalln("persistence: LicenseConfig not initiated")
	}
	return licenseConfig
}
