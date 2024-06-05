package service

import (
	"context"
	"log"
	v "root/genprotos/election"
	pb "root/genprotos/public_vote"
	"root/storage/postgres"
)

type PublicVoteService struct {
	stg *postgres.PublicVoteStorage
	pb.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(stg *postgres.PublicVoteStorage) *PublicVoteService {
	return &PublicVoteService{stg: stg}
}

func (c *PublicVoteService) CreatePublicVote(ctx context.Context, publicVote *pb.PublicVote) (*v.Void, error) {
	v, err := c.stg.CreatePublicVote(publicVote)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *PublicVoteService) GetAllPublicVotes(ctx context.Context, v *v.Void) (*pb.GetAllPublicVote, error) {
	publicVotes, err := c.stg.GetAllPublicVote(v)
	if err != nil {
		log.Print(err)
	}

	return publicVotes, err
}

func (c *PublicVoteService) GetByIdPublicVote(ctx context.Context, id *v.ById) (*pb.PublicVote, error) {
	prod, err := c.stg.GetByIdPublicVote(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *PublicVoteService) UpdatePublicVote(ctx context.Context, publicVote *pb.PublicVote) (*v.Void, error) {
	v, err := c.stg.UpdatePublicVote(publicVote)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *PublicVoteService) DeletePublicVote(ctx context.Context, id *v.ById) (*v.Void, error) {
	v, err := c.stg.DeletePublicVote(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
