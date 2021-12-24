 
 1640041730
 
 2,147,483,647
 
 integer ->
 
 status --> Running /  inactive / suspended
 
 status --> 1  0  2
 -------------------------------
 
 client -> customer, contact,
 id
 name
 mobile
 email
 address 
 status
 
 product-> handset, charger, 
 service/repair: 
 
 item_group:
 id
 group_name
 description
 status
 
 item: 
 id
 item_name: SAMSUNG-GALAXY13
 item_group_id -> CHARGER
 price: 100.50 --> NUMERIC
 warranty-> yes/no -> 1,0
 warranty_period: 30 days
 stock -> integer
 status -> 1,0
 
 
 order: invoice
 id
 order_number
 order_date
 delivery_date
 item
 quantity
 price
 client_id
 order_discount
 order_total
 status
