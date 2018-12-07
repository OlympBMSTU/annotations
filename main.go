package main


func Init() (*pgx.ConnPool, error) {
	conf, err := config.GetConfigInstance()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	port, err := strconv.Atoi(conf.GetDBPort())
	if err != nil {
		log.Print(err)
		return nil, err
	}
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     conf.GetDBHost(),
			User:     conf.GetDBUser(),
			Port:     uint16(port),
			Password: conf.GetDBPassword(),
			Database: conf.GetDatabase(),
		},
		MaxConnections: 5,
	}

	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return pool, nil
}

func main() {
    pool, err := Init() 
    if err != nil {
        log.Print(err)
    }
}