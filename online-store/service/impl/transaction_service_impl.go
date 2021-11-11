package service_impl

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"online-store/config"
	"online-store/entity"
	"online-store/exception"
	"online-store/logger"
	"online-store/model"
	"online-store/repository"
	"online-store/service"
	"online-store/validation"
	"sync"
	"time"
)

func NewTransactionService(transactionRepository *repository.TransactionRepository, productRepository *repository.ProductRepository, userRepository *repository.UserRepository, configuration config.Config) service.TransactionService {
	return &transactionServiceImpl{
		TransactionRepository: *transactionRepository,
		Configuration: configuration,
		ProductRepository: *productRepository,
		UserRepository: *userRepository,
	}
}

type transactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	Configuration config.Config
	ProductRepository repository.ProductRepository
	UserRepository repository.UserRepository
}

type HandleRaceCondition struct {
	val int64
	m   sync.Mutex
}

func (service *transactionServiceImpl) Insert(request model.InsertTransactionRequest) {
	//run validation
	validation.ValidateCreateTransaction(request)
	value := &HandleRaceCondition{}

	//start database transaction
	client, trxOpts := service.ProductRepository.SessionTransaction()
	session, err := client.StartSession()
	exception.PanicIfNeeded(err)
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		//check product
		productTransaction,_ := service.ProductRepository.FindById(request.ProductId)
		//lock stock
		value.Set(productTransaction.Stock)
		productStock := value.Get()-request.Quantity

		//send error if insufficient stock
		if productStock<0{
			panic(exception.ValidationError{
				Message: "insufficient stock",
				Status:  http.StatusBadRequest,
			})
		}

		product := entity.Product{
			Id:        productTransaction.Id,
			Stock:     productStock,
			UpdatedAt: time.Now().UnixNano(),
		}

		//check user
		user,_ := service.UserRepository.FindById(request.UserId)
		//loock balance
		value.Set(user.Balance)
		balance := value.Get()-productTransaction.Price*(request.Quantity)

		//send error if insufficient balance
		if balance<0{
			panic(exception.ValidationError{
				Message: "insufficient balance",
				Status:  http.StatusBadRequest,
			})
		}

		user = entity.User{
			Id:        request.UserId,
			Balance:   balance,
			UpdatedAt: time.Now().UnixNano(),
		}
		//update user balance
		err = service.UserRepository.UpdateBalance(user,sessionContext)
		if err != nil{
			return nil, err
		}
		//update product stock
		err  = service.ProductRepository.UpdateStock(product, sessionContext)
		if err != nil{
			return nil, err
		}

		//add transaction
		transaction := entity.Transaction{
			Id:         uuid.New().String(),
			UserId:     request.UserId,
			ProductId:  request.ProductId,
			Quantity:   request.Quantity,
			Price:      productTransaction.Price,
			TotalPrice: productTransaction.Price*request.Quantity,
			Address:    request.Address,
			CreatedAt:  time.Now().UnixNano(),
			UpdatedAt:  time.Now().UnixNano(),
		}
		service.TransactionRepository.Insert(transaction)

		return nil, nil
	}
	//check if error in transaction
	_, err = session.WithTransaction(context.Background(), callback, trxOpts)
	//send error if error in transaction
	exception.PanicIfNeeded(err)

	logger.Info().Interface("Transaction: ", request).Msg("Add transaction success")
}

func (service *transactionServiceImpl) Get(userId string) (responses []model.GetTransactionResponse) {
	transactions := service.TransactionRepository.FindByUserId(userId)
	for _, transaction := range transactions{
		product,_ := service.ProductRepository.FindById(transaction.ProductId)
		responses = append(responses,model.GetTransactionResponse{
			Id:         transaction.Id,
			UserId:     transaction.UserId,
			ProductId:  model.GetProductResponse{
				Id:        product.Id,
				Name:      product.Name,
				Category:  product.Category,
				Image:     product.Image,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			},
			Quantity:   transaction.Quantity,
			Price:      transaction.Price,
			TotalPrice: transaction.TotalPrice,
			Address:    transaction.Address,
			CreatedAt:  transaction.CreatedAt,
			UpdatedAt:  transaction.UpdatedAt,
		})
	}

	logger.Info().Interface("Transaction: ", nil).Msg("Get transaction success")
	return responses
}

func (i *HandleRaceCondition) Get() int64 {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *HandleRaceCondition) Set(val int64) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}