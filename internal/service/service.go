package service

type Service interface {

}

type service struct {

}

func New() (Service, error) {
	srv := &service{}
	return srv, nil
}
