## SQL Queries

```sql
INSERT INTO item_group (group_name,description,status) VALUES("Handset", "mobile handset",1);
```

 > SELECT * from  item_group;

```sql
INSERT INTO item (item_name,item_group_id,price,warranty,warranty_period,status) VALUES("Samsung Galaxy A5",1,15000,1,"1 Year",1);
```

> SELECT * FROM item;

```sql
INSERT INTO invoice (invoice_number,invoice_date,delivery_date,item_id,quantity,price,client_id) VALUES("INVSL#00001","2022-02-04", "2022-02-04", 1,1,15000,2);
```

> SELECT * FROM invoice;

```sql
SELECT invoice.invoice_number,invoice.invoice_date,invoice.price,invoice.quantity,
item.item_name,
client.name as client_name,client.mobile
 FROM invoice
LEFT JOIN item ON invoice.item_id=item.id 
LEFT JOIN client ON invoice.client_id=client.id
WHERE invoice.status=1;
```

## Learning Resource
* [SQL Join](https://www.w3schools.com/sql/sql_join.asp)
* [Join Multiple table](https://www.sqlshack.com/learn-sql-join-multiple-tables/)