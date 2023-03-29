package memory

import "github.com/josealvaradoo/saime-status-bot/src/model"

type Memory struct {
	Saime model.Saime
}

func NewMemory() Memory {
	return Memory{
		Saime: model.Saime{
			Status: "",
		},
	}
}

func (m *Memory) Get() (model.Saime, error) {
	saime := m.Saime

	if saime.Status == "" {
		return model.Saime{}, model.ErrStatusNotExists
	}

	return saime, nil

}

func (m *Memory) Post(value string) error {
	if value == "" {
		return model.ErrStatusCanNotBeUpdated
	}

	saime := model.Saime{Status: value}
	m.Saime = saime

	return nil
}
