package orderbook

import (
	// "encoding/json"
	// "log"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	side      Side
	orderID   int64 
	userID    int64 
	orderType OrderType
	status    OrderStatus
	price     decimal.Decimal
	volume    decimal.Decimal
	createdAt time.Time
	volumeMu  sync.RWMutex
}

func NewOrder(side Side, userID uuid.UUID, orderType OrderType, price, volume decimal.Decimal, partialAllowed bool) *Order {
	return &Order{
		side:      side,
		orderID:   uuid.New(),
		userID:    userID,
		orderType: orderType,
		status:    Open,
		price:     price,
		volume:    volume,
		createdAt: time.Now(),
	}
}

// ID returns orderID field copy
func (o *Order) OrderID() uuid.UUID {
	return o.orderID
}

// shortOrderID returns first 4 characters of orderID (for debugging purposes)
func (o *Order) shortOrderID() string {
	return o.orderID.String()[:4]
}

// Status returns status field copy
func (o *Order) Status() OrderStatus {
	return o.status
}

// Side returns side field copy
func (o *Order) Side() Side {
	return o.side
}

// volume returns volume field copy
func (o *Order) Volume() decimal.Decimal {
	o.volumeMu.RLock()
	defer o.volumeMu.RUnlock()
	return o.volume
}

// Price returns price field copy
func (o *Order) Price() decimal.Decimal {
	return o.price
}

func (o *Order) OrderType() OrderType {
	return o.orderType
}

func (o *Order) UserID() uuid.UUID {
	return o.userID
}

func (o *Order) setStatusToPartiallyFilled(remaining decimal.Decimal) {
	o.volumeMu.Lock()
	o.volume = remaining
	o.volumeMu.Unlock()
	o.status = PartiallyFilled
	// logMsg := fmt.Sprintf("Order partially filled. Remaining volume: %s", remaining)
	// o.AppendLog(logMsg)
}

func (o *Order) setStatusToFilled() {
	o.volumeMu.Lock()
	o.volume = decimal.Zero
	o.volumeMu.Unlock()
	o.status = Filled
	// o.AppendLog("Order fully filled.")
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

// String implements Stringer interface
func (o *Order) String() string {
	return fmt.Sprintf("\norder %s:\n\tside: %s\n\ttype: %s\n\tvolume: %s\n\tprice: %s\n\ttime: %d\n", o.shortOrderID(), o.Side(), o.OrderType(), o.Volume(), o.Price(), o.CreatedAt().String())
}

// // MarshalJSON implements json.Marshaler interface
// func (o *Order) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(
// 		&struct {
// 			S         Side            `json:"side"`
// 			OrderID   string          `json:"orderID"`
// 			Timestamp time.Time       `json:"timestamp"`
// 			Volume    decimal.Decimal `json:"volume"`
// 			Price     decimal.Decimal `json:"price"`
// 		}{
// 			S:         o.Side(),
// 			OrderID:   o.OrderID(),
// 			Timestamp: o.Time(),
// 			Volume:    o.Volume(),
// 			Price:     o.Price(),
// 		},
// 	)
// }

// // UnmarshalJSON implements json.Unmarshaler interface
// func (o *Order) UnmarshalJSON(data []byte) error {
// 	obj := struct {
// 		S         Side            `json:"side"`
// 		OrderID   string          `json:"orderID"`
// 		Timestamp time.Time       `json:"timestamp"`
// 		Volume    decimal.Decimal `json:"volume"`
// 		Price     decimal.Decimal `json:"price"`
// 	}{}

// 	if err := json.Unmarshal(data, &obj); err != nil {
// 		return err
// 	}
// 	orderID, err := uuid.Parse(obj.OrderID)
// 	if err != nil {
// 		log.Fatalf("failed to parse UUID: %v", err)
// 	}
// 	o.orderID = orderID
// 	o.side = obj.S
// 	o.timestamp = obj.Timestamp
// 	o.volume = obj.Volume
// 	o.price = obj.Price
// 	return nil
// }
