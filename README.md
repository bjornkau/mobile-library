# mobile-library
This repo contains the server side implementation for a library tracking application.
It is supposed to handle information about books from https://openlibrary.org/dev/docs/api/books

Specification:

calls:

Book API
  
  POST
  root/library/api/book
  
  POST
	root/library/api/author
	
User API

  POST
  root/library/users/login
  
  POST
  root/library/users/register
  
  GET
  root/library/users/logout
	
  GET
  root/library/users/books
	
  GET
  root/library/users/authors
	
  POST
  root/library/users/registerbook
  
  These calls shall be used together with a frontend that is to be developed.
  
  ##things that went well
  Code was fast to develop. All calls give apropriate responses.
  
  ##Things that did not go well
  I spent way to much time to working on manually parsing the json from openlibrary.
  There was not enough time to fix the json web tokens I wanted to use for user validation.
  Openlibrary does not have an api with information about authors, will need to look into additional sources for information
  
