package init

import (
	"testTrood/internal/adapter/db_adapter/mock_db"
	"testTrood/internal/adapter/nlp_adapter/mock_nlp"
	"testTrood/internal/handler/ask_handler"
	"testTrood/internal/repository/intent_repository"
	"testTrood/internal/repository/nlp_repository"
	"testTrood/internal/service/intent_service"
	"testTrood/internal/service/nlp_service"
	"testTrood/internal/usecase/intent_usecase"

	"github.com/sirupsen/logrus"
)

type Init struct {
	dbAdapter     *mock_db.MockDB
	nlpAdapter    *mock_nlp.MockNLP
	intentRepo    *intent_repository.IntentRepository
	nlpRepo       *nlp_repository.NlpRepository
	intentService *intent_service.IntentService
	nlpService    *nlp_service.NlpService
	intentUsecase *intent_usecase.IntentUsecase
	askHandler    *ask_handler.AskHandler
}

func NewInit() *Init {
	return &Init{}
}

func (i *Init) GetDbAdapter() *mock_db.MockDB {
	return i.dbAdapter
}

func (i *Init) GetNlpAdapter() *mock_nlp.MockNLP {
	return i.nlpAdapter
}

func (i *Init) GetAskHandler() *ask_handler.AskHandler {
	return i.askHandler
}

func (i *Init) GetIntentService() *intent_service.IntentService {
	return i.intentService
}

func (i *Init) GetNlpService() *nlp_service.NlpService {
	return i.nlpService
}

func (i *Init) GetIntentRepo() *intent_repository.IntentRepository {
	return i.intentRepo
}

func (i *Init) GetNlpRepo() *nlp_repository.NlpRepository {
	return i.nlpRepo
}

func (i *Init) GetIntentUsecase() *intent_usecase.IntentUsecase {
	return i.intentUsecase
}

func (i *Init) MustInit() {
	{
		logrus.Info("init adapters")
		i.dbAdapter = mock_db.NewMockDB()
		i.nlpAdapter = mock_nlp.NewMockNLP()

	}
	{
		logrus.Info("init repositories")
		i.intentRepo = intent_repository.NewIntentRepository(i.dbAdapter)
		i.nlpRepo = nlp_repository.NewNlpRepository(i.nlpAdapter)
	}
	{
		logrus.Info("init services")
		i.intentService = intent_service.NewIntentService(i.intentRepo)
		i.nlpService = nlp_service.NewNlpService(i.nlpRepo)
	}
	{
		logrus.Info("init usecases")
		i.intentUsecase = intent_usecase.NewIntentUsecase(i.intentService, i.nlpService)
	}
	{
		logrus.Info("init handlers")
		i.askHandler = ask_handler.NewAskHandler(i.intentUsecase)
	}

	logrus.Info("init done!")
}
