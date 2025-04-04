package mock_db

import (
	"context"
	"testTrood/internal/model/intent_model"

	"github.com/sirupsen/logrus"
)

// this may be adapter for postgres/mysql/mongodb/etc with abstract methods
type MockDB struct {
	db map[string]interface{}
}

func NewMockDB() *MockDB {
	mockDB := &MockDB{
		db: make(map[string]interface{}),
	}
	mockDB.AddMocks()
	return mockDB
}

func (m *MockDB) AddMocks() {
	intentModel1 := intent_model.IntentModel{
		Id:         1,
		IntentWord: "change_password",
		Answer:     "You can change password in your profile...",
	}
	intentModel2 := intent_model.IntentModel{
		Id:         2,
		IntentWord: "change_email",
		Answer:     "You can change email in your profile...",
	}
	m.db["change_password"] = intentModel1
	m.db["change_email"] = intentModel2
}

func (m *MockDB) Exec(ctx context.Context, intent string, dest *intent_model.IntentModel) error {
	m.db[intent] = *dest
	return nil
}

func (m *MockDB) Get(ctx context.Context, intent string, dest *intent_model.IntentModel) error {
	*dest = m.db[intent].(intent_model.IntentModel)
	logrus.Info(dest)
	return nil
}

func (m *MockDB) Delete(ctx context.Context, intent string) error {
	delete(m.db, intent)
	return nil
}

func (m *MockDB) GetList(ctx context.Context) ([]string, error) {
	intents := make([]string, 0)
	for intent, _ := range m.db {
		intents = append(intents, intent)
	}
	return intents, nil
}

func (m *MockDB) Close() error {
	return nil
}
