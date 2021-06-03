
My code isn't ready to production, you can't start it, because Repository, Notifier and Cmd should be implemented. 
There are some examples, how I used to implement the similar ones in specific folders.

I decided to spend more time to make flexible microservice, based on Clean Architecture. This service is easy to be scaled.

I chose some approaches:
* project is based on Clean Architecture;
* users.go file contains main domain model and intefaces;
* use ulid, instead of id - protect against brute force;
* use soft delete, because we can have a lot of relations with our user;
* use Lock for User changes;
* search is based on keyset pagination (without offset), based on Postgresql b-tree - CreatedAt field;
* mock generator for interfaces;
* use middlewares for logging;
* don't log user password;

Improvements:
* add more tests;
* implement other Repository, Notifier and cmd;
* add API docs, use SWAGGER;
