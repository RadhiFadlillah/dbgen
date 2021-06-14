package storage

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// InsertAccount is wrapper for exec query "InsertAccount".
func InsertAccount(acc Accessor, namedArgs InsertAccountArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertAccount, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteAccount is wrapper for exec query "DeleteAccount".
func DeleteAccount(acc Accessor, namedArgs DeleteAccountArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteAccount, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateAccount is wrapper for exec query "UpdateAccount".
func UpdateAccount(acc Accessor, namedArgs UpdateAccountArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateAccount, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateAccountPassword is wrapper for exec query "UpdateAccountPassword".
func UpdateAccountPassword(acc Accessor, namedArgs UpdateAccountPasswordArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateAccountPassword, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertCorrectionDetail is wrapper for exec query "InsertCorrectionDetail".
func InsertCorrectionDetail(acc Accessor, namedArgs InsertCorrectionDetailArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertCorrectionDetail, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteCorrectionDetailByCorrectionId is wrapper for exec query "DeleteCorrectionDetailByCorrectionId".
func DeleteCorrectionDetailByCorrectionId(acc Accessor, namedArgs DeleteCorrectionDetailByCorrectionIdArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteCorrectionDetailByCorrectionId, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteCorrectionDetailByProduct is wrapper for exec query "DeleteCorrectionDetailByProduct".
func DeleteCorrectionDetailByProduct(acc Accessor, namedArgs DeleteCorrectionDetailByProductArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteCorrectionDetailByProduct, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertCorrection is wrapper for exec query "InsertCorrection".
func InsertCorrection(acc Accessor, namedArgs InsertCorrectionArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertCorrection, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteCorrection is wrapper for exec query "DeleteCorrection".
func DeleteCorrection(acc Accessor, namedArgs DeleteCorrectionArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteCorrection, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateCorrection is wrapper for exec query "UpdateCorrection".
func UpdateCorrection(acc Accessor, namedArgs UpdateCorrectionArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateCorrection, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertCustomer is wrapper for exec query "InsertCustomer".
func InsertCustomer(acc Accessor, namedArgs InsertCustomerArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertCustomer, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteCustomer is wrapper for exec query "DeleteCustomer".
func DeleteCustomer(acc Accessor, namedArgs DeleteCustomerArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteCustomer, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateCustomer is wrapper for exec query "UpdateCustomer".
func UpdateCustomer(acc Accessor, namedArgs UpdateCustomerArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateCustomer, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertDevice is wrapper for exec query "InsertDevice".
func InsertDevice(acc Accessor, namedArgs InsertDeviceArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertDevice, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteDevice is wrapper for exec query "DeleteDevice".
func DeleteDevice(acc Accessor, namedArgs DeleteDeviceArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteDevice, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateDevice is wrapper for exec query "UpdateDevice".
func UpdateDevice(acc Accessor, namedArgs UpdateDeviceArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateDevice, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// SetDeviceStore is wrapper for exec query "SetDeviceStore".
func SetDeviceStore(acc Accessor, namedArgs SetDeviceStoreArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlSetDeviceStore, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// MarkDeviceAllowed is wrapper for exec query "MarkDeviceAllowed".
func MarkDeviceAllowed(acc Accessor, namedArgs MarkDeviceAllowedArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlMarkDeviceAllowed, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertEmployee is wrapper for exec query "InsertEmployee".
func InsertEmployee(acc Accessor, namedArgs InsertEmployeeArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertEmployee, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteEmployee is wrapper for exec query "DeleteEmployee".
func DeleteEmployee(acc Accessor, namedArgs DeleteEmployeeArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteEmployee, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateEmployee is wrapper for exec query "UpdateEmployee".
func UpdateEmployee(acc Accessor, namedArgs UpdateEmployeeArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateEmployee, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertEventDetail is wrapper for exec query "InsertEventDetail".
func InsertEventDetail(acc Accessor, namedArgs InsertEventDetailArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertEventDetail, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateEventDetailQty is wrapper for exec query "UpdateEventDetailQty".
func UpdateEventDetailQty(acc Accessor, namedArgs UpdateEventDetailQtyArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateEventDetailQty, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateEventDetailPrice is wrapper for exec query "UpdateEventDetailPrice".
func UpdateEventDetailPrice(acc Accessor, namedArgs UpdateEventDetailPriceArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateEventDetailPrice, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteEventDetailByEventId is wrapper for exec query "DeleteEventDetailByEventId".
func DeleteEventDetailByEventId(acc Accessor, namedArgs DeleteEventDetailByEventIdArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteEventDetailByEventId, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteEventProduct is wrapper for exec query "DeleteEventProduct".
func DeleteEventProduct(acc Accessor, namedArgs DeleteEventProductArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteEventProduct, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// CreateEvent is wrapper for exec query "CreateEvent".
func CreateEvent(acc Accessor, namedArgs CreateEventArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlCreateEvent, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteEvent is wrapper for exec query "DeleteEvent".
func DeleteEvent(acc Accessor, namedArgs DeleteEventArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteEvent, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateEvent is wrapper for exec query "UpdateEvent".
func UpdateEvent(acc Accessor, namedArgs UpdateEventArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateEvent, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertExpenseCategory is wrapper for exec query "InsertExpenseCategory".
func InsertExpenseCategory(acc Accessor, namedArgs InsertExpenseCategoryArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertExpenseCategory, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteExpenseCategory is wrapper for exec query "DeleteExpenseCategory".
func DeleteExpenseCategory(acc Accessor, namedArgs DeleteExpenseCategoryArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteExpenseCategory, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateExpenseCategory is wrapper for exec query "UpdateExpenseCategory".
func UpdateExpenseCategory(acc Accessor, namedArgs UpdateExpenseCategoryArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateExpenseCategory, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertExpense is wrapper for exec query "InsertExpense".
func InsertExpense(acc Accessor, namedArgs InsertExpenseArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertExpense, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteExpense is wrapper for exec query "DeleteExpense".
func DeleteExpense(acc Accessor, namedArgs DeleteExpenseArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteExpense, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateExpense is wrapper for exec query "UpdateExpense".
func UpdateExpense(acc Accessor, namedArgs UpdateExpenseArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateExpense, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertPaymentMethod is wrapper for exec query "InsertPaymentMethod".
func InsertPaymentMethod(acc Accessor, namedArgs InsertPaymentMethodArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertPaymentMethod, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeletePaymentMethod is wrapper for exec query "DeletePaymentMethod".
func DeletePaymentMethod(acc Accessor, namedArgs DeletePaymentMethodArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeletePaymentMethod, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdatePaymentMethod is wrapper for exec query "UpdatePaymentMethod".
func UpdatePaymentMethod(acc Accessor, namedArgs UpdatePaymentMethodArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdatePaymentMethod, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertProductCategory is wrapper for exec query "InsertProductCategory".
func InsertProductCategory(acc Accessor, namedArgs InsertProductCategoryArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertProductCategory, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteProductCategory is wrapper for exec query "DeleteProductCategory".
func DeleteProductCategory(acc Accessor, namedArgs DeleteProductCategoryArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteProductCategory, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateProductCategory is wrapper for exec query "UpdateProductCategory".
func UpdateProductCategory(acc Accessor, namedArgs UpdateProductCategoryArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateProductCategory, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertProduct is wrapper for exec query "InsertProduct".
func InsertProduct(acc Accessor, namedArgs InsertProductArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertProduct, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// SetProductIdentifier is wrapper for exec query "SetProductIdentifier".
func SetProductIdentifier(acc Accessor, namedArgs SetProductIdentifierArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlSetProductIdentifier, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertPurchase is wrapper for exec query "InsertPurchase".
func InsertPurchase(acc Accessor, namedArgs InsertPurchaseArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertPurchase, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdatePurchase is wrapper for exec query "UpdatePurchase".
func UpdatePurchase(acc Accessor, namedArgs UpdatePurchaseArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdatePurchase, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdatePurchasePayment is wrapper for exec query "UpdatePurchasePayment".
func UpdatePurchasePayment(acc Accessor, namedArgs UpdatePurchasePaymentArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdatePurchasePayment, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertPurchaseDetail is wrapper for exec query "InsertPurchaseDetail".
func InsertPurchaseDetail(acc Accessor, namedArgs InsertPurchaseDetailArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertPurchaseDetail, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertStore is wrapper for exec query "InsertStore".
func InsertStore(acc Accessor, namedArgs InsertStoreArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertStore, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteStore is wrapper for exec query "DeleteStore".
func DeleteStore(acc Accessor, namedArgs DeleteStoreArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteStore, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateStore is wrapper for exec query "UpdateStore".
func UpdateStore(acc Accessor, namedArgs UpdateStoreArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateStore, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// InsertSupplier is wrapper for exec query "InsertSupplier".
func InsertSupplier(acc Accessor, namedArgs InsertSupplierArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlInsertSupplier, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// DeleteSupplier is wrapper for exec query "DeleteSupplier".
func DeleteSupplier(acc Accessor, namedArgs DeleteSupplierArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlDeleteSupplier, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// UpdateSupplier is wrapper for exec query "UpdateSupplier".
func UpdateSupplier(acc Accessor, namedArgs UpdateSupplierArgs) (sql.Result, error) {
	query, args, err := sqlx.Named(sqlUpdateSupplier, namedArgs)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return acc.Exec(query)
}

// StmtInsertAccount is used to prepare exec query "InsertAccount".
func StmtInsertAccount(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertAccount)
}

// StmtDeleteAccount is used to prepare exec query "DeleteAccount".
func StmtDeleteAccount(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteAccount)
}

// StmtUpdateAccount is used to prepare exec query "UpdateAccount".
func StmtUpdateAccount(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateAccount)
}

// StmtUpdateAccountPassword is used to prepare exec query "UpdateAccountPassword".
func StmtUpdateAccountPassword(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateAccountPassword)
}

// StmtInsertCorrectionDetail is used to prepare exec query "InsertCorrectionDetail".
func StmtInsertCorrectionDetail(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertCorrectionDetail)
}

// StmtDeleteCorrectionDetailByCorrectionId is used to prepare exec query "DeleteCorrectionDetailByCorrectionId".
func StmtDeleteCorrectionDetailByCorrectionId(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteCorrectionDetailByCorrectionId)
}

// StmtDeleteCorrectionDetailByProduct is used to prepare exec query "DeleteCorrectionDetailByProduct".
func StmtDeleteCorrectionDetailByProduct(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteCorrectionDetailByProduct)
}

// StmtInsertCorrection is used to prepare exec query "InsertCorrection".
func StmtInsertCorrection(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertCorrection)
}

// StmtDeleteCorrection is used to prepare exec query "DeleteCorrection".
func StmtDeleteCorrection(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteCorrection)
}

// StmtUpdateCorrection is used to prepare exec query "UpdateCorrection".
func StmtUpdateCorrection(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateCorrection)
}

// StmtInsertCustomer is used to prepare exec query "InsertCustomer".
func StmtInsertCustomer(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertCustomer)
}

// StmtDeleteCustomer is used to prepare exec query "DeleteCustomer".
func StmtDeleteCustomer(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteCustomer)
}

// StmtUpdateCustomer is used to prepare exec query "UpdateCustomer".
func StmtUpdateCustomer(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateCustomer)
}

// StmtInsertDevice is used to prepare exec query "InsertDevice".
func StmtInsertDevice(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertDevice)
}

// StmtDeleteDevice is used to prepare exec query "DeleteDevice".
func StmtDeleteDevice(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteDevice)
}

// StmtUpdateDevice is used to prepare exec query "UpdateDevice".
func StmtUpdateDevice(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateDevice)
}

// StmtSetDeviceStore is used to prepare exec query "SetDeviceStore".
func StmtSetDeviceStore(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlSetDeviceStore)
}

// StmtMarkDeviceAllowed is used to prepare exec query "MarkDeviceAllowed".
func StmtMarkDeviceAllowed(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlMarkDeviceAllowed)
}

// StmtInsertEmployee is used to prepare exec query "InsertEmployee".
func StmtInsertEmployee(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertEmployee)
}

// StmtDeleteEmployee is used to prepare exec query "DeleteEmployee".
func StmtDeleteEmployee(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteEmployee)
}

// StmtUpdateEmployee is used to prepare exec query "UpdateEmployee".
func StmtUpdateEmployee(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateEmployee)
}

// StmtInsertEventDetail is used to prepare exec query "InsertEventDetail".
func StmtInsertEventDetail(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertEventDetail)
}

// StmtUpdateEventDetailQty is used to prepare exec query "UpdateEventDetailQty".
func StmtUpdateEventDetailQty(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateEventDetailQty)
}

// StmtUpdateEventDetailPrice is used to prepare exec query "UpdateEventDetailPrice".
func StmtUpdateEventDetailPrice(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateEventDetailPrice)
}

// StmtDeleteEventDetailByEventId is used to prepare exec query "DeleteEventDetailByEventId".
func StmtDeleteEventDetailByEventId(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteEventDetailByEventId)
}

// StmtDeleteEventProduct is used to prepare exec query "DeleteEventProduct".
func StmtDeleteEventProduct(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteEventProduct)
}

// StmtCreateEvent is used to prepare exec query "CreateEvent".
func StmtCreateEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlCreateEvent)
}

// StmtDeleteEvent is used to prepare exec query "DeleteEvent".
func StmtDeleteEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteEvent)
}

// StmtUpdateEvent is used to prepare exec query "UpdateEvent".
func StmtUpdateEvent(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateEvent)
}

// StmtInsertExpenseCategory is used to prepare exec query "InsertExpenseCategory".
func StmtInsertExpenseCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertExpenseCategory)
}

// StmtDeleteExpenseCategory is used to prepare exec query "DeleteExpenseCategory".
func StmtDeleteExpenseCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteExpenseCategory)
}

// StmtUpdateExpenseCategory is used to prepare exec query "UpdateExpenseCategory".
func StmtUpdateExpenseCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateExpenseCategory)
}

// StmtInsertExpense is used to prepare exec query "InsertExpense".
func StmtInsertExpense(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertExpense)
}

// StmtDeleteExpense is used to prepare exec query "DeleteExpense".
func StmtDeleteExpense(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteExpense)
}

// StmtUpdateExpense is used to prepare exec query "UpdateExpense".
func StmtUpdateExpense(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateExpense)
}

// StmtInsertPaymentMethod is used to prepare exec query "InsertPaymentMethod".
func StmtInsertPaymentMethod(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertPaymentMethod)
}

// StmtDeletePaymentMethod is used to prepare exec query "DeletePaymentMethod".
func StmtDeletePaymentMethod(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeletePaymentMethod)
}

// StmtUpdatePaymentMethod is used to prepare exec query "UpdatePaymentMethod".
func StmtUpdatePaymentMethod(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdatePaymentMethod)
}

// StmtInsertProductCategory is used to prepare exec query "InsertProductCategory".
func StmtInsertProductCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertProductCategory)
}

// StmtDeleteProductCategory is used to prepare exec query "DeleteProductCategory".
func StmtDeleteProductCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteProductCategory)
}

// StmtUpdateProductCategory is used to prepare exec query "UpdateProductCategory".
func StmtUpdateProductCategory(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateProductCategory)
}

// StmtInsertProduct is used to prepare exec query "InsertProduct".
func StmtInsertProduct(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertProduct)
}

// StmtSetProductIdentifier is used to prepare exec query "SetProductIdentifier".
func StmtSetProductIdentifier(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlSetProductIdentifier)
}

// StmtInsertPurchase is used to prepare exec query "InsertPurchase".
func StmtInsertPurchase(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertPurchase)
}

// StmtUpdatePurchase is used to prepare exec query "UpdatePurchase".
func StmtUpdatePurchase(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdatePurchase)
}

