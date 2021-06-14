package storage

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

// LoadAccount is wrapper for select query "LoadAccount".
func LoadAccount(acc Accessor, additionalQueries ...string) (res []LoadAccountResult, err error) {
	query := sqlLoadAccount

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetAccountById is wrapper for select query "GetAccountById".
func GetAccountById(acc Accessor, namedArgs GetAccountByIdArgs, additionalQueries ...string) (res GetAccountByIdResult, err error) {
	query, args, err := sqlx.Named(sqlGetAccountById, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetAccountByUsername is wrapper for select query "GetAccountByUsername".
func GetAccountByUsername(acc Accessor, namedArgs GetAccountByUsernameArgs, additionalQueries ...string) (res GetAccountByUsernameResult, err error) {
	query, args, err := sqlx.Named(sqlGetAccountByUsername, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetAdminCount is wrapper for select query "GetAdminCount".
func GetAdminCount(acc Accessor, additionalQueries ...string) (res GetAdminCountResult, err error) {
	query := sqlGetAdminCount

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetCorrectionDetail is wrapper for select query "GetCorrectionDetail".
func GetCorrectionDetail(acc Accessor, namedArgs GetCorrectionDetailArgs, additionalQueries ...string) (res []GetCorrectionDetailResult, err error) {
	query, args, err := sqlx.Named(sqlGetCorrectionDetail, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadCorrection is wrapper for select query "LoadCorrection".
func LoadCorrection(acc Accessor, namedArgs LoadCorrectionArgs, additionalQueries ...string) (res []LoadCorrectionResult, err error) {
	query, args, err := sqlx.Named(sqlLoadCorrection, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetCorrectionById is wrapper for select query "GetCorrectionById".
func GetCorrectionById(acc Accessor, namedArgs GetCorrectionByIdArgs, additionalQueries ...string) (res GetCorrectionByIdResult, err error) {
	query, args, err := sqlx.Named(sqlGetCorrectionById, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadCustomer is wrapper for select query "LoadCustomer".
func LoadCustomer(acc Accessor, additionalQueries ...string) (res []LoadCustomerResult, err error) {
	query := sqlLoadCustomer

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadDevice is wrapper for select query "LoadDevice".
func LoadDevice(acc Accessor, additionalQueries ...string) (res []LoadDeviceResult, err error) {
	query := sqlLoadDevice

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetDeviceByIdentifier is wrapper for select query "GetDeviceByIdentifier".
func GetDeviceByIdentifier(acc Accessor, namedArgs GetDeviceByIdentifierArgs, additionalQueries ...string) (res GetDeviceByIdentifierResult, err error) {
	query, args, err := sqlx.Named(sqlGetDeviceByIdentifier, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadEmployee is wrapper for select query "LoadEmployee".
func LoadEmployee(acc Accessor, additionalQueries ...string) (res []LoadEmployeeResult, err error) {
	query := sqlLoadEmployee

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadProductInEvent is wrapper for select query "LoadProductInEvent".
func LoadProductInEvent(acc Accessor, namedArgs LoadProductInEventArgs, additionalQueries ...string) (res []LoadProductInEventResult, err error) {
	query, args, err := sqlx.Named(sqlLoadProductInEvent, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadEvent is wrapper for select query "LoadEvent".
func LoadEvent(acc Accessor, namedArgs LoadEventArgs, additionalQueries ...string) (res []LoadEventResult, err error) {
	query, args, err := sqlx.Named(sqlLoadEvent, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetEvent is wrapper for select query "GetEvent".
func GetEvent(acc Accessor, namedArgs GetEventArgs, additionalQueries ...string) (res GetEventResult, err error) {
	query, args, err := sqlx.Named(sqlGetEvent, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetProductActiveEvent is wrapper for select query "GetProductActiveEvent".
func GetProductActiveEvent(acc Accessor, namedArgs GetProductActiveEventArgs, additionalQueries ...string) (res GetProductActiveEventResult, err error) {
	query, args, err := sqlx.Named(sqlGetProductActiveEvent, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetProductPastEvent is wrapper for select query "GetProductPastEvent".
func GetProductPastEvent(acc Accessor, namedArgs GetProductPastEventArgs, additionalQueries ...string) (res GetProductPastEventResult, err error) {
	query, args, err := sqlx.Named(sqlGetProductPastEvent, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadExpenseCategory is wrapper for select query "LoadExpenseCategory".
func LoadExpenseCategory(acc Accessor, namedArgs LoadExpenseCategoryArgs, additionalQueries ...string) (res []LoadExpenseCategoryResult, err error) {
	query, args, err := sqlx.Named(sqlLoadExpenseCategory, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetExpenseCategoryById is wrapper for select query "GetExpenseCategoryById".
func GetExpenseCategoryById(acc Accessor, namedArgs GetExpenseCategoryByIdArgs, additionalQueries ...string) (res GetExpenseCategoryByIdResult, err error) {
	query, args, err := sqlx.Named(sqlGetExpenseCategoryById, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadExpense is wrapper for select query "LoadExpense".
func LoadExpense(acc Accessor, namedArgs LoadExpenseArgs, additionalQueries ...string) (res []LoadExpenseResult, err error) {
	query, args, err := sqlx.Named(sqlLoadExpense, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetSumExpense is wrapper for select query "GetSumExpense".
func GetSumExpense(acc Accessor, namedArgs GetSumExpenseArgs, additionalQueries ...string) (res GetSumExpenseResult, err error) {
	query, args, err := sqlx.Named(sqlGetSumExpense, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadPaymentMethod is wrapper for select query "LoadPaymentMethod".
func LoadPaymentMethod(acc Accessor, additionalQueries ...string) (res []LoadPaymentMethodResult, err error) {
	query := sqlLoadPaymentMethod

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadProductCategory is wrapper for select query "LoadProductCategory".
func LoadProductCategory(acc Accessor, namedArgs LoadProductCategoryArgs, additionalQueries ...string) (res []LoadProductCategoryResult, err error) {
	query, args, err := sqlx.Named(sqlLoadProductCategory, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetProduct is wrapper for select query "GetProduct".
func GetProduct(acc Accessor, namedArgs GetProductArgs, additionalQueries ...string) (res GetProductResult, err error) {
	query, args, err := sqlx.Named(sqlGetProduct, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetProductDetail is wrapper for select query "GetProductDetail".
func GetProductDetail(acc Accessor, namedArgs GetProductDetailArgs, additionalQueries ...string) (res GetProductDetailResult, err error) {
	query, args, err := sqlx.Named(sqlGetProductDetail, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetProductLeftover is wrapper for select query "GetProductLeftover".
func GetProductLeftover(acc Accessor, namedArgs GetProductLeftoverArgs, additionalQueries ...string) (res GetProductLeftoverResult, err error) {
	query, args, err := sqlx.Named(sqlGetProductLeftover, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadAllPurchase is wrapper for select query "LoadAllPurchase".
func LoadAllPurchase(acc Accessor, namedArgs LoadAllPurchaseArgs, additionalQueries ...string) (res []LoadAllPurchaseResult, err error) {
	query, args, err := sqlx.Named(sqlLoadAllPurchase, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadActivePurchase is wrapper for select query "LoadActivePurchase".
func LoadActivePurchase(acc Accessor, namedArgs LoadActivePurchaseArgs, additionalQueries ...string) (res []LoadActivePurchaseResult, err error) {
	query, args, err := sqlx.Named(sqlLoadActivePurchase, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetPurchase is wrapper for select query "GetPurchase".
func GetPurchase(acc Accessor, namedArgs GetPurchaseArgs, additionalQueries ...string) (res GetPurchaseResult, err error) {
	query, args, err := sqlx.Named(sqlGetPurchase, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetPurchaseDetail is wrapper for select query "GetPurchaseDetail".
func GetPurchaseDetail(acc Accessor, namedArgs GetPurchaseDetailArgs, additionalQueries ...string) (res []GetPurchaseDetailResult, err error) {
	query, args, err := sqlx.Named(sqlGetPurchaseDetail, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetPurchaseDetailSum is wrapper for select query "GetPurchaseDetailSum".
func GetPurchaseDetailSum(acc Accessor, namedArgs GetPurchaseDetailSumArgs, additionalQueries ...string) (res GetPurchaseDetailSumResult, err error) {
	query, args, err := sqlx.Named(sqlGetPurchaseDetailSum, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadStore is wrapper for select query "LoadStore".
func LoadStore(acc Accessor, additionalQueries ...string) (res []LoadStoreResult, err error) {
	query := sqlLoadStore

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// GetStoreById is wrapper for select query "GetStoreById".
func GetStoreById(acc Accessor, namedArgs GetStoreByIdArgs, additionalQueries ...string) (res GetStoreByIdResult, err error) {
	query, args, err := sqlx.Named(sqlGetStoreById, namedArgs)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// LoadSupplier is wrapper for select query "LoadSupplier".
func LoadSupplier(acc Accessor, additionalQueries ...string) (res []LoadSupplierResult, err error) {
	query := sqlLoadSupplier

	if len(additionalQueries) > 0 {
		query += " " + strings.Join(additionalQueries, " ")
	}

	err = acc.Select(&res, query)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}

// StmtLoadAccount is used to prepare select query "LoadAccount".
func StmtLoadAccount(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadAccount)
}

// StmtGetAccountById is used to prepare select query "GetAccountById".
func StmtGetAccountById(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetAccountById)
}

// StmtGetAccountByUsername is used to prepare select query "GetAccountByUsername".
func StmtGetAccountByUsername(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetAccountByUsername)
}

// StmtGetAdminCount is used to prepare select query "GetAdminCount".
func StmtGetAdminCount(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetAdminCount)
}

// StmtGetCorrectionDetail is used to prepare select query "GetCorrectionDetail".
func StmtGetCorrectionDetail(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetCorrectionDetail)
}

// StmtLoadCorrection is used to prepare select query "LoadCorrection".
func StmtLoadCorrection(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadCorrection)
}

// StmtGetCorrectionById is used to prepare select query "GetCorrectionById".
func StmtGetCorrectionById(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetCorrectionById)
}

// StmtLoadCustomer is used to prepare select query "LoadCustomer".
func StmtLoadCustomer(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadCustomer)
}

// StmtLoadDevice is used to prepare select query "LoadDevice".
func StmtLoadDevice(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadDevice)
}

// StmtGetDeviceByIdentifier is used to prepare select query "GetDeviceByIdentifier".
func StmtGetDeviceByIdentifier(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetDeviceByIdentifier)
}

// StmtLoadEmployee is used to prepare select query "LoadEmployee".
func StmtLoadEmployee(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadEmployee)
}

// StmtLoadProductInEvent is used to prepare select query "LoadProductInEvent".
func StmtLoadProductInEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadProductInEvent)
}

// StmtLoadEvent is used to prepare select query "LoadEvent".
func StmtLoadEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadEvent)
}

// StmtGetEvent is used to prepare select query "GetEvent".
func StmtGetEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetEvent)
}

// StmtGetProductActiveEvent is used to prepare select query "GetProductActiveEvent".
func StmtGetProductActiveEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetProductActiveEvent)
}

// StmtGetProductPastEvent is used to prepare select query "GetProductPastEvent".
func StmtGetProductPastEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetProductPastEvent)
}

// StmtLoadExpenseCategory is used to prepare select query "LoadExpenseCategory".
func StmtLoadExpenseCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadExpenseCategory)
}

// StmtGetExpenseCategoryById is used to prepare select query "GetExpenseCategoryById".
func StmtGetExpenseCategoryById(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetExpenseCategoryById)
}

// StmtLoadExpense is used to prepare select query "LoadExpense".
func StmtLoadExpense(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadExpense)
}

// StmtGetSumExpense is used to prepare select query "GetSumExpense".
func StmtGetSumExpense(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetSumExpense)
}

// StmtLoadPaymentMethod is used to prepare select query "LoadPaymentMethod".
func StmtLoadPaymentMethod(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadPaymentMethod)
}

// StmtLoadProductCategory is used to prepare select query "LoadProductCategory".
func StmtLoadProductCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadProductCategory)
}

// StmtGetProduct is used to prepare select query "GetProduct".
func StmtGetProduct(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetProduct)
}

// StmtGetProductDetail is used to prepare select query "GetProductDetail".
func StmtGetProductDetail(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetProductDetail)
}

// StmtGetProductLeftover is used to prepare select query "GetProductLeftover".
func StmtGetProductLeftover(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetProductLeftover)
}

// StmtLoadAllPurchase is used to prepare select query "LoadAllPurchase".
func StmtLoadAllPurchase(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadAllPurchase)
}

// StmtLoadActivePurchase is used to prepare select query "LoadActivePurchase".
func StmtLoadActivePurchase(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadActivePurchase)
}

// StmtGetPurchase is used to prepare select query "GetPurchase".
func StmtGetPurchase(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetPurchase)
}

// StmtGetPurchaseDetail is used to prepare select query "GetPurchaseDetail".
func StmtGetPurchaseDetail(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetPurchaseDetail)
}

// StmtGetPurchaseDetailSum is used to prepare select query "GetPurchaseDetailSum".
func StmtGetPurchaseDetailSum(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetPurchaseDetailSum)
}

// StmtLoadStore is used to prepare select query "LoadStore".
func StmtLoadStore(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadStore)
}

// StmtGetStoreById is used to prepare select query "GetStoreById".
func StmtGetStoreById(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlGetStoreById)
}

// StmtLoadSupplier is used to prepare select query "LoadSupplier".
func StmtLoadSupplier(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlLoadSupplier)
}

const sqlLoadAccount = `
SELECT id, username, name, permission 
FROM account WHERE deleted = 0`

const sqlGetAccountById = `
SELECT id, username, name, permission
FROM account WHERE id = :id`

const sqlGetAccountByUsername = `
SELECT id, username, name, permission
FROM account WHERE username = :username`

const sqlGetAdminCount = `
SELECT COUNT(id) FROM account 
WHERE deleted = 0
AND JSON_CONTAINS(permission, '"admin"')`

const sqlGetCorrectionDetail = `
SELECT p.id, p.identifier, p.name, pc.name category, cd.qty, p.capital, p.price
FROM correction_detail cd
LEFT JOIN product p ON cd.product_id = p.id
LEFT JOIN product_category pc ON p.category_id = pc.id
WHERE cd.correction_id = :correction_id AND cd.deleted = 0
ORDER BY p.identifier`

const sqlLoadCorrection = `
WITH correction_summary AS (
	SELECT correction_id id, COUNT(product_id) n_products, SUM(qty) qty
	FROM correction_detail cd
	WHERE cd.deleted = 0
	GROUP BY correction_id),
final_data AS (
	SELECT c.id, c.account_id, a.name account, c.description,
		CONVERT_TZ(c.correction_time, "+00::00", :timezone) correction_time,
		IFNULL(cs.qty, 0) qty,
		IFNULL(cs.n_products, 0) n_products
	FROM correction c
	LEFT JOIN account a ON c.account_id = a.id
	LEFT JOIN correction_summary cs ON cs.id = c.id
	WHERE c.deleted = 0)
SELECT * FROM final_data`

const sqlGetCorrectionById = `
SELECT id, account_id, description, correction_time
FROM correction WHERE id = :id`

const sqlLoadCustomer = `
SELECT id, name, identifier, phone, address, city, info
FROM customer
WHERE deleted = 0`

const sqlLoadDevice = `
SELECT d.id, d.name, d.store_id, s.name store, d.identifier, d.allowed
FROM device d
LEFT JOIN store s ON d.store_id = s.id`

const sqlGetDeviceByIdentifier = `
SELECT d.id, d.name, d.store_id, s.name store, d.identifier, d.allowed
FROM device d
LEFT JOIN store s ON d.store_id = s.id
WHERE d.identifier = :identifier`

const sqlLoadEmployee = `
SELECT id, username, name, gender, religion,
	phone, address, city, start_year, end_year
FROM employee WHERE deleted = 0`

const sqlLoadProductInEvent = `
SELECT p.id, p.identifier, p.name, pc.name category, ed.qty, ed.capital, ed.price
FROM event_detail ed
LEFT JOIN product p ON ed.product_id = p.id
LEFT JOIN product_category pc ON p.category_id = pc.id
WHERE ed.event_id = :event_id AND ed.deleted = 0`

const sqlLoadEvent = `
WITH n_event_products AS (
	SELECT event_id id, SUM(qty) qty, COUNT(product_id) n_products
	FROM event_detail
	WHERE deleted = 0
	GROUP BY event_id),
final_data AS (
	SELECT e.id, e.name, e.percentage, e.rounding, e.fixed_price,
		IFNULL(nep.qty, 0) qty,
		IFNULL(nep.n_products, 0) n_products,
		CONVERT_TZ(e.start, "+00::00", :timezone) start,
		CONVERT_TZ(e.end, "+00::00", :timezone) end
	FROM event e
	LEFT JOIN n_event_products nep ON e.id = nep.id
	WHERE e.deleted = 0)
SELECT * FROM final_data`

const sqlGetEvent = `
SELECT id, name, percentage, rounding, fixed_price, start, end
FROM event
WHERE id = :id AND deleted = 0`

const sqlGetProductActiveEvent = `
SELECT e.id, e.name, e.percentage, e.rounding, e.fixed_price, e.start, e.end
FROM event_detail ed
LEFT JOIN event e ON ed.event_id = e.id 
WHERE ed.product_id = :product_id
	AND e.deleted = 0
	AND ed.deleted = 0
	AND e.start <= UTC_TIMESTAMP()
	AND (e.end >= UTC_TIMESTAMP() OR e.end IS NULL)
ORDER BY event_id DESC 
LIMIT 1`

const sqlGetProductPastEvent = `
WITH fixed_event AS (
	SELECT id, name, percentage, rounding, fixed_price,
		start, IFNULL(end, UTC_TIMESTAMP()) end
	FROM event
	WHERE deleted = 0)
SELECT e.id, e.name, e.percentage, e.rounding, e.fixed_price, e.start, e.end
FROM event_detail ed
LEFT JOIN fixed_event e ON ed.event_id = e.id 
WHERE ed.product_id = :product_id
AND NOT(ed.event_id <=> :event_id)
AND ed.deleted = 0
AND (
	(e.start <= :start_time AND e.end >= :end_time) OR
	(e.start <= IFNULL(:start_time, UTC_TIMESTAMP) AND e.end >= IFNULL(:end_time, UTC_TIMESTAMP)) OR
	(e.start >= :start_time AND e.end <= IFNULL(:end_time, UTC_TIMESTAMP))
)
ORDER BY event_id DESC 
LIMIT 1`

const sqlLoadExpenseCategory = `
SELECT id as value, name as caption
FROM expense_category
WHERE deleted = 0 AND parent_id <=> :parent_id
ORDER BY name`

const sqlGetExpenseCategoryById = `
SELECT id, name, parent_id
FROM expense_category WHERE id = :id`

const sqlLoadExpense = `
SELECT e.id, e.name, e.category_id, c.name as category, e.amount
FROM expense e
LEFT JOIN expense_category c ON e.category_id = c.id
WHERE e.deleted = 0
AND e.expense_date = :expense_date`

const sqlGetSumExpense = `
SELECT IFNULL(SUM(amount), 0) amount
FROM expense
WHERE deleted = 0
AND expense_date = :expense_date`

const sqlLoadPaymentMethod = `
SELECT id, name FROM payment_method WHERE deleted = 0`

const sqlLoadProductCategory = `
SELECT id as value, name as caption
FROM product_category
WHERE deleted = 0 AND parent_id <=> :parent_id
ORDER BY name`

const sqlGetProduct = `
SELECT id, purchase_id, category_id, identifier,
	name, capital, price, specs
FROM product
WHERE deleted = 0 AND identifier = :identifier`

const sqlGetProductDetail = `
SELECT p.id, p.purchase_id, p.identifier, pp.supplier_id, s.name supplier,
	p.category_id, pc.name category, p.name, p.capital, p.price, p.specs
FROM product p
LEFT JOIN purchase pp ON p.purchase_id = pp.id
LEFT JOIN supplier s ON pp.supplier_id = s.id
LEFT JOIN product_category pc ON p.category_id = pc.id
WHERE p.deleted = 0 AND p.identifier = :identifier`

const sqlGetProductLeftover = `
WITH product_qty AS (
	SELECT pd.product_id, SUM(pd.qty) qty
	FROM purchase_detail pd 
	LEFT JOIN purchase p ON pd.purchase_id = p.id
	WHERE p.deleted = 0
	GROUP BY pd.product_id),
product_correction AS (
	SELECT cd.product_id, SUM(cd.qty) qty
	FROM correction_detail cd
	LEFT JOIN correction c ON cd.correction_id = c.id
	WHERE c.deleted = 0
	AND cd.deleted = 0
	GROUP BY cd.product_id),
product_sold AS (
	SELECT sd.product_id, SUM(sd.qty) qty
	FROM sales_detail sd 
	LEFT JOIN sales s ON sd.sales_id = s.id
	AND s.deleted = 0
	GROUP BY sd.product_id)
SELECT p.id, p.identifier, p.name, cat.name category, p.capital, p.price,
	IFNULL(pq.qty, 0) + IFNULL(pc.qty, 0) - IFNULL(ps.qty, 0) qty
FROM product p
LEFT JOIN product_category cat ON p.category_id = cat.id
LEFT JOIN product_qty pq ON pq.product_id = p.id
LEFT JOIN product_correction pc ON pc.product_id = p.id
LEFT JOIN product_sold ps ON ps.product_id = p.id
WHERE p.deleted = 0 AND p.identifier = :identifier`

const sqlLoadAllPurchase = `
WITH purchase_by_category AS (
	SELECT pd.purchase_id id, p.category_id,
		SUM(pd.qty) qty, pc.name category
	FROM purchase_detail pd
	LEFT JOIN product p ON pd.product_id = p.id
	LEFT JOIN product_category pc ON p.category_id = pc.id
	GROUP BY pd.purchase_id, p.category_id
	ORDER BY id ASC, qty DESC),
purchase_by_category_grouped AS (
	SELECT id, category
	FROM purchase_by_category
	GROUP BY id),
purchase_finished AS (
	SELECT pd.purchase_id id,
		(MIN(pd.price + p.price) <> 0 OR p.deleted = 1) finished
	FROM purchase_detail pd
	LEFT JOIN product p ON pd.product_id = p.id
	GROUP BY pd.purchase_id),
purchase_profit AS (
	SELECT purchase_id, ROUND(IFNULL((SUM(qty * price) / SUM(qty * capital) - 1) * 100, 0), 4) percent_profit
	FROM purchase_detail pd
	GROUP BY purchase_id),
n_purchase_products AS (
	SELECT purchase_id id, COUNT(product_id) n_products
	FROM purchase_detail
	GROUP BY purchase_id),
final_data AS (
	SELECT p.id,
		a.id account_id,
		a.name account,
		s.id supplier_id,
		s.name supplier,
		IFNULL(pbcg.category, "") category,
		p.receipt_date,
		CONVERT_TZ(p.input_time, "+00::00", :timezone) input_time,
		IFNULL(npp.n_products, 0) n_products,
		p.qty,
		p.total_purchase,
		p.payment_due,
		IFNULL(p.paid_amount, 0) paid_amount,
		p.paid_date,
		IFNULL(pp.percent_profit, 0) percent_profit,
		IFNULL(pf.finished, 0) finished
	FROM purchase p
	LEFT JOIN supplier s ON p.supplier_id = s.id
	LEFT JOIN account a ON p.account_id = a.id
	LEFT JOIN purchase_finished pf ON pf.id = p.id
	LEFT JOIN purchase_by_category_grouped pbcg ON pbcg.id = p.id
	LEFT JOIN purchase_profit pp ON pp.purchase_id = p.id
	LEFT JOIN n_purchase_products npp ON npp.id = p.id
	WHERE p.account_id IS NOT NULL
	AND p.deleted = 0)
SELECT * FROM final_data`

const sqlLoadActivePurchase = `
WITH purchase_by_category AS (
	SELECT pd.purchase_id id, p.category_id,
		SUM(pd.qty) qty, pc.name category
	FROM purchase_detail pd
	LEFT JOIN product p ON pd.product_id = p.id
	LEFT JOIN product_category pc ON p.category_id = pc.id
	GROUP BY pd.purchase_id, p.category_id
	ORDER BY id ASC, qty DESC),
purchase_by_category_grouped AS (
	SELECT id, category
	FROM purchase_by_category
	GROUP BY id),
purchase_finished AS (
	SELECT pd.purchase_id id,
		(MIN(pd.price + p.price) <> 0 OR p.deleted = 1) finished
	FROM purchase_detail pd
	LEFT JOIN product p ON pd.product_id = p.id
	GROUP BY pd.purchase_id),
n_purchase_products AS (
	SELECT purchase_id id, COUNT(product_id) n_products
	FROM purchase_detail
	GROUP BY purchase_id),
visible_purchase AS (
	SELECT p.id,
		a.id account_id,
		a.name account,
		s.id supplier_id,
		s.name supplier,
		p.receipt_date,
		CONVERT_TZ(p.input_time, "+00::00", :timezone) input_time,
		IFNULL(npp.n_products, 0) n_products,
		p.qty,
		p.total_purchase
	FROM purchase p
	LEFT JOIN supplier s ON p.supplier_id = s.id
	LEFT JOIN account a ON p.account_id = a.id
	LEFT JOIN n_purchase_products npp ON p.id = npp.id
	WHERE p.input_time >= :start_date
	AND p.account_id IS NOT NULL
	AND p.deleted = 0)
SELECT vp.id, vp.account_id, vp.account, vp.supplier_id, vp.supplier,
	IFNULL(pc.category, "") category, vp.receipt_date, vp.input_time,
	vp.n_products, vp.qty, vp.total_purchase,
	IFNULL(pf.finished, 0) finished
FROM visible_purchase vp
LEFT JOIN purchase_finished pf ON pf.id = vp.id
LEFT JOIN purchase_by_category_grouped pc ON pc.id = vp.id
WHERE 1`

const sqlGetPurchase = `
SELECT id, account_id, 
	supplier_id, receipt_date, input_time, qty,
	total_purchase, payment_due, paid_amount, paid_date
FROM purchase
WHERE deleted = 0 AND id = :id`

const sqlGetPurchaseDetail = `
SELECT p.id, p.identifier, p.name, pc.name category, pd.qty, pd.capital, 
	IF(pd.price > 0, pd.price, p.price) price,
	pd.qty * pd.capital total_capital
FROM purchase_detail pd
LEFT JOIN product p ON pd.product_id = p.id
LEFT JOIN product_category pc ON p.category_id = pc.id
WHERE pd.purchase_id = :purchase_id
ORDER BY pd.position`

const sqlGetPurchaseDetailSum = `
SELECT IFNULL(SUM(qty), 0) qty,
	IFNULL(SUM(qty * capital), 0) total_capital
FROM purchase_detail
WHERE purchase_id = :purchase_id`

const sqlLoadStore = `
SELECT id, name, phone, address, city, info
FROM store WHERE deleted = 0`

const sqlGetStoreById = `
SELECT id, name, phone, address, city, info
FROM store WHERE deleted = 0 AND id = :id`

const sqlLoadSupplier = `
WITH supplier_by_category AS (
	SELECT pu.supplier_id id, p.category_id, SUM(p.qty) qty
	FROM product p
	LEFT JOIN purchase pu ON p.purchase_id = pu.id
	WHERE p.deleted = 0
	GROUP BY pu.supplier_id, p.category_id
	ORDER BY id ASC, qty DESC),
supplier_by_category_grouped AS (
	SELECT id, category_id
	FROM supplier_by_category
	GROUP BY id),
final_data AS (
	SELECT s.id, s.name, s.phone, s.address, s.city, s.info, 
		sbcg.category_id, pc.name category
	FROM supplier s
	LEFT JOIN supplier_by_category_grouped sbcg ON s.id = sbcg.id
	LEFT JOIN product_category pc ON pc.id = sbcg.category_id
	WHERE s.deleted = 0)
SELECT * FROM final_data`
