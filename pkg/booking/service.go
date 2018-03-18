package booking

import (
  "gomua/pkg/models"
  "github.com/rickb777/date"
)

type Service interface {
  BookAppointment(date date.Date, client models.Client)
}