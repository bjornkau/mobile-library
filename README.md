# mobile-library
This repo contains the server side implementation for a library tracking application.
It is supposed to handle information about books from https://openlibrary.org/dev/docs/api/books

Specification:

calls:

Book API
  
  POST
  root/library/api/book
  Request: {"isbn":isbn_value_is_int}
  
  POST
	root/library/api/author
  Request: {"name":"name_of_author"}
	
User API

  POST
  root/library/users/login
  Request: {"username":"username","password":"password"}
  
  POST
	root/library/users/register
  Request: {"name":"name_of_author"}
  
  GET
	root/library/users/logout
	
  GET
  root/library/users/books
	
  GET
  root/library/users/authors
	
  POST
  root/library/users/registerbook
  Request: {"name":"name_of_author"}
  
  
