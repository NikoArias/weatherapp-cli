package models

import(
  "time"
)

type Crop struct {
    id string
    name string
    createdAt time.Time
}

type CropRepository interface {
    Insert(c *Crop) error
    Update(c *Crop) error
    GetById(id string)(*Crop error)
}
