package storage

import (
	"encoding/json"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

// Structs for DDL queries.

// Account contains data for table "account".
type Account struct {
	Id         int             `db:"id" json:"id"`
	Username   string          `db:"username" json:"username"`
	Name       string          `db:"name" json:"name"`
	Password   string          `db:"password" json:"password"`
	Permission json.RawMessage `db:"permission" json:"permission"`
	Deleted    bool            `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// Device contains data for table "device".
type Device struct {
	Id         int      `db:"id" json:"id"`
	StoreId    null.Int `db:"store_id" json:"storeId"`
	Name       string   `db:"name" json:"name"`
	Identifier string   `db:"identifier" json:"identifier"`
	Allowed    bool     `db:"allowed" json:"allowed"`
}

// Employee contains data for table "employee".
type Employee struct {
	Id         int            `db:"id" json:"id"`
	Name       string         `db:"name" json:"name"`
	Username   string         `db:"username" json:"username"`
	Gender     string         `db:"gender" json:"gender"`
	Religion   string         `db:"religion" json:"religion"`
	Phone      string         `db:"phone" json:"phone"`
	Address    string         `db:"address" json:"address"`
	City       string         `db:"city" json:"city"`
	StartYear  int            `db:"start_year" json:"startYear"`
	EndYear    null.Int       `db:"end_year" json:"endYear"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// Supplier contains data for table "supplier".
type Supplier struct {
	Id         int            `db:"id" json:"id"`
	Name       string         `db:"name" json:"name"`
	Phone      string         `db:"phone" json:"phone"`
	Address    string         `db:"address" json:"address"`
	City       string         `db:"city" json:"city"`
	Info       null.String    `db:"info" json:"info"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// Store contains data for table "store".
type Store struct {
	Id         int            `db:"id" json:"id"`
	Name       string         `db:"name" json:"name"`
	Phone      string         `db:"phone" json:"phone"`
	Address    string         `db:"address" json:"address"`
	City       string         `db:"city" json:"city"`
	Info       null.String    `db:"info" json:"info"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// Customer contains data for table "customer".
type Customer struct {
	Id         int            `db:"id" json:"id"`
	Name       string         `db:"name" json:"name"`
	Identifier string         `db:"identifier" json:"identifier"`
	Phone      string         `db:"phone" json:"phone"`
	Address    string         `db:"address" json:"address"`
	City       string         `db:"city" json:"city"`
	Info       null.String    `db:"info" json:"info"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// ExpenseCategory contains data for table "expense_category".
type ExpenseCategory struct {
	Id         int            `db:"id" json:"id"`
	ParentId   null.Int       `db:"parent_id" json:"parentId"`
	Name       string         `db:"name" json:"name"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// ProductCategory contains data for table "product_category".
type ProductCategory struct {
	Id         int            `db:"id" json:"id"`
	ParentId   null.Int       `db:"parent_id" json:"parentId"`
	Name       string         `db:"name" json:"name"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// Expense contains data for table "expense".
type Expense struct {
	Id          int             `db:"id" json:"id"`
	AccountId   null.Int        `db:"account_id" json:"accountId"`
	EmployeeId  null.Int        `db:"employee_id" json:"employeeId"`
	StoreId     null.Int        `db:"store_id" json:"storeId"`
	CategoryId  null.Int        `db:"category_id" json:"categoryId"`
	Identifier  null.String     `db:"identifier" json:"identifier"`
	Name        string          `db:"name" json:"name"`
	ExpenseDate time.Time       `db:"expense_date" json:"expenseDate"`
	Amount      decimal.Decimal `db:"amount" json:"amount"`
	Deleted     bool            `db:"deleted" json:"deleted"`
	DeleteTime  mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime  time.Time       `db:"update_time" json:"updateTime"`
}

// Purchase contains data for table "purchase".
type Purchase struct {
	Id            int                 `db:"id" json:"id"`
	AccountId     null.Int            `db:"account_id" json:"accountId"`
	SupplierId    int                 `db:"supplier_id" json:"supplierId"`
	ReceiptDate   time.Time           `db:"receipt_date" json:"receiptDate"`
	InputTime     time.Time           `db:"input_time" json:"inputTime"`
	Qty           decimal.Decimal     `db:"qty" json:"qty"`
	TotalPurchase decimal.Decimal     `db:"total_purchase" json:"totalPurchase"`
	PaymentDue    mysql.NullTime      `db:"payment_due" json:"paymentDue"`
	PaidAmount    decimal.NullDecimal `db:"paid_amount" json:"paidAmount"`
	PaidDate      mysql.NullTime      `db:"paid_date" json:"paidDate"`
	Deleted       bool                `db:"deleted" json:"deleted"`
	DeleteTime    mysql.NullTime      `db:"delete_time" json:"deleteTime"`
	UpdateTime    time.Time           `db:"update_time" json:"updateTime"`
}

// Product contains data for table "product".
type Product struct {
	Id         int             `db:"id" json:"id"`
	PurchaseId null.Int        `db:"purchase_id" json:"purchaseId"`
	CategoryId int             `db:"category_id" json:"categoryId"`
	Identifier null.String     `db:"identifier" json:"identifier"`
	Name       string          `db:"name" json:"name"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	Specs      json.RawMessage `db:"specs" json:"specs"`
	Deleted    bool            `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// PurchaseDetail contains data for table "purchase_detail".
type PurchaseDetail struct {
	PurchaseId int             `db:"purchase_id" json:"purchaseId"`
	ProductId  int             `db:"product_id" json:"productId"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	Position   int             `db:"position" json:"position"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// Event contains data for table "event".
type Event struct {
	Id         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	Percentage decimal.Decimal `db:"percentage" json:"percentage"`
	Rounding   decimal.Decimal `db:"rounding" json:"rounding"`
	FixedPrice decimal.Decimal `db:"fixed_price" json:"fixedPrice"`
	Start      time.Time       `db:"start" json:"start"`
	End        mysql.NullTime  `db:"end" json:"end"`
	Deleted    bool            `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// EventDetail contains data for table "event_detail".
type EventDetail struct {
	EventId    int             `db:"event_id" json:"eventId"`
	ProductId  int             `db:"product_id" json:"productId"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	Deleted    bool            `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// PaymentMethod contains data for table "payment_method".
type PaymentMethod struct {
	Id         int            `db:"id" json:"id"`
	Name       string         `db:"name" json:"name"`
	Deleted    bool           `db:"deleted" json:"deleted"`
	DeleteTime mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime time.Time      `db:"update_time" json:"updateTime"`
}

// Display contains data for table "display".
type Display struct {
	Id           int             `db:"id" json:"id"`
	EmployeeId   null.Int        `db:"employee_id" json:"employeeId"`
	StoreId      null.Int        `db:"store_id" json:"storeId"`
	Identifier   null.String     `db:"identifier" json:"identifier"`
	TransferTime time.Time       `db:"transfer_time" json:"transferTime"`
	Qty          decimal.Decimal `db:"qty" json:"qty"`
	Deleted      bool            `db:"deleted" json:"deleted"`
	DeleteTime   mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime   time.Time       `db:"update_time" json:"updateTime"`
}

// DisplayDetail contains data for table "display_detail".
type DisplayDetail struct {
	DisplayId  int             `db:"display_id" json:"displayId"`
	ProductId  int             `db:"product_id" json:"productId"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// Sales contains data for table "sales".
type Sales struct {
	Id              int                 `db:"id" json:"id"`
	EmployeeId      null.Int            `db:"employee_id" json:"employeeId"`
	CustomerId      null.Int            `db:"customer_id" json:"customerId"`
	StoreId         null.Int            `db:"store_id" json:"storeId"`
	PaymentMethodId null.Int            `db:"payment_method_id" json:"paymentMethodId"`
	Identifier      null.String         `db:"identifier" json:"identifier"`
	TransactionTime time.Time           `db:"transaction_time" json:"transactionTime"`
	Qty             decimal.Decimal     `db:"qty" json:"qty"`
	TotalCapital    decimal.Decimal     `db:"total_capital" json:"totalCapital"`
	TotalSales      decimal.Decimal     `db:"total_sales" json:"totalSales"`
	PaymentDue      mysql.NullTime      `db:"payment_due" json:"paymentDue"`
	PaidAmount      decimal.NullDecimal `db:"paid_amount" json:"paidAmount"`
	PaidDate        mysql.NullTime      `db:"paid_date" json:"paidDate"`
	Deleted         bool                `db:"deleted" json:"deleted"`
	DeleteTime      mysql.NullTime      `db:"delete_time" json:"deleteTime"`
	UpdateTime      time.Time           `db:"update_time" json:"updateTime"`
}

// SalesDetail contains data for table "sales_detail".
type SalesDetail struct {
	SalesId    int             `db:"sales_id" json:"salesId"`
	ProductId  int             `db:"product_id" json:"productId"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	UpdateTime time.Time       `db:"update_time" json:"updateTime"`
}

// Correction contains data for table "correction".
type Correction struct {
	Id             int            `db:"id" json:"id"`
	AccountId      null.Int       `db:"account_id" json:"accountId"`
	Description    null.String    `db:"description" json:"description"`
	CorrectionTime time.Time      `db:"correction_time" json:"correctionTime"`
	Deleted        bool           `db:"deleted" json:"deleted"`
	DeleteTime     mysql.NullTime `db:"delete_time" json:"deleteTime"`
	UpdateTime     time.Time      `db:"update_time" json:"updateTime"`
}

// CorrectionDetail contains data for table "correction_detail".
type CorrectionDetail struct {
	CorrectionId int             `db:"correction_id" json:"correctionId"`
	ProductId    int             `db:"product_id" json:"productId"`
	Qty          decimal.Decimal `db:"qty" json:"qty"`
	Deleted      bool            `db:"deleted" json:"deleted"`
	DeleteTime   mysql.NullTime  `db:"delete_time" json:"deleteTime"`
	UpdateTime   time.Time       `db:"update_time" json:"updateTime"`
}

// Structs for SELECT queries.

// LoadAccountResult contains result of select query "LoadAccount".
type LoadAccountResult struct {
	Id         int             `db:"id" json:"id"`
	Username   string          `db:"username" json:"username"`
	Name       string          `db:"name" json:"name"`
	Permission json.RawMessage `db:"permission" json:"permission"`
}

// GetAccountByIdResult contains result of select query "GetAccountById".
type GetAccountByIdResult struct {
	Id         int             `db:"id" json:"id"`
	Username   string          `db:"username" json:"username"`
	Name       string          `db:"name" json:"name"`
	Permission json.RawMessage `db:"permission" json:"permission"`
}

// GetAccountByIdArgs is input parameter for select query "GetAccountById".
type GetAccountByIdArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// GetAccountByUsernameResult contains result of select query "GetAccountByUsername".
type GetAccountByUsernameResult struct {
	Id         int             `db:"id" json:"id"`
	Username   string          `db:"username" json:"username"`
	Name       string          `db:"name" json:"name"`
	Permission json.RawMessage `db:"permission" json:"permission"`
}

// GetAccountByUsernameArgs is input parameter for select query "GetAccountByUsername".
type GetAccountByUsernameArgs struct {
	Username interface{} `db:"username" json:"username"`
}

// GetAdminCountResult contains result of select query "GetAdminCount".
type GetAdminCountResult struct {
	COUNTid int `db:"COUNT(id)" json:"cOUNTid"`
}

// GetCorrectionDetailResult contains result of select query "GetCorrectionDetail".
type GetCorrectionDetailResult struct {
	Id         int             `db:"id" json:"id"`
	Identifier string          `db:"identifier" json:"identifier"`
	Name       string          `db:"name" json:"name"`
	Category   string          `db:"category" json:"category"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
}

// GetCorrectionDetailArgs is input parameter for select query "GetCorrectionDetail".
type GetCorrectionDetailArgs struct {
	CorrectionId interface{} `db:"correction_id" json:"correctionId"`
}

// LoadCorrectionResult contains result of select query "LoadCorrection".
type LoadCorrectionResult struct {
	Id             int             `db:"id" json:"id"`
	AccountId      int             `db:"account_id" json:"accountId"`
	Account        string          `db:"account" json:"account"`
	Description    string          `db:"description" json:"description"`
	CorrectionTime time.Time       `db:"correction_time" json:"correctionTime"`
	Qty            decimal.Decimal `db:"qty" json:"qty"`
	NProducts      int             `db:"n_products" json:"nProducts"`
}

// LoadCorrectionArgs is input parameter for select query "LoadCorrection".
type LoadCorrectionArgs struct {
	Timezone interface{} `db:"timezone" json:"timezone"`
}

// GetCorrectionByIdResult contains result of select query "GetCorrectionById".
type GetCorrectionByIdResult struct {
	Id             int       `db:"id" json:"id"`
	AccountId      int       `db:"account_id" json:"accountId"`
	Description    string    `db:"description" json:"description"`
	CorrectionTime time.Time `db:"correction_time" json:"correctionTime"`
}

// GetCorrectionByIdArgs is input parameter for select query "GetCorrectionById".
type GetCorrectionByIdArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// LoadCustomerResult contains result of select query "LoadCustomer".
type LoadCustomerResult struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Identifier string `db:"identifier" json:"identifier"`
	Phone      string `db:"phone" json:"phone"`
	Address    string `db:"address" json:"address"`
	City       string `db:"city" json:"city"`
	Info       string `db:"info" json:"info"`
}

// LoadDeviceResult contains result of select query "LoadDevice".
type LoadDeviceResult struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	StoreId    int    `db:"store_id" json:"storeId"`
	Store      string `db:"store" json:"store"`
	Identifier string `db:"identifier" json:"identifier"`
	Allowed    bool   `db:"allowed" json:"allowed"`
}

// GetDeviceByIdentifierResult contains result of select query "GetDeviceByIdentifier".
type GetDeviceByIdentifierResult struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	StoreId    int    `db:"store_id" json:"storeId"`
	Store      string `db:"store" json:"store"`
	Identifier string `db:"identifier" json:"identifier"`
	Allowed    bool   `db:"allowed" json:"allowed"`
}

// GetDeviceByIdentifierArgs is input parameter for select query "GetDeviceByIdentifier".
type GetDeviceByIdentifierArgs struct {
	Identifier interface{} `db:"identifier" json:"identifier"`
}

// LoadEmployeeResult contains result of select query "LoadEmployee".
type LoadEmployeeResult struct {
	Id        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Name      string `db:"name" json:"name"`
	Gender    string `db:"gender" json:"gender"`
	Religion  string `db:"religion" json:"religion"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
	City      string `db:"city" json:"city"`
	StartYear int    `db:"start_year" json:"startYear"`
	EndYear   int    `db:"end_year" json:"endYear"`
}

// LoadProductInEventResult contains result of select query "LoadProductInEvent".
type LoadProductInEventResult struct {
	Id         int             `db:"id" json:"id"`
	Identifier string          `db:"identifier" json:"identifier"`
	Name       string          `db:"name" json:"name"`
	Category   string          `db:"category" json:"category"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
}

// LoadProductInEventArgs is input parameter for select query "LoadProductInEvent".
type LoadProductInEventArgs struct {
	EventId interface{} `db:"event_id" json:"eventId"`
}

// LoadEventResult contains result of select query "LoadEvent".
type LoadEventResult struct {
	Id         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	Percentage decimal.Decimal `db:"percentage" json:"percentage"`
	Rounding   decimal.Decimal `db:"rounding" json:"rounding"`
	FixedPrice decimal.Decimal `db:"fixed_price" json:"fixedPrice"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
	NProducts  int             `db:"n_products" json:"nProducts"`
	Start      time.Time       `db:"start" json:"start"`
	End        time.Time       `db:"end" json:"end"`
}

// LoadEventArgs is input parameter for select query "LoadEvent".
type LoadEventArgs struct {
	Timezone interface{} `db:"timezone" json:"timezone"`
}

// GetEventResult contains result of select query "GetEvent".
type GetEventResult struct {
	Id         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	Percentage decimal.Decimal `db:"percentage" json:"percentage"`
	Rounding   decimal.Decimal `db:"rounding" json:"rounding"`
	FixedPrice decimal.Decimal `db:"fixed_price" json:"fixedPrice"`
	Start      time.Time       `db:"start" json:"start"`
	End        time.Time       `db:"end" json:"end"`
}

// GetEventArgs is input parameter for select query "GetEvent".
type GetEventArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// GetProductActiveEventResult contains result of select query "GetProductActiveEvent".
type GetProductActiveEventResult struct {
	Id         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	Percentage decimal.Decimal `db:"percentage" json:"percentage"`
	Rounding   decimal.Decimal `db:"rounding" json:"rounding"`
	FixedPrice decimal.Decimal `db:"fixed_price" json:"fixedPrice"`
	Start      time.Time       `db:"start" json:"start"`
	End        time.Time       `db:"end" json:"end"`
}

// GetProductActiveEventArgs is input parameter for select query "GetProductActiveEvent".
type GetProductActiveEventArgs struct {
	ProductId interface{} `db:"product_id" json:"productId"`
}

// GetProductPastEventResult contains result of select query "GetProductPastEvent".
type GetProductPastEventResult struct {
	Id         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	Percentage decimal.Decimal `db:"percentage" json:"percentage"`
	Rounding   decimal.Decimal `db:"rounding" json:"rounding"`
	FixedPrice decimal.Decimal `db:"fixed_price" json:"fixedPrice"`
	Start      time.Time       `db:"start" json:"start"`
	End        time.Time       `db:"end" json:"end"`
}

// GetProductPastEventArgs is input parameter for select query "GetProductPastEvent".
type GetProductPastEventArgs struct {
	ProductId interface{} `db:"product_id" json:"productId"`
	EventId   interface{} `db:"event_id" json:"eventId"`
	StartTime interface{} `db:"start_time" json:"startTime"`
	EndTime   interface{} `db:"end_time" json:"endTime"`
}

// LoadExpenseCategoryResult contains result of select query "LoadExpenseCategory".
type LoadExpenseCategoryResult struct {
	Value   int    `db:"value" json:"value"`
	Caption string `db:"caption" json:"caption"`
}

// LoadExpenseCategoryArgs is input parameter for select query "LoadExpenseCategory".
type LoadExpenseCategoryArgs struct {
	ParentId interface{} `db:"parent_id" json:"parentId"`
}

// GetExpenseCategoryByIdResult contains result of select query "GetExpenseCategoryById".
type GetExpenseCategoryByIdResult struct {
	Id       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ParentId int    `db:"parent_id" json:"parentId"`
}

// GetExpenseCategoryByIdArgs is input parameter for select query "GetExpenseCategoryById".
type GetExpenseCategoryByIdArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// LoadExpenseResult contains result of select query "LoadExpense".
type LoadExpenseResult struct {
	Id         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	CategoryId int             `db:"category_id" json:"categoryId"`
	Category   string          `db:"category" json:"category"`
	Amount     decimal.Decimal `db:"amount" json:"amount"`
}

// LoadExpenseArgs is input parameter for select query "LoadExpense".
type LoadExpenseArgs struct {
	ExpenseDate interface{} `db:"expense_date" json:"expenseDate"`
}

// GetSumExpenseResult contains result of select query "GetSumExpense".
type GetSumExpenseResult struct {
	Amount decimal.Decimal `db:"amount" json:"amount"`
}

// GetSumExpenseArgs is input parameter for select query "GetSumExpense".
type GetSumExpenseArgs struct {
	ExpenseDate interface{} `db:"expense_date" json:"expenseDate"`
}

// LoadPaymentMethodResult contains result of select query "LoadPaymentMethod".
type LoadPaymentMethodResult struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// LoadProductCategoryResult contains result of select query "LoadProductCategory".
type LoadProductCategoryResult struct {
	Value   int    `db:"value" json:"value"`
	Caption string `db:"caption" json:"caption"`
}

// LoadProductCategoryArgs is input parameter for select query "LoadProductCategory".
type LoadProductCategoryArgs struct {
	ParentId interface{} `db:"parent_id" json:"parentId"`
}

// GetProductResult contains result of select query "GetProduct".
type GetProductResult struct {
	Id         int             `db:"id" json:"id"`
	PurchaseId int             `db:"purchase_id" json:"purchaseId"`
	CategoryId int             `db:"category_id" json:"categoryId"`
	Identifier string          `db:"identifier" json:"identifier"`
	Name       string          `db:"name" json:"name"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	Specs      json.RawMessage `db:"specs" json:"specs"`
}

// GetProductArgs is input parameter for select query "GetProduct".
type GetProductArgs struct {
	Identifier interface{} `db:"identifier" json:"identifier"`
}

// GetProductDetailResult contains result of select query "GetProductDetail".
type GetProductDetailResult struct {
	Id         int             `db:"id" json:"id"`
	PurchaseId int             `db:"purchase_id" json:"purchaseId"`
	Identifier string          `db:"identifier" json:"identifier"`
	SupplierId int             `db:"supplier_id" json:"supplierId"`
	Supplier   string          `db:"supplier" json:"supplier"`
	CategoryId int             `db:"category_id" json:"categoryId"`
	Category   string          `db:"category" json:"category"`
	Name       string          `db:"name" json:"name"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	Specs      json.RawMessage `db:"specs" json:"specs"`
}

// GetProductDetailArgs is input parameter for select query "GetProductDetail".
type GetProductDetailArgs struct {
	Identifier interface{} `db:"identifier" json:"identifier"`
}

// GetProductLeftoverResult contains result of select query "GetProductLeftover".
type GetProductLeftoverResult struct {
	Id         int             `db:"id" json:"id"`
	Identifier string          `db:"identifier" json:"identifier"`
	Name       string          `db:"name" json:"name"`
	Category   string          `db:"category" json:"category"`
	Capital    decimal.Decimal `db:"capital" json:"capital"`
	Price      decimal.Decimal `db:"price" json:"price"`
	Qty        decimal.Decimal `db:"qty" json:"qty"`
}

// GetProductLeftoverArgs is input parameter for select query "GetProductLeftover".
type GetProductLeftoverArgs struct {
	Identifier interface{} `db:"identifier" json:"identifier"`
}

// LoadAllPurchaseResult contains result of select query "LoadAllPurchase".
type LoadAllPurchaseResult struct {
	Id            int             `db:"id" json:"id"`
	AccountId     int             `db:"account_id" json:"accountId"`
	Account       string          `db:"account" json:"account"`
	SupplierId    int             `db:"supplier_id" json:"supplierId"`
	Supplier      string          `db:"supplier" json:"supplier"`
	Category      string          `db:"category" json:"category"`
	ReceiptDate   time.Time       `db:"receipt_date" json:"receiptDate"`
	InputTime     time.Time       `db:"input_time" json:"inputTime"`
	NProducts     int             `db:"n_products" json:"nProducts"`
	Qty           decimal.Decimal `db:"qty" json:"qty"`
	TotalPurchase decimal.Decimal `db:"total_purchase" json:"totalPurchase"`
	PaymentDue    time.Time       `db:"payment_due" json:"paymentDue"`
	PaidAmount    decimal.Decimal `db:"paid_amount" json:"paidAmount"`
	PaidDate      time.Time       `db:"paid_date" json:"paidDate"`
	PercentProfit decimal.Decimal `db:"percent_profit" json:"percentProfit"`
	Finished      int             `db:"finished" json:"finished"`
}

// LoadAllPurchaseArgs is input parameter for select query "LoadAllPurchase".
type LoadAllPurchaseArgs struct {
	Timezone interface{} `db:"timezone" json:"timezone"`
}

// LoadActivePurchaseResult contains result of select query "LoadActivePurchase".
type LoadActivePurchaseResult struct {
	Id            int             `db:"id" json:"id"`
	AccountId     int             `db:"account_id" json:"accountId"`
	Account       string          `db:"account" json:"account"`
	SupplierId    int             `db:"supplier_id" json:"supplierId"`
	Supplier      string          `db:"supplier" json:"supplier"`
	Category      string          `db:"category" json:"category"`
	ReceiptDate   time.Time       `db:"receipt_date" json:"receiptDate"`
	InputTime     time.Time       `db:"input_time" json:"inputTime"`
	NProducts     int             `db:"n_products" json:"nProducts"`
	Qty           decimal.Decimal `db:"qty" json:"qty"`
	TotalPurchase decimal.Decimal `db:"total_purchase" json:"totalPurchase"`
	Finished      int             `db:"finished" json:"finished"`
}

// LoadActivePurchaseArgs is input parameter for select query "LoadActivePurchase".
type LoadActivePurchaseArgs struct {
	Timezone  interface{} `db:"timezone" json:"timezone"`
	StartDate interface{} `db:"start_date" json:"startDate"`
}

// GetPurchaseResult contains result of select query "GetPurchase".
type GetPurchaseResult struct {
	Id            int             `db:"id" json:"id"`
	AccountId     int             `db:"account_id" json:"accountId"`
	SupplierId    int             `db:"supplier_id" json:"supplierId"`
	ReceiptDate   time.Time       `db:"receipt_date" json:"receiptDate"`
	InputTime     time.Time       `db:"input_time" json:"inputTime"`
	Qty           decimal.Decimal `db:"qty" json:"qty"`
	TotalPurchase decimal.Decimal `db:"total_purchase" json:"totalPurchase"`
	PaymentDue    time.Time       `db:"payment_due" json:"paymentDue"`
	PaidAmount    decimal.Decimal `db:"paid_amount" json:"paidAmount"`
	PaidDate      time.Time       `db:"paid_date" json:"paidDate"`
}

// GetPurchaseArgs is input parameter for select query "GetPurchase".
type GetPurchaseArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// GetPurchaseDetailResult contains result of select query "GetPurchaseDetail".
type GetPurchaseDetailResult struct {
	Id           int             `db:"id" json:"id"`
	Identifier   string          `db:"identifier" json:"identifier"`
	Name         string          `db:"name" json:"name"`
	Category     string          `db:"category" json:"category"`
	Qty          decimal.Decimal `db:"qty" json:"qty"`
	Capital      decimal.Decimal `db:"capital" json:"capital"`
	Price        decimal.Decimal `db:"price" json:"price"`
	TotalCapital decimal.Decimal `db:"total_capital" json:"totalCapital"`
}

// GetPurchaseDetailArgs is input parameter for select query "GetPurchaseDetail".
type GetPurchaseDetailArgs struct {
	PurchaseId interface{} `db:"purchase_id" json:"purchaseId"`
}

// GetPurchaseDetailSumResult contains result of select query "GetPurchaseDetailSum".
type GetPurchaseDetailSumResult struct {
	Qty          decimal.Decimal `db:"qty" json:"qty"`
	TotalCapital decimal.Decimal `db:"total_capital" json:"totalCapital"`
}

// GetPurchaseDetailSumArgs is input parameter for select query "GetPurchaseDetailSum".
type GetPurchaseDetailSumArgs struct {
	PurchaseId interface{} `db:"purchase_id" json:"purchaseId"`
}

// LoadStoreResult contains result of select query "LoadStore".
type LoadStoreResult struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Phone   string `db:"phone" json:"phone"`
	Address string `db:"address" json:"address"`
	City    string `db:"city" json:"city"`
	Info    string `db:"info" json:"info"`
}

// GetStoreByIdResult contains result of select query "GetStoreById".
type GetStoreByIdResult struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Phone   string `db:"phone" json:"phone"`
	Address string `db:"address" json:"address"`
	City    string `db:"city" json:"city"`
	Info    string `db:"info" json:"info"`
}

// GetStoreByIdArgs is input parameter for select query "GetStoreById".
type GetStoreByIdArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// LoadSupplierResult contains result of select query "LoadSupplier".
type LoadSupplierResult struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Phone      string `db:"phone" json:"phone"`
	Address    string `db:"address" json:"address"`
	City       string `db:"city" json:"city"`
	Info       string `db:"info" json:"info"`
	CategoryId int    `db:"category_id" json:"categoryId"`
	Category   string `db:"category" json:"category"`
}

// Structs for EXEC queries.

// InsertAccountArgs is input parameter for exec query "InsertAccount".
type InsertAccountArgs struct {
	Username interface{} `db:"username" json:"username"`
	Name     interface{} `db:"name" json:"name"`
}

// DeleteAccountArgs is input parameter for exec query "DeleteAccount".
type DeleteAccountArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateAccountArgs is input parameter for exec query "UpdateAccount".
type UpdateAccountArgs struct {
	Username   interface{} `db:"username" json:"username"`
	Name       interface{} `db:"name" json:"name"`
	Permission interface{} `db:"permission" json:"permission"`
	Id         interface{} `db:"id" json:"id"`
}

// UpdateAccountPasswordArgs is input parameter for exec query "UpdateAccountPassword".
type UpdateAccountPasswordArgs struct {
	Password interface{} `db:"password" json:"password"`
	Id       interface{} `db:"id" json:"id"`
}

// InsertCorrectionDetailArgs is input parameter for exec query "InsertCorrectionDetail".
type InsertCorrectionDetailArgs struct {
	CorrectionId interface{} `db:"correction_id" json:"correctionId"`
	ProductId    interface{} `db:"product_id" json:"productId"`
	Qty          interface{} `db:"qty" json:"qty"`
}

// DeleteCorrectionDetailByCorrectionIdArgs is input parameter for exec query "DeleteCorrectionDetailByCorrectionId".
type DeleteCorrectionDetailByCorrectionIdArgs struct {
	CorrectionId interface{} `db:"correction_id" json:"correctionId"`
}

// DeleteCorrectionDetailByProductArgs is input parameter for exec query "DeleteCorrectionDetailByProduct".
type DeleteCorrectionDetailByProductArgs struct {
	CorrectionId interface{} `db:"correction_id" json:"correctionId"`
	Identifier   interface{} `db:"identifier" json:"identifier"`
}

// InsertCorrectionArgs is input parameter for exec query "InsertCorrection".
type InsertCorrectionArgs struct {
	AccountId      interface{} `db:"account_id" json:"accountId"`
	Description    interface{} `db:"description" json:"description"`
	CorrectionTime interface{} `db:"correction_time" json:"correctionTime"`
}

// DeleteCorrectionArgs is input parameter for exec query "DeleteCorrection".
type DeleteCorrectionArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateCorrectionArgs is input parameter for exec query "UpdateCorrection".
type UpdateCorrectionArgs struct {
	Description interface{} `db:"description" json:"description"`
	Id          interface{} `db:"id" json:"id"`
}

// InsertCustomerArgs is input parameter for exec query "InsertCustomer".
type InsertCustomerArgs struct {
	Name       interface{} `db:"name" json:"name"`
	Identifier interface{} `db:"identifier" json:"identifier"`
	Phone      interface{} `db:"phone" json:"phone"`
	Address    interface{} `db:"address" json:"address"`
	City       interface{} `db:"city" json:"city"`
	Info       interface{} `db:"info" json:"info"`
}

// DeleteCustomerArgs is input parameter for exec query "DeleteCustomer".
type DeleteCustomerArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateCustomerArgs is input parameter for exec query "UpdateCustomer".
type UpdateCustomerArgs struct {
	Name    interface{} `db:"name" json:"name"`
	Phone   interface{} `db:"phone" json:"phone"`
	Address interface{} `db:"address" json:"address"`
	City    interface{} `db:"city" json:"city"`
	Info    interface{} `db:"info" json:"info"`
	Id      interface{} `db:"id" json:"id"`
}

// InsertDeviceArgs is input parameter for exec query "InsertDevice".
type InsertDeviceArgs struct {
	Name       interface{} `db:"name" json:"name"`
	Identifier interface{} `db:"identifier" json:"identifier"`
}

// DeleteDeviceArgs is input parameter for exec query "DeleteDevice".
type DeleteDeviceArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateDeviceArgs is input parameter for exec query "UpdateDevice".
type UpdateDeviceArgs struct {
	Name    interface{} `db:"name" json:"name"`
	Allowed interface{} `db:"allowed" json:"allowed"`
	Id      interface{} `db:"id" json:"id"`
}

// SetDeviceStoreArgs is input parameter for exec query "SetDeviceStore".
type SetDeviceStoreArgs struct {
	StoreId interface{} `db:"store_id" json:"storeId"`
	Id      interface{} `db:"id" json:"id"`
}

// MarkDeviceAllowedArgs is input parameter for exec query "MarkDeviceAllowed".
type MarkDeviceAllowedArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// InsertEmployeeArgs is input parameter for exec query "InsertEmployee".
type InsertEmployeeArgs struct {
	Username  interface{} `db:"username" json:"username"`
	Name      interface{} `db:"name" json:"name"`
	Gender    interface{} `db:"gender" json:"gender"`
	Religion  interface{} `db:"religion" json:"religion"`
	Phone     interface{} `db:"phone" json:"phone"`
	Address   interface{} `db:"address" json:"address"`
	City      interface{} `db:"city" json:"city"`
	StartYear interface{} `db:"start_year" json:"startYear"`
	EndYear   interface{} `db:"end_year" json:"endYear"`
}

// DeleteEmployeeArgs is input parameter for exec query "DeleteEmployee".
type DeleteEmployeeArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateEmployeeArgs is input parameter for exec query "UpdateEmployee".
type UpdateEmployeeArgs struct {
	Username  interface{} `db:"username" json:"username"`
	Name      interface{} `db:"name" json:"name"`
	Gender    interface{} `db:"gender" json:"gender"`
	Religion  interface{} `db:"religion" json:"religion"`
	Phone     interface{} `db:"phone" json:"phone"`
	Address   interface{} `db:"address" json:"address"`
	City      interface{} `db:"city" json:"city"`
	StartYear interface{} `db:"start_year" json:"startYear"`
	EndYear   interface{} `db:"end_year" json:"endYear"`
	Id        interface{} `db:"id" json:"id"`
}

// InsertEventDetailArgs is input parameter for exec query "InsertEventDetail".
type InsertEventDetailArgs struct {
	EventId   interface{} `db:"event_id" json:"eventId"`
	ProductId interface{} `db:"product_id" json:"productId"`
	Qty       interface{} `db:"qty" json:"qty"`
	Capital   interface{} `db:"capital" json:"capital"`
	Price     interface{} `db:"price" json:"price"`
}

// UpdateEventDetailQtyArgs is input parameter for exec query "UpdateEventDetailQty".
type UpdateEventDetailQtyArgs struct {
	Qty       interface{} `db:"qty" json:"qty"`
	ProductId interface{} `db:"product_id" json:"productId"`
	EventId   interface{} `db:"event_id" json:"eventId"`
}

// UpdateEventDetailPriceArgs is input parameter for exec query "UpdateEventDetailPrice".
type UpdateEventDetailPriceArgs struct {
	ProductId interface{} `db:"product_id" json:"productId"`
	EventId   interface{} `db:"event_id" json:"eventId"`
}

// DeleteEventDetailByEventIdArgs is input parameter for exec query "DeleteEventDetailByEventId".
type DeleteEventDetailByEventIdArgs struct {
	EventId interface{} `db:"event_id" json:"eventId"`
}

// DeleteEventProductArgs is input parameter for exec query "DeleteEventProduct".
type DeleteEventProductArgs struct {
	EventId    interface{} `db:"event_id" json:"eventId"`
	Identifier interface{} `db:"identifier" json:"identifier"`
}

// CreateEventArgs is input parameter for exec query "CreateEvent".
type CreateEventArgs struct {
	Name       interface{} `db:"name" json:"name"`
	Percentage interface{} `db:"percentage" json:"percentage"`
	Rounding   interface{} `db:"rounding" json:"rounding"`
	FixedPrice interface{} `db:"fixed_price" json:"fixedPrice"`
	Start      interface{} `db:"start" json:"start"`
	End        interface{} `db:"end" json:"end"`
}

// DeleteEventArgs is input parameter for exec query "DeleteEvent".
type DeleteEventArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateEventArgs is input parameter for exec query "UpdateEvent".
type UpdateEventArgs struct {
	Name       interface{} `db:"name" json:"name"`
	Percentage interface{} `db:"percentage" json:"percentage"`
	Rounding   interface{} `db:"rounding" json:"rounding"`
	FixedPrice interface{} `db:"fixed_price" json:"fixedPrice"`
	Start      interface{} `db:"start" json:"start"`
	End        interface{} `db:"end" json:"end"`
	Id         interface{} `db:"id" json:"id"`
}

// InsertExpenseCategoryArgs is input parameter for exec query "InsertExpenseCategory".
type InsertExpenseCategoryArgs struct {
	ParentId interface{} `db:"parent_id" json:"parentId"`
	Name     interface{} `db:"name" json:"name"`
}

// DeleteExpenseCategoryArgs is input parameter for exec query "DeleteExpenseCategory".
type DeleteExpenseCategoryArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateExpenseCategoryArgs is input parameter for exec query "UpdateExpenseCategory".
type UpdateExpenseCategoryArgs struct {
	Name interface{} `db:"name" json:"name"`
	Id   interface{} `db:"id" json:"id"`
}

// InsertExpenseArgs is input parameter for exec query "InsertExpense".
type InsertExpenseArgs struct {
	AccountId   interface{} `db:"account_id" json:"accountId"`
	CategoryId  interface{} `db:"category_id" json:"categoryId"`
	Identifier  interface{} `db:"identifier" json:"identifier"`
	Name        interface{} `db:"name" json:"name"`
	ExpenseDate interface{} `db:"expense_date" json:"expenseDate"`
	Amount      interface{} `db:"amount" json:"amount"`
}

// DeleteExpenseArgs is input parameter for exec query "DeleteExpense".
type DeleteExpenseArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateExpenseArgs is input parameter for exec query "UpdateExpense".
type UpdateExpenseArgs struct {
	AccountId  interface{} `db:"account_id" json:"accountId"`
	CategoryId interface{} `db:"category_id" json:"categoryId"`
	Name       interface{} `db:"name" json:"name"`
	Amount     interface{} `db:"amount" json:"amount"`
	Id         interface{} `db:"id" json:"id"`
}

// InsertPaymentMethodArgs is input parameter for exec query "InsertPaymentMethod".
type InsertPaymentMethodArgs struct {
	Name interface{} `db:"name" json:"name"`
}

// DeletePaymentMethodArgs is input parameter for exec query "DeletePaymentMethod".
type DeletePaymentMethodArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdatePaymentMethodArgs is input parameter for exec query "UpdatePaymentMethod".
type UpdatePaymentMethodArgs struct {
	Name interface{} `db:"name" json:"name"`
	Id   interface{} `db:"id" json:"id"`
}

// InsertProductCategoryArgs is input parameter for exec query "InsertProductCategory".
type InsertProductCategoryArgs struct {
	ParentId interface{} `db:"parent_id" json:"parentId"`
	Name     interface{} `db:"name" json:"name"`
}

// DeleteProductCategoryArgs is input parameter for exec query "DeleteProductCategory".
type DeleteProductCategoryArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateProductCategoryArgs is input parameter for exec query "UpdateProductCategory".
type UpdateProductCategoryArgs struct {
	Name interface{} `db:"name" json:"name"`
	Id   interface{} `db:"id" json:"id"`
}

// InsertProductArgs is input parameter for exec query "InsertProduct".
type InsertProductArgs struct {
	PurchaseId interface{} `db:"purchase_id" json:"purchaseId"`
	CategoryId interface{} `db:"category_id" json:"categoryId"`
	Name       interface{} `db:"name" json:"name"`
	Qty        interface{} `db:"qty" json:"qty"`
	Capital    interface{} `db:"capital" json:"capital"`
	Price      interface{} `db:"price" json:"price"`
	Specs      interface{} `db:"specs" json:"specs"`
}

// SetProductIdentifierArgs is input parameter for exec query "SetProductIdentifier".
type SetProductIdentifierArgs struct {
	PurchaseId interface{} `db:"purchase_id" json:"purchaseId"`
}

// InsertPurchaseArgs is input parameter for exec query "InsertPurchase".
type InsertPurchaseArgs struct {
	AccountId     interface{} `db:"account_id" json:"accountId"`
	SupplierId    interface{} `db:"supplier_id" json:"supplierId"`
	ReceiptDate   interface{} `db:"receipt_date" json:"receiptDate"`
	InputTime     interface{} `db:"input_time" json:"inputTime"`
	Qty           interface{} `db:"qty" json:"qty"`
	TotalPurchase interface{} `db:"total_purchase" json:"totalPurchase"`
}

// UpdatePurchaseArgs is input parameter for exec query "UpdatePurchase".
type UpdatePurchaseArgs struct {
	ReceiptDate interface{} `db:"receipt_date" json:"receiptDate"`
	PaymentDue  interface{} `db:"payment_due" json:"paymentDue"`
	PaidAmount  interface{} `db:"paid_amount" json:"paidAmount"`
	PaidDate    interface{} `db:"paid_date" json:"paidDate"`
	Id          interface{} `db:"id" json:"id"`
}

// UpdatePurchasePaymentArgs is input parameter for exec query "UpdatePurchasePayment".
type UpdatePurchasePaymentArgs struct {
	PaidAmount interface{} `db:"paid_amount" json:"paidAmount"`
	PaidDate   interface{} `db:"paid_date" json:"paidDate"`
	Id         interface{} `db:"id" json:"id"`
}

// InsertPurchaseDetailArgs is input parameter for exec query "InsertPurchaseDetail".
type InsertPurchaseDetailArgs struct {
	PurchaseId interface{} `db:"purchase_id" json:"purchaseId"`
	ProductId  interface{} `db:"product_id" json:"productId"`
	Qty        interface{} `db:"qty" json:"qty"`
	Capital    interface{} `db:"capital" json:"capital"`
	Price      interface{} `db:"price" json:"price"`
	Position   interface{} `db:"position" json:"position"`
}

// InsertStoreArgs is input parameter for exec query "InsertStore".
type InsertStoreArgs struct {
	Name    interface{} `db:"name" json:"name"`
	Phone   interface{} `db:"phone" json:"phone"`
	Address interface{} `db:"address" json:"address"`
	City    interface{} `db:"city" json:"city"`
	Info    interface{} `db:"info" json:"info"`
}

// DeleteStoreArgs is input parameter for exec query "DeleteStore".
type DeleteStoreArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateStoreArgs is input parameter for exec query "UpdateStore".
type UpdateStoreArgs struct {
	Name    interface{} `db:"name" json:"name"`
	Phone   interface{} `db:"phone" json:"phone"`
	Address interface{} `db:"address" json:"address"`
	City    interface{} `db:"city" json:"city"`
	Info    interface{} `db:"info" json:"info"`
	Id      interface{} `db:"id" json:"id"`
}

// InsertSupplierArgs is input parameter for exec query "InsertSupplier".
type InsertSupplierArgs struct {
	Name    interface{} `db:"name" json:"name"`
	Phone   interface{} `db:"phone" json:"phone"`
	Address interface{} `db:"address" json:"address"`
	City    interface{} `db:"city" json:"city"`
	Info    interface{} `db:"info" json:"info"`
}

// DeleteSupplierArgs is input parameter for exec query "DeleteSupplier".
type DeleteSupplierArgs struct {
	Id interface{} `db:"id" json:"id"`
}

// UpdateSupplierArgs is input parameter for exec query "UpdateSupplier".
type UpdateSupplierArgs struct {
	Name    interface{} `db:"name" json:"name"`
	Phone   interface{} `db:"phone" json:"phone"`
	Address interface{} `db:"address" json:"address"`
	City    interface{} `db:"city" json:"city"`
	Info    interface{} `db:"info" json:"info"`
	Id      interface{} `db:"id" json:"id"`
}