// StmtUpdatePurchasePayment is used to prepare exec query "UpdatePurchasePayment".
func StmtUpdatePurchasePayment(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdatePurchasePayment)
}

// StmtInsertPurchaseDetail is used to prepare exec query "InsertPurchaseDetail".
func StmtInsertPurchaseDetail(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertPurchaseDetail)
}

// StmtInsertStore is used to prepare exec query "InsertStore".
func StmtInsertStore(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertStore)
}

// StmtDeleteStore is used to prepare exec query "DeleteStore".
func StmtDeleteStore(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteStore)
}

// StmtUpdateStore is used to prepare exec query "UpdateStore".
func StmtUpdateStore(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateStore)
}

// StmtInsertSupplier is used to prepare exec query "InsertSupplier".
func StmtInsertSupplier(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlInsertSupplier)
}

// StmtDeleteSupplier is used to prepare exec query "DeleteSupplier".
func StmtDeleteSupplier(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlDeleteSupplier)
}

// StmtUpdateSupplier is used to prepare exec query "UpdateSupplier".
func StmtUpdateSupplier(acc Accessor) (*sqlx.NamedStmt, error) {
	return acc.PrepareNamed(sqlUpdateSupplier)
}

const sqlInsertAccount = `
INSERT INTO account (username, name, password, permission)
VALUES (:username, :name, password, permission)`

