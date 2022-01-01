## SQL queries
 

### Data inster to Table `Client`
```sql
INSERT INTO client(name,mobile,email,address) VALUES("NASARUL HASAN", "01552457194","nasarulhasan@gmail.com","Dhaka");

INSERT INTO client(name,mobile,email,address) VALUES("NASARUL MIA", "01552457195","nasarulhasan2@gmail.com","Dhaka");

INSERT INTO client(name,mobile,email,address) VALUES("TAREQ IBNE RAHMAN", "01700000000","DRTAREQ@gmail.com","Dhaka");

INSERT INTO client(name,mobile,email,address) VALUES("Eshita", "01600000000","eshita@gmail.com","Dhaka");

```
### Data inster to Table `item_group`
```sql
INSERT INTO item_group (group_name,description) VALUES("Charger", "samsung phone charger");
```

### Data inster to Table `item`
```sql
INSERT INTO item (item_name,item_group_id,price) VALUES("Galaxy A50", 2, 200.50);
```
### Data inster to Table `invoice`

```sql
INSERT INTO invoice (invoice_number,invoice_date,item_id,quantity,price,client_id,invoice_total) VALUES("ORDER#1","2021-12-31",1,5,200.00,1, 1000.00)
```
### if you need sqlite_sequence change then use this command  
```sql
UPDATE sqlite_sequence SET seq=0 WHERE name="client";
```



### Data fetch
```sql
SELECT invoice.*,client.name, item.item_name 
FROM invoice,client,item 
WHERE invoice.client_id=client.id AND invoice.item_id=item.id;
```
```sql
SELECT invoice.*,client.name, item.item_name, item.price, item_group.group_name 
FROM invoice,client,item,item_group 
WHERE invoice.client_id=client.id AND invoice.item_id=item.id AND item.item_group_id=item_group.id;
```
### Table delete 
```sql
DELETE FROM client;

DELETE FROM client WHERE id=1;
```

-----------------------------------------------

### Programming by
- golang
### Database
 - sql,nosql
 - sql
 - mssql
 - sqlite, postgrey

-----------------------------------------------

### Front-end - `html5/css and javascript`
-----------------------------------------------
### Desktop APP `fyne.io`


## SQL Query Tutorial
* [Youtube video lecture](https://www.youtube.com/watch?v=7S_tz1z_5bA&t=7469s)

