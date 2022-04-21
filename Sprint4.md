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
1. Fixed issues in integration of login and sign up pages to ensure proper implementation.

![Screenshot (31)](https://user-images.githubusercontent.com/58160789/164346819-8a9f0b41-19d1-4598-8fee-f561d2cbee93.png)

After successfully logging in you will be displayed with your username in above way.Your progress will be saved with provided credentials.

![Screenshot (32)](https://user-images.githubusercontent.com/58160789/164346160-ad449c5d-1c98-445a-83e2-135c2797c6ee.png)

You can logout anytime in login page, For securing your account progress.

3. Updated code to use session id of user once he login or register the Prepboost.

![Screenshot (33)](https://user-images.githubusercontent.com/58160789/164348624-271352d0-01fe-42d8-8412-1a60ee293444.png)

The above is an example of signed up user named anudeepbar, with his credentials given he can successfully login after filling up signup details.

5. Updated code to get bookmark questions of individual users to frequently have a look of those questions.

![Screenshot (29)](https://user-images.githubusercontent.com/58160789/164350251-9ab0dbfa-0e55-4646-ad76-0771f96af660.png)

Successfully logged in users can see their bookmarked questions through their account details.

7. Added API links in respective yearwise and topic wise question pages to fetch the questions from backend accordingly.
 
![api link example](https://user-images.githubusercontent.com/58160789/164351088-4c02edbe-077f-4371-a231-b7134a27061b.png)

Here, we have written API links wherever necessary to fetch data from backend such as yearwise questions and topic wise questions and subtopics as well by using command fetch

![Screenshot (27)](https://user-images.githubusercontent.com/58160789/164351884-1c579b9d-bf80-4046-9714-bb9a88507d06.png)

Here in above image we can see we have fetched questions for yearwise questions as well as score is displayed at the bottom as we move on attempting questions. 

9. Updated the stylings like background colours,paddings,font sizes,alignment of the Prepboost.
10. Performed end to end Cypress testing starting from home page.

Features completed for updating Data base
1. Updation of questions  topic wise for past 10 years.
2. Updation of questions year wise for last 25 years.
3. Updated code at backend to extract the data of all the questions.
4. Initially we tried to update database directly at the database level but later, we tried to update it at backend level. Which demonstrates the levels of updation of code.
