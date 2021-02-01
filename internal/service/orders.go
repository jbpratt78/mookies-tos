package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/jbpratt/tos/internal/models"
	"github.com/jbpratt/tos/internal/pb"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// OrderDB is everything that interacts with the database involving the orders
type OrderDB interface {
	GetOrders(context.Context) ([]*pb.Order, error)
	CompleteOrder(context.Context, int64) error
	SubmitOrder(context.Context, *pb.Order) error
}

// OrderService is the abstraction of the db layer
type OrderService interface {
	OrderDB
}

var _ OrderDB = &orderService{}

type orderService struct {
	OrderDB
}

type orderDB struct {
	rw sync.RWMutex
	db *sql.DB
}

// NewOrderService creates a new order service for interacting with the
// database
func NewOrderService(db *sql.DB) (OrderService, error) {
	return &orderService{&orderDB{db: db}}, nil
}

func (o *orderDB) SubmitOrder(ctx context.Context, order *pb.Order) error {
	o.rw.Lock()
	defer o.rw.Unlock()

	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ord := models.Order{CreatedAt: time.Now().UnixNano()}

	if err = ord.Insert(ctx, tx, boil.Infer()); err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	for _, item := range order.GetItems() {
		var itm *models.Item
		itm, err = models.Items(
			qm.Where(models.ItemColumns.Name, item.GetName()),
			//			qm.Load(nil),
		).One(ctx, tx)
		if err != nil {
			return fmt.Errorf("failed to find item: %w", err)
		}

		ordItm := models.OrderItem{
			OrderID: ord.ID.Int64,
			ItemID:  itm.ID.Int64,
			Price:   null.Int64From(item.GetPrice()),
		}
		if err = ordItm.Insert(ctx, tx, boil.Infer()); err != nil {
			return fmt.Errorf("failed to insert order_item: %w", err)
		}

		for _, opt := range item.GetOptions() {
			_ = opt
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// GetOrders returns orders that are not yet marked as complete
func (o *orderDB) GetOrders(ctx context.Context) ([]*pb.Order, error) {
	o.rw.RLock()
	defer o.rw.RUnlock()

	orders, err := models.Orders(
		models.OrderWhere.CompletedAt.IsNull(),
		qm.Load(models.OrderRels.OrderItems),
		qm.Load(models.OrderRels.OrderItemOptions),
	).AllG(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders that are not complete: %w", err)
	}

	output := make([]*pb.Order, 0, len(orders))

	for _, order := range orders {
		ord := &pb.Order{}
		_ = order
		output = append(output, ord)
	}
	/*
		orders, err := o.getOrders()
		if err != nil {
			return nil, err
		}
		for _, order := range orders {
			err = o.getOrderItems(order)
			if err != nil {
				return nil, err
			}
			for _, item := range order.GetItems() {
				err = o.getOrderItemOptions(item, order.GetId())
				if err != nil {
					return nil, err
				}
			}
		}
		return orders, nil
	*/
	return output, errors.New("unimplemented")
}

func (o *orderDB) CompleteOrder(ctx context.Context, id int64) error {
	o.rw.Lock()
	defer o.rw.Unlock()

	order, err := models.FindOrderG(ctx, null.Int64From(id))
	if err != nil {
		return fmt.Errorf("failed to find order: %w", err)
	}

	order.CompletedAt = null.Int64From(time.Now().UnixNano())
	if _, err = order.UpdateG(ctx, boil.Infer()); err != nil {
		return fmt.Errorf("failed to update CompletedAt for order: %w", err)
	}
	return nil
}
