package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserId(userId int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error) {

	//get campaign
	//check campaign.userid != user_id yang melakukan request

	campaign, err := s.campaignRepository.FindByID(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.Id {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignId(input.Id)

	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionsByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