const sqlDeleteAccount = `
UPDATE account
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateAccount = `
UPDATE account
SET username = :username, name = :name, permission = :permission
WHERE id = :id`

const sqlUpdateAccountPassword = `
UPDATE account SET password = :password WHERE id = :id`

const sqlInsertCorrectionDetail = `
INSERT INTO correction_detail (correction_id, product_id, qty)
VALUES (:correction_id, :product_id, :qty)
ON DUPLICATE KEY UPDATE qty = qty + VALUES(qty)`

const sqlDeleteCorrectionDetailByCorrectionId = `
UPDATE correction_detail
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE correction_id = :correction_id`

const sqlDeleteCorrectionDetailByProduct = `
UPDATE correction_detail cd
LEFT JOIN product p ON p.id = cd.product_id
SET cd.deleted = 1, cd.delete_time = UTC_TIMESTAMP()
WHERE cd.correction_id = :correction_id AND p.identifier = :identifier`

const sqlInsertCorrection = `
INSERT INTO correction (account_id, description, correction_time)
VALUES (:account_id, :description, :correction_time)`

const sqlDeleteCorrection = `
UPDATE correction
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateCorrection = `
UPDATE correction
SET description = :description
WHERE id = :id`

const sqlInsertCustomer = `
INSERT INTO customer (name, identifier, phone, address, city, info)
VALUES (:name, :identifier, :phone, :address, :city, :info)`

