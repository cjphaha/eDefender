package service

type Service interface {

}

type service struct {
	c *Config
}

func New(c *Config) (Service, error) {
	srv := &service{
		c: c,
	}
	return srv, nil
}
