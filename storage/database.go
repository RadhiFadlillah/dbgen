package storage

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// OpenDatabase open database on specified data source.
func OpenDatabase(dataSourceName string) (db *sqlx.DB, err error) {
	// Connect to database
	db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		err = fmt.Errorf("failed to open database: %v", err)
		return
	}
	db.SetConnMaxLifetime(time.Minute)

	// Disable foreign keys
	_, err = db.Exec(`SET FOREIGN_KEY_CHECKS = 0`)
	if err != nil {
		err = fmt.Errorf("failed to disable foreign key: %v", err)
		return
	}

	// If error happened, close database and enable foreign keys
	// before returning result
	defer func() {
		if err != nil {
			if db != nil {
				db.Exec(`SET FOREIGN_KEY_CHECKS = 1`)
				db.Close()
			}
			db = nil
		}
	}()

	// Generate tables
	_, err = db.Exec(ddlCreateAccount)
	if err != nil {
		err = fmt.Errorf("failed to create table \"account\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateDevice)
	if err != nil {
		err = fmt.Errorf("failed to create table \"device\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateEmployee)
	if err != nil {
		err = fmt.Errorf("failed to create table \"employee\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateSupplier)
	if err != nil {
		err = fmt.Errorf("failed to create table \"supplier\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateStore)
	if err != nil {
		err = fmt.Errorf("failed to create table \"store\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateCustomer)
	if err != nil {
		err = fmt.Errorf("failed to create table \"customer\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateExpenseCategory)
	if err != nil {
		err = fmt.Errorf("failed to create table \"expense_category\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateProductCategory)
	if err != nil {
		err = fmt.Errorf("failed to create table \"product_category\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateExpense)
	if err != nil {
		err = fmt.Errorf("failed to create table \"expense\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreatePurchase)
	if err != nil {
		err = fmt.Errorf("failed to create table \"purchase\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateProduct)
	if err != nil {
		err = fmt.Errorf("failed to create table \"product\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreatePurchaseDetail)
	if err != nil {
		err = fmt.Errorf("failed to create table \"purchase_detail\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateEvent)
	if err != nil {
		err = fmt.Errorf("failed to create table \"event\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateEventDetail)
	if err != nil {
		err = fmt.Errorf("failed to create table \"event_detail\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreatePaymentMethod)
	if err != nil {
		err = fmt.Errorf("failed to create table \"payment_method\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateDisplay)
	if err != nil {
		err = fmt.Errorf("failed to create table \"display\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateDisplayDetail)
	if err != nil {
		err = fmt.Errorf("failed to create table \"display_detail\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateSales)
	if err != nil {
		err = fmt.Errorf("failed to create table \"sales\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateSalesDetail)
	if err != nil {
		err = fmt.Errorf("failed to create table \"sales_detail\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateCorrection)
	if err != nil {
		err = fmt.Errorf("failed to create table \"correction\": %v", err)
		return
	}

	_, err = db.Exec(ddlCreateCorrectionDetail)
	if err != nil {
		err = fmt.Errorf("failed to create table \"correction_detail\": %v", err)
		return
	}

	// Return database and error (if any)
	return
}

const ddlCreateAccount = `
CREATE TABLE IF NOT EXISTS account (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	username    VARCHAR(40)  NOT NULL,
	name        VARCHAR(80)  NOT NULL,
	password    BINARY(60)   NOT NULL,
	permission  JSON         NOT NULL DEFAULT '[]',
	deleted     BOOLEAN      NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY account_username_UNIQUE (username)
) CHARACTER SET utf8mb4`

const ddlCreateDevice = `
CREATE TABLE IF NOT EXISTS device (
	id         INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	store_id   INT UNSIGNED  DEFAULT NULL,
	name       VARCHAR(80)   NOT NULL,
	identifier VARBINARY(64) NOT NULL,
	allowed    BOOLEAN       NOT NULL DEFAULT '0',
	PRIMARY KEY (id),
	UNIQUE KEY device_identifier_UNIQUE (identifier),
	FOREIGN KEY device_store_id_FK (store_id) REFERENCES store (id)
) CHARACTER SET utf8mb4`

const ddlCreateEmployee = `
CREATE TABLE IF NOT EXISTS employee (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	name        VARCHAR(80)  NOT NULL,
	username    VARCHAR(40)  NOT NULL,
	gender      VARCHAR(40)  NOT NULL,
	religion    VARCHAR(40)  NOT NULL,
	phone       VARCHAR(80)  NOT NULL,
	address     VARCHAR(150) NOT NULL,
	city        VARCHAR(100) NOT NULL,
	start_year  INT          NOT NULL,
	end_year    INT          DEFAULT NULL,
	deleted     BOOLEAN      NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY employee_username_UNIQUE (username)
) CHARACTER SET utf8mb4`

const ddlCreateSupplier = `
CREATE TABLE IF NOT EXISTS supplier (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	name        VARCHAR(80)  NOT NULL,
	phone       VARCHAR(80)  NOT NULL,
	address     VARCHAR(150) NOT NULL,
	city        VARCHAR(100) NOT NULL,
	info        VARCHAR(560) DEFAULT NULL,
	deleted     BOOLEAN      NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
) CHARACTER SET utf8mb4`

const ddlCreateStore = `
CREATE TABLE IF NOT EXISTS store (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	name        VARCHAR(80)  NOT NULL,
	phone       VARCHAR(80)  NOT NULL,
	address     VARCHAR(150) NOT NULL,
	city        VARCHAR(100) NOT NULL,
	info        VARCHAR(560) DEFAULT NULL,
	deleted     BOOLEAN      NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
) CHARACTER SET utf8mb4`

const ddlCreateCustomer = `
CREATE TABLE IF NOT EXISTS customer (
	id          INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	name        VARCHAR(80)   NOT NULL,
	identifier  VARBINARY(20) NOT NULL,
	phone       VARCHAR(80)   NOT NULL,
	address     VARCHAR(150)  NOT NULL,
	city        VARCHAR(100)  NOT NULL,
	info        VARCHAR(560)  DEFAULT NULL,
	deleted     BOOLEAN       NOT NULL DEFAULT '0',
	delete_time DATETIME      DEFAULT NULL,
	update_time TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY customer_identifier_UNIQUE (identifier)
) CHARACTER SET utf8mb4`

const ddlCreateExpenseCategory = `
CREATE TABLE IF NOT EXISTS expense_category (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	parent_id   INT UNSIGNED DEFAULT NULL,
	name        VARCHAR(80)  NOT NULL,
	deleted     BOOLEAN      NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	INDEX (parent_id),
	FOREIGN KEY expense_category_parent_id_FK (parent_id) REFERENCES expense_category (id)
		ON UPDATE cascade ON DELETE cascade
) CHARACTER SET utf8mb4`

const ddlCreateProductCategory = `
CREATE TABLE IF NOT EXISTS product_category (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	parent_id   INT UNSIGNED DEFAULT NULL,
	name        VARCHAR(80)  NOT NULL,
	deleted     BOOLEAN      NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	INDEX (parent_id),
	FOREIGN KEY product_category_parent_id_FK (parent_id) REFERENCES product_category (id)
		ON UPDATE cascade ON DELETE cascade
) CHARACTER SET utf8mb4`

const ddlCreateExpense = `
CREATE TABLE IF NOT EXISTS expense (
	id           INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	account_id   INT UNSIGNED  DEFAULT NULL,
	employee_id  INT UNSIGNED  DEFAULT NULL,
	store_id     INT UNSIGNED  DEFAULT NULL,
	category_id  INT UNSIGNED  DEFAULT NULL,
	identifier   VARBINARY(20) DEFAULT NULL,
	name         VARCHAR(80)   NOT NULL,
	expense_date DATE          NOT NULL,
	amount       DECIMAL(20,4) NOT NULL,
	deleted      BOOLEAN       NOT NULL DEFAULT '0',
	delete_time  DATETIME      DEFAULT NULL,
	update_time  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY expense_identifier_UNIQUE (identifier),
	INDEX (account_id),
	INDEX (store_id),
	INDEX (category_id),
	INDEX (expense_date),
	FOREIGN KEY expense_account_id_FK (account_id) REFERENCES account (id),
	FOREIGN KEY expense_employee_id_FK (employee_id) REFERENCES employee (id),
	FOREIGN KEY expense_store_id_FK (store_id) REFERENCES store (id),
	FOREIGN KEY expense_category_id_FK (category_id) REFERENCES expense_category (id)
) CHARACTER SET utf8mb4`

const ddlCreatePurchase = `
CREATE TABLE IF NOT EXISTS purchase (
	id             INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	account_id     INT UNSIGNED  DEFAULT NULL,
	supplier_id    INT UNSIGNED  NOT NULL,
	receipt_date   DATE          NOT NULL,
	input_time     DATETIME      NOT NULL,
	qty            DECIMAL(20,4) NOT NULL,
	total_purchase DECIMAL(20,4) NOT NULL,
	payment_due    DATE          DEFAULT NULL,
	paid_amount    DECIMAL(20,4) DEFAULT NULL,
	paid_date      DATE          DEFAULT NULL,
	deleted        TINYINT(1)    NOT NULL DEFAULT '0',
	delete_time    DATETIME      DEFAULT NULL,
	update_time    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	INDEX (account_id),
	INDEX (supplier_id),
	FOREIGN KEY purchase_account_id_FK (account_id) REFERENCES account (id),
	FOREIGN KEY purchase_supplier_id_FK (supplier_id) REFERENCES supplier (id)
) CHARACTER SET utf8mb4`

const ddlCreateProduct = `
CREATE TABLE IF NOT EXISTS product (
	id           INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	purchase_id  INT UNSIGNED  DEFAULT NULL,
	category_id  INT UNSIGNED  NOT NULL,
	identifier   VARBINARY(20) DEFAULT NULL,
	name         VARCHAR(80)   NOT NULL,
	qty          DECIMAL(20,4) NOT NULL,
	capital      DECIMAL(20,4) NOT NULL,
	price        DECIMAL(20,4) NOT NULL,
	specs        JSON          DEFAULT NULL,
	deleted      TINYINT(1)    NOT NULL DEFAULT '0',
	delete_time  DATETIME      DEFAULT NULL,
	update_time  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY product_identifier_UNIQUE (identifier),
	INDEX (purchase_id),
	INDEX (category_id),
	FOREIGN KEY product_purchase_id_FK (purchase_id) REFERENCES purchase (id),
	FOREIGN KEY product_category_id_FK (category_id) REFERENCES product_category (id)
) CHARACTER SET utf8mb4`

const ddlCreatePurchaseDetail = `
CREATE TABLE IF NOT EXISTS purchase_detail (
	purchase_id  INT UNSIGNED  NOT NULL,
	product_id   INT UNSIGNED  NOT NULL,
	qty          DECIMAL(20,4) NOT NULL,
	capital      DECIMAL(20,4) NOT NULL,
	price        DECIMAL(20,4) NOT NULL,
	position     INT UNSIGNED  NOT NULL DEFAULT '0',
	update_time  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (purchase_id, product_id),
	INDEX (purchase_id),
	INDEX (product_id),
	FOREIGN KEY purchase_detail_purchase_id_FK (purchase_id) REFERENCES purchase (id),
	FOREIGN KEY purchase_detail_product_id_FK (product_id) REFERENCES product (id)
) CHARACTER SET utf8mb4`

const ddlCreateEvent = `
CREATE TABLE IF NOT EXISTS event (
	id          INT UNSIGNED           NOT NULL AUTO_INCREMENT,
	name        VARCHAR(80)            NOT NULL,
	percentage  DECIMAL(20,4)          NOT NULL DEFAULT '0',
	rounding    DECIMAL(20,4) UNSIGNED NOT NULL DEFAULT '0',
	fixed_price DECIMAL(20,4)          NOT NULL DEFAULT '0',
	start       DATETIME               NOT NULL,
	end         DATETIME               DEFAULT NULL,
	deleted     TINYINT(1)             NOT NULL DEFAULT '0',
	delete_time DATETIME               DEFAULT NULL,
	update_time TIMESTAMP              NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
) CHARACTER SET utf8mb4`

const ddlCreateEventDetail = `
CREATE TABLE IF NOT EXISTS event_detail (
	event_id    INT UNSIGNED  NOT NULL,
	product_id  INT UNSIGNED  NOT NULL,
	qty         DECIMAL(20,4) NOT NULL,
	capital     DECIMAL(20,4) NOT NULL,
	price       DECIMAL(20,4) NOT NULL,
	deleted     TINYINT(1)    NOT NULL DEFAULT '0',
	delete_time DATETIME      DEFAULT NULL,
	update_time TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (event_id, product_id),
	INDEX (event_id),
	INDEX (product_id),
	FOREIGN KEY event_detail_event_id_FK (event_id) REFERENCES event (id),
	FOREIGN KEY event_detail_product_id_FK (product_id) REFERENCES product (id)
) CHARACTER SET utf8mb4`

const ddlCreatePaymentMethod = `
CREATE TABLE IF NOT EXISTS payment_method (
	id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
	name        VARCHAR(80)  NOT NULL,
	deleted     TINYINT(1)   NOT NULL DEFAULT '0',
	delete_time DATETIME     DEFAULT NULL,
	update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
) CHARACTER SET utf8mb4`

const ddlCreateDisplay = `
CREATE TABLE IF NOT EXISTS display (
	id            INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	employee_id   INT UNSIGNED  DEFAULT NULL,
	store_id      INT UNSIGNED  DEFAULT NULL,
	identifier    VARBINARY(20) DEFAULT NULL,
	transfer_time DATETIME      NOT NULL,
	qty           DECIMAL(20,4) NOT NULL,
	deleted       TINYINT(1)    NOT NULL DEFAULT '0',
	delete_time   DATETIME      DEFAULT NULL,
	update_time   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY display_identifier_UNIQUE (identifier),
	INDEX (employee_id),
	INDEX (store_id),
	FOREIGN KEY display_employee_id_FK (employee_id) REFERENCES employee (id),
	FOREIGN KEY display_store_id_FK (store_id) REFERENCES store (id)
) CHARACTER SET utf8mb4`

const ddlCreateDisplayDetail = `
CREATE TABLE IF NOT EXISTS display_detail (
	display_id   INT UNSIGNED  NOT NULL,
	product_id   INT UNSIGNED  NOT NULL,
	qty          DECIMAL(20,4) NOT NULL,
	update_time  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (display_id, product_id),
	INDEX (display_id),
	INDEX (product_id),
	FOREIGN KEY display_detail_sales_id_FK (display_id) REFERENCES display (id),
	FOREIGN KEY display_detail_product_id_FK (product_id) REFERENCES product (id)
) CHARACTER SET utf8mb4`

const ddlCreateSales = `
CREATE TABLE IF NOT EXISTS sales (
	id                INT UNSIGNED  NOT NULL AUTO_INCREMENT,
	employee_id       INT UNSIGNED  DEFAULT NULL,
	customer_id       INT UNSIGNED  DEFAULT NULL,
	store_id          INT UNSIGNED  DEFAULT NULL,
	payment_method_id INT UNSIGNED  DEFAULT NULL,
	identifier        VARBINARY(20) DEFAULT NULL,
	transaction_time  DATETIME      NOT NULL,
	qty               DECIMAL(20,4) NOT NULL,
	total_capital     DECIMAL(20,4) NOT NULL,
	total_sales       DECIMAL(20,4) NOT NULL,
	payment_due       DATE          DEFAULT NULL,
	paid_amount       DECIMAL(20,4) DEFAULT NULL,
	paid_date         DATE          DEFAULT NULL,
	deleted           TINYINT(1)    NOT NULL DEFAULT '0',
	delete_time       DATETIME      DEFAULT NULL,
	update_time       TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY sales_identifier_UNIQUE (identifier),
	INDEX (employee_id),
	INDEX (customer_id),
	INDEX (store_id),
	INDEX (payment_method_id),
	INDEX (transaction_time),
	FOREIGN KEY purchase_employee_id_FK (employee_id) REFERENCES employee (id),
	FOREIGN KEY purchase_customer_id_FK (customer_id) REFERENCES customer (id),
	FOREIGN KEY purchase_store_id_FK (store_id) REFERENCES store (id),
	FOREIGN KEY purchase_payment_method_id_FK (payment_method_id) REFERENCES payment_method (id)
) CHARACTER SET utf8mb4`

const ddlCreateSalesDetail = `
CREATE TABLE IF NOT EXISTS sales_detail (
	sales_id     INT UNSIGNED  NOT NULL,
	product_id   INT UNSIGNED  NOT NULL,
	qty          DECIMAL(20,4) NOT NULL,
	capital      DECIMAL(20,4) NOT NULL,
	price        DECIMAL(20,4) NOT NULL,
	update_time  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (sales_id, product_id),
	INDEX (sales_id),
	INDEX (product_id),
	FOREIGN KEY sales_detail_sales_id_FK (sales_id) REFERENCES sales (id),
	FOREIGN KEY sales_detail_product_id_FK (product_id) REFERENCES product (id)
) CHARACTER SET utf8mb4`

const ddlCreateCorrection = `
CREATE TABLE IF NOT EXISTS correction (
	id              INT UNSIGNED NOT NULL AUTO_INCREMENT,
	account_id      INT UNSIGNED DEFAULT NULL,
	description     VARCHAR(256) DEFAULT NULL,
	correction_time DATETIME     NOT NULL,
	deleted         TINYINT(1)   NOT NULL DEFAULT '0',
	delete_time     DATETIME     DEFAULT NULL,
	update_time     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	INDEX (account_id),
	FOREIGN KEY correction_account_id_FK (account_id) REFERENCES account (id)
) CHARACTER SET utf8mb4`

const ddlCreateCorrectionDetail = `
CREATE TABLE IF NOT EXISTS correction_detail (
	correction_id INT UNSIGNED  NOT NULL,
	product_id    INT UNSIGNED  NOT NULL,
	qty           DECIMAL(20,4) NOT NULL,
	deleted       TINYINT(1)    NOT NULL DEFAULT '0',
	delete_time   DATETIME      DEFAULT NULL,
	update_time   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (correction_id, product_id),
	INDEX (correction_id),
	INDEX (product_id),
	FOREIGN KEY correction_detail_correction_id_FK (correction_id) REFERENCES correction (id),
	FOREIGN KEY correction_detail_product_id_FK (product_id) REFERENCES product (id)
) CHARACTER SET utf8mb4`
