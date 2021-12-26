## Database Table Structure
 -------------------------------
 
 ## client -> customer, contact,
 - id
 - name
 - mobile
 - email
 - address 
 - status


 ## product-> handset, charger, 
 ## service/repair: 
 
-  item_group:
- id
- group_name
- description
- status
 
 ## item: 
 - id
 - item_name: SAMSUNG-GALAXY13
 - item_group_id -> CHARGER
 - price: 100.50 --> NUMERIC
 - warranty-> yes/no -> 1,0
 - warranty_period: 30 days
 - stock -> integer
 - status -> 1,0
 
 
 ## order: invoice
 - id
 - order_number
 - order_date
 - delivery_date
 - item
 - quantity
 - price
 - client_id
 - order_discount
 - order_total
 - status

 ## client -> customer, contact,
 - id
 - name
 - mobile
 - email
 - address 
 - status
 
 ## product-> handset, charger, 
 ## service/repair: 
 
 - item_group:
 - id
 - group_name
 - description
 - status
 
 ## item: 
 - id
 - item_name: SAMSUNG-GALAXY13
 - item_group_id -> CHARGER
 - price: 100.50 --> NUMERIC
 - warranty-> yes/no -> 1,0
 - warranty_period: 30 days
 - stock -> integer
 - status -> 1,0
 
 
 ## order: invoice
 - id
 - order_number
 - order_date
 - delivery_date
 - item
 - quantity
 - price
 - client_id
 - order_discount
 - order_total
 - status
 
 
 ```
 CREATE TABLE "client" (
	"id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL,
	"mobile"	TEXT,
	"email"	TEXT,
	"address"	TEXT,
	"status"	INTEGER NOT NULL DEFAULT 1,
	PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE "item_group" (
	"id"	INTEGER NOT NULL,
	"group_name"	TEXT NOT NULL UNIQUE,
	"description"	TEXT,
	"status"	INTEGER NOT NULL DEFAULT 1,
	PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE "item" (
	"id"	INTEGER NOT NULL,
	"item_name"	TEXT NOT NULL UNIQUE,
	"item_group_id"	INTEGER NOT NULL,
	"price"	NUMERIC NOT NULL DEFAULT 0.00,
	"warranty"	INTEGER NOT NULL DEFAULT 0,
	"warranty_period"	INTEGER,
	"stock"	INTEGER NOT NULL DEFAULT 0,
	"status"	INTEGER NOT NULL DEFAULT 1,
	PRIMARY KEY("id"),
	FOREIGN KEY("item_group_id") REFERENCES "item_group"("id")
);

CREATE TABLE "invoice" (
	"id"	INTEGER NOT NULL,
	"invoice_number"	TEXT NOT NULL,
	"invoice_date"	TEXT,
	"delivery_date"	TEXT,
	"item_id"	INTEGER NOT NULL,
	"quantity"	INTEGER NOT NULL DEFAULT 1,
	"price"	NUMERIC NOT NULL DEFAULT 0.00,
	"client_id"	INTEGER NOT NULL,
	"invoice_discount"	NUMERIC,
	"invoice_total"	NUMERIC NOT NULL DEFAULT 0,
	"status"	INTEGER NOT NULL DEFAULT 1,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("item_id") REFERENCES "item"("id"),
	FOREIGN KEY("client_id") REFERENCES "client"("id")
);

CREATE TABLE "company" (
	"id"	INTEGER NOT NULL,
	"company_name"	TEXT NOT NULL,
	"address"	TEXT,
	"website"	TEXT,
	"email"	TEXT,
	"mobile"	TEXT,
	"logo"	TEXT,
	"status"	INTEGER NOT NULL DEFAULT 1,
	PRIMARY KEY("id" AUTOINCREMENT)
);
```


 ## Mysql Keywords
 * https://dev.mysql.com/doc/refman/8.0/en/keywords.html
 * https://sqlitebrowser.org/