const sqlDeleteCustomer = `
UPDATE customer
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateCustomer = `
UPDATE customer
SET name = :name, phone = :phone, 
	address = :address, city = :city, info = :info
WHERE id = :id`

const sqlInsertDevice = `
INSERT INTO device (name, identifier)
VALUES (:name, :identifier)`

const sqlDeleteDevice = `
DELETE FROM device WHERE id = :id`

const sqlUpdateDevice = `
UPDATE device 
SET name = :name, allowed = :allowed 
WHERE id = :id`

const sqlSetDeviceStore = `
UPDATE device
SET store_id = :store_id
WHERE id = :id`

const sqlMarkDeviceAllowed = `
UPDATE device SET allowed = 1 WHERE id = :id`

const sqlInsertEmployee = `
INSERT INTO employee (
	username, name, gender, religion, phone,
	address, city, start_year, end_year)
VALUES (
	:username, :name, :gender, :religion, :phone,
	:address, :city, :start_year, :end_year)`

const sqlDeleteEmployee = `
UPDATE employee
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateEmployee = `
UPDATE employee
SET username = :username, name = :name, gender = :gender,
	religion = :religion, phone = :phone, address = :address,
	city = :city, start_year = :start_year, end_year = :end_year
WHERE id = :id`

const sqlInsertEventDetail = `
INSERT INTO event_detail (
	event_id, product_id, qty, capital, price
) VALUES (
	:event_id, :product_id, :qty, :capital, :price
) ON DUPLICATE KEY UPDATE 
	qty = VALUES(qty),
	capital = VALUES(capital),
	price = VALUES(price),
	deleted = 0`

