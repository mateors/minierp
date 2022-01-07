# Manual Data Entry to sqlite

## Insert Data

```sql
INSERT INTO client(name, mobile, email, address) VALUES("Md Abu Raihan", "01853566901", "raihanmahmudi35@gmail.com", "Kushtia");
```

## Delete data

```sql
DELETE FROM client;
```

## Delete data with condition

```sql
DELETE FROM client WHERE id=2;
```

## Data Insert to item_group table

```sql
INSERT INTO item_group(group_name, description) VALUES("Charger","samsung");
```

## Data Insert to item table

```sql
INSERT INTO item(item_name, item_group_id, price) VALUES("Sumsung A50", 1, 200.33);
```

## Data Insert to invoice table

```sql
INSERT INTO invoice(invoice_number, invoice_date, item_id, quantity, price, client_id, invoice_total) VALUES ("23333", "2022-01-07", 1, 1, 200.00, 1, 100.00);
```

## Fetch Data

```sql
SELECT * FROM invoice i,client c WHERE i.client_id=c.id;
```
