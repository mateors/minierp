## Database Table Create {class 02 (24-12-2021)}

### Database Table Structure
client -> customer, contact,

    id
    name
    mobile
    email
    address
    status
   ![client](https://github.com/Fatema04/minierp/blob/main/assignment/02_Database_structure/images/sqlite/img1.png?raw=true)

### item_group ( product-> handset, charger, service / repair )

    id
    group_name
    description
    status
![item_group](https://github.com/Fatema04/minierp/blob/main/assignment/02_Database_structure/images/sqlite/img3.png)

### item:

    id
    item_name: SAMSUNG-GALAXY13
    item_group_id -> CHARGER
    price: 100.50 --> NUMERIC
    warranty-> yes/no -> 1,0
    warranty_period: 30 days
    stock -> integer
    status -> 1,0

![item](https://github.com/Fatema04/minierp/blob/main/assignment/02_Database_structure/images/sqlite/img4.png)


### invoice (order)

    id
    invoice_number
    invoice_date
    delivery_date
    item_id
    quantity
    price
    client_id
    invoice_discount
    invoice_total
    status

![invoice](https://github.com/Fatema04/minierp/blob/main/assignment/02_Database_structure/images/sqlite/img5.png)

### Company

    id
    company name
    address
    website
    email
    mobile
    logo
    status

![company](https://github.com/Fatema04/minierp/blob/main/assignment/02_Database_structure/images/sqlite/img2.png)