const sqlUpdateEventDetailQty = `
UPDATE event_detail ed
LEFT JOIN product p ON ed.product_id = p.id
SET ed.qty = ed.qty + :qty, 
	ed.capital = p.capital,
	ed.price = p.price
WHERE ed.product_id = :product_id
AND ed.event_id = :event_id`

const sqlUpdateEventDetailPrice = `
UPDATE event_detail ed
LEFT JOIN product p ON ed.product_id = p.id
SET ed.price = p.price
WHERE ed.product_id = :product_id
AND ed.event_id = :event_id`

const sqlDeleteEventDetailByEventId = `
UPDATE event_detail
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE event_id = :event_id AND deleted = 0`

const sqlDeleteEventProduct = `
UPDATE event_detail ed
LEFT JOIN product p ON p.id = ed.product_id
SET ed.deleted = 1, ed.delete_time = UTC_TIMESTAMP()
WHERE ed.event_id = :event_id AND p.identifier = :identifier`

const sqlCreateEvent = `
INSERT INTO event (
	name, percentage, rounding, fixed_price, start, end
) VALUES (
	:name, :percentage, :rounding, :fixed_price, :start, :end
)`

const sqlDeleteEvent = `
UPDATE event
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateEvent = `
UPDATE event
SET name = :name, percentage = :percentage, rounding = :rounding,
	fixed_price = :fixed_price, start = :start, end = :end
