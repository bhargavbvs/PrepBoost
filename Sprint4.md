# Frontend server
To run the Froentend Server run the given command below :
>`cd /frontend/prepboost/`

>`npm start`

# Backend server
> `cd ../backend`

> `go run main.go `

First the localhost:4200 will start upon ng serve and then a new tab for cypress test will also open upon running the cypress related commands.

# Sprint 4
Features completed in Sprint4 for FRONTEND:
1. Fixed issues in integration of login and sign up pages to ensure proper implementation.\n
![Screenshot (32)](https://user-images.githubusercontent.com/58160789/164346160-ad449c5d-1c98-445a-83e2-135c2797c6ee.png)

3. Updated code to use session id of user once he login or register the Prepboost.
4. Updated code to get bookmark questions of individual users to frequently have a look of those questions.
5. Added API links in respective yearwise and topic wise question pages to fetch the questions from backend accordingly.
6. Updated the stylings like background colours,paddings,font sizes,alignment of the Prepboost.
7. Performed end to end Cypress testing starting from home page.

Features completed for updating Data base
1. updation of questions  topic wise for past 10 years.
2. updation of questions year wise for last 25 years.
3. Updated code at backend to extract the data of all the questions.
