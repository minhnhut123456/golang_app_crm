package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migratemysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/minhnhut123456/golang_app_crm/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlConfig struct {
	Username string     `json:"username" yaml:"username"`
	Password string     `json:"password" yaml:"password"` 
	Server string       `json:"server" yaml:"server"`
	Schema string 			`json:"schema" yaml:"schema"`
}

type MySql struct {
	config *MySqlConfig
	DB *gorm.DB
	migrationFolder string
}

type Option func (*MySql)

func NewMySql (cfgPath string, options ...Option) (*MySql,error) {
	cfg := MySqlConfig{}
	err := helper.ReadYaml(cfgPath, &cfg)

	if(err != nil){
		return nil, err
	}

	mysql := &MySql{config: &cfg}

	for _,o := range options {
		o(mysql)
	}

	return mysql, nil
}

func WithMigation(migrationFolder string) Option{
	return func(ms *MySql) {
		ms.migrationFolder = migrationFolder
	}
}

func (m *MySql) MigrateUp() error{
	if(m.migrationFolder == "") {
		return fmt.Errorf("not found migration folder, not run migrate")
	}

	if(m.DB == nil) {
		return fmt.Errorf("please connect db before migrate")
	}

	sqlDB, err := m.DB.DB()
	if err != nil {
		return err
	}

	driver, err := migratemysql.WithInstance(sqlDB, &migratemysql.Config{})
	if err != nil {
			return err
	}

	mig, err := migrate.NewWithDatabaseInstance(	fmt.Sprintf("file://./%s",m.migrationFolder), "mysql", driver)
	if err != nil {
		return err
	}

	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
			return err
	}

	return nil
}

func (m *MySql) Connect() error{
	cfg := m.config
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Server, cfg.Schema)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})

	m.DB = db

	if err != nil{
		return err
	}

	return nil
}

