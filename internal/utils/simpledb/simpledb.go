package simpledb

import(
  "github.com/sdomino/scribble"
)

func ConnectDB(databaseName string) (*scribble.Driver, error) {
  db, err := scribble.New(databaseName, nil)
  if err != nil {
    return nil, err
  }
  return db, nil
}
