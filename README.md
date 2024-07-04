Razor pay api integration:


### example

This is an example of how to use and give json data using api .
* Payload:
  ```sh
    {
    "user_key": "rzp_test_vOsKKSWnOE803Q",
    "email": "payeemail@mail.com",
    "name": "payeeName",
    "amount": 2500,
    "contact": "8787874562",
    "currency": "INR",
    "receipt": "receipt_data",
    "org_name": "org_name",
    "description": "description of service",
    "img_url": "imgUrlOfOrg"
    }
  
  ```

* Endpoint:
  ```sh
  "POST":http://localhost:8085/order ; 
  resp: 
    {
    "payUrl": "localhost:8085/pay/order_OUUWizlFwiIN1p"
    }
  
  ```