WHERE id = :id`

const sqlInsertExpenseCategory = `
INSERT INTO expense_category (parent_id, name)
VALUES (:parent_id, :name)`

const sqlDeleteExpenseCategory = `
UPDATE expense_category
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateExpenseCategory = `
UPDATE expense_category 
SET name = :name
WHERE id = :id`

const sqlInsertExpense = `
INSERT INTO expense (
	account_id, category_id, identifier,
	name, expense_date, amount
) VALUES (
	:account_id, :category_id, :identifier,
	:name, :expense_date, :amount
)`

const sqlDeleteExpense = `
UPDATE expense
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateExpense = `
UPDATE expense
SET account_id = :account_id,
	category_id = :category_id,
	name = :name, amount = :amount
WHERE id = :id`

const sqlInsertPaymentMethod = `
INSERT INTO payment_method (name) VALUES (:name)`

const sqlDeletePaymentMethod = `
UPDATE payment_method
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdatePaymentMethod = `
UPDATE payment_method 
SET name = :name
WHERE id = :id`

const sqlInsertProductCategory = `
INSERT INTO product_category (parent_id, name)
VALUES (:parent_id, :name)`

const sqlDeleteProductCategory = `
UPDATE product_category
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateProductCategory = `
UPDATE product_category 
SET name = :name
WHERE id = :id`

const sqlInsertProduct = `
INSERT INTO product (
	purchase_id, category_id, name, qty, capital, price, specs
) VALUES(
	:purchase_id, :category_id, :name, :qty, :capital, :price, :specs
) ON DUPLICATE KEY UPDATE
	purchase_id = VALUES(purchase_id),
	category_id = VALUES(category_id),
	name        = VALUES(name),
	capital     = ((qty * capital) + (VALUES(qty) * VALUES(capital))) / (qty + VALUES(qty)),
	qty         = qty + VALUES(qty),
	price       = VALUES(price),
	specs       = VALUES(specs)`

const sqlSetProductIdentifier = `
UPDATE product
SET identifier = CONCAT(YEAR(NOW())-2000, id)
WHERE purchase_id = :purchase_id AND identifier IS NULL`

const sqlInsertPurchase = `
INSERT INTO purchase (
	account_id, supplier_id, receipt_date, input_time, qty, total_purchase
) VALUES(
	:account_id, :supplier_id, :receipt_date, :input_time, :qty, :total_purchase
)`

const sqlUpdatePurchase = `
UPDATE purchase
SET receipt_date = :receipt_date, payment_due = :payment_due,
	paid_amount = :paid_amount, paid_date = :paid_date
WHERE id = :id`

const sqlUpdatePurchasePayment = `
UPDATE purchase
SET paid_amount = :paid_amount, paid_date = :paid_date
WHERE id = :id`

const sqlInsertPurchaseDetail = `
INSERT INTO purchase_detail (
	purchase_id, product_id, qty, capital, price, position
) VALUES(
	:purchase_id, :product_id, :qty, :capital, :price, :position
)`

const sqlInsertStore = `
INSERT INTO store (name, phone, address, city, info)
VALUES (:name, :phone, :address, :city, :info)`

const sqlDeleteStore = `
UPDATE store
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateStore = `
UPDATE store
SET name = :name, phone = :phone, 
	address = :address, city = :city, info = :info
WHERE id = :id`

const sqlInsertSupplier = `
INSERT INTO supplier (name, phone, address, city, info)
VALUES (:name, :phone, :address, :city, :info)`

const sqlDeleteSupplier = `
UPDATE supplier
SET deleted = 1, delete_time = UTC_TIMESTAMP()
WHERE id = :id`

const sqlUpdateSupplier = `
UPDATE supplier
SET name = :name, phone = :phone, 
	address = :address, city = :city, info = :info
WHERE id = :id`
