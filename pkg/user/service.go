package user

type Service interface {
  NewUser()
}

type service struct {
  Service
}