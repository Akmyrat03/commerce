Commerce backend API

------------------------------------------------------------------------------------
1) User can sign-up. Every user have unique username, unique email, role default “user” and password. All fields are required. Password must be at least 4 character. After entering all data it hashes the password and save all data to database.
http://localhost:8000/api/users/sign-up

2) User can login. All fields are required. If user enters username and password correctly then user can login. If user has role “user” it will redirect to "http://localhost:8000/api/user/profile“.
If user has role “admin” it will redirect to "http://localhost:8000/api/admin/dashboard“
And token generated for user. 

3) User can sign-out. All fields are required. If user enters correct username and password then user can sign-out. It will blacklist token. Then we cannot use token 
"http://localhost:8000/api/users/sign-out“

4) User can view profile when signed in
"http://localhost:8000/api/users/profile


--------------------------------------------------------------------------------------
1) Admin can create a category. Category name field must not be empty.
http://localhost:8000/api/categories/add

2) Admin can update category by id. Category name field is required.
http://localhost:8000/api/categories/update/id

3) Admin can delete category by id.
http://localhost:8000/api/categories/delete/id

4) View all categories.
http://localhost:8000/api/categories/view


--------------------------------------------------------------------------------------
1) Admin can create products
http://localhost:8000/api/products/add-product

2) Admin see all drafted and published status products
http://localhost:8000/api/products/view-all

3) Users can see just published status products
http://localhost:8000/api/products/published

4) We can group products by category name
http://localhost:8000/api/products/view/name

5) Admin can delete products by id
http://localhost:8000/api/products/delete/id

6) Admin can update products by id
http://localhost:8000/api/products/update/id


--------------------------------------------------------------------------------------
Cart section

http://localhost:8000/api/cart/add-cart

http://localhost:8000/api/cart/get/id

http://localhost:8000/api/item/add-cart-item

http://localhost:8000/api/item/get-all/id

-------------------------------------------------------------------------------------
Order section

http://localhost:8000/api/orders/add-order

http://localhost:8000/api/orders/view-all

