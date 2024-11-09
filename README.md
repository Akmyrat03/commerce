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
http://localhost:8000/api/categories/update/3
3) Admin can delete category by id.
http://localhost:8000/api/categories/delete/2
4) View all categories.
http://localhost:8000/api/categories/view
