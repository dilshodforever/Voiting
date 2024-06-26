package postgres

import (
	"database/sql"
	"fmt"
	pb "root/genprotos"

	"github.com/google/uuid"
)

type ElectionStorage struct {
	db *sql.DB
}

func NewElectionStorage(db *sql.DB) *ElectionStorage {
	return &ElectionStorage{db: db}
}

func (bc *ElectionStorage) CreateElection(pb *pb.Election) (*pb.Void, error) {
	id := uuid.NewString()
	_, err := bc.db.Exec(`insert into election(id, name, date) 
						values($1, $2, $3)`,
		id, pb.Name, pb.Date)
	return nil, err
}

func (p *ElectionStorage) GetByIdElection(id *pb.ById) (*pb.Election, error) {
	query := `
		SELECT id, name, date
		FROM election
		WHERE id = $1
	`
	row := p.db.QueryRow(query, id.Id)

	var elec pb.Election
	err := row.Scan(&elec.Id,
		&elec.Name,
		&elec.Date)
	if err != nil {
		return nil, err
	}

	return &elec, nil
}

func (p *ElectionStorage) GetAllElection(e *pb.Election) (*pb.GetAllElection, error) {
	elecs := &pb.GetAllElection{}
	query := `select id, name, date from election where delated_at=0 `
	count := 1
	var arr []interface{}
	if len(e.Name) > 0 {
		query += fmt.Sprintf(` and name=$%d`, count)
		arr = append(arr, e.Name)
	}
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		elec := &pb.Election{}
		err = row.Scan(&elec.Id, &elec.Name, &elec.Date)
		if err != nil {
			return nil, err
		}
		elecs.Elections = append(elecs.Elections, elec)
	}
	return elecs, nil
}

func (p *ElectionStorage) UpdateElection(pb *pb.Election) (*pb.Void, error) {
	query := `
	UPDATE election
	SET  name=$1, date=$2
	WHERE id = $3
`
	_, err := p.db.Exec(query, pb.Name, pb.Date)
	return nil, err
}

func (p *ElectionStorage) DeleteElection(id *pb.ById) (*pb.Void, error) {
	query := `
	delete from election  where id = $1
`
	_, err := p.db.Exec(query, id.Id)
	return nil, err
}
