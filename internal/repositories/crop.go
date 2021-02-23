package repositories

import(
  "github.com/sdomino/scribble"
)

type CropRepo struct {
  db *scribble.Driver
}

func NewCropRepo (*scribble.Driver) (*CropRepo, error) {
    return &CropRepo{
      db: db,
    }
}
