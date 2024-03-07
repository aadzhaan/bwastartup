package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	//dependecies atau ketergantungan terhadap repository
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserID(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.Id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
