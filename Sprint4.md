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

DEMONSTRATION VIDEO

https://user-images.githubusercontent.com/58160789/164365619-497292fb-5492-4e8c-a088-74312af9d44f.mp4

CYPRESS TEST VIDEO

https://user-images.githubusercontent.com/58160789/164365680-2ab965d4-d22f-4bf9-9d2f-7de24d1555c2.mp4

Features completed in Sprint4 for BACKEND:

1. Updation of questions  topic wise for past 10 years.
2. Updation of questions year wise for last 25 years.
3. Updated code at backend to extract the data of all the questions.
4. Initially we tried to update database directly at the database level but later, we tried to update it at backend level. Which demonstrates the levels of updation of code.
![Screenshot (36)]![image](https://user-images.githubusercontent.com/93815375/164362221-b4b96a59-83b8-4632-bee2-3c8d523efcfc.png)
Link to script_for_db_update : https://github.com/bhargavbvs/PrepBoost/blob/main/script_for_db_update.py

5. questions  topic wise for past 10 years
![Screenshot (34)]![image](https://user-images.githubusercontent.com/93815375/164361438-187ec519-f556-4aaa-b40b-bfc8c2783ea1.png)

6. questions year wise for last 25 years
![Screenshot (35)]![image](https://user-images.githubusercontent.com/93815375/164361886-29a521ca-2b95-4c6e-afe5-b4c7d1594208.png)

7. Worked on writing Unit test cases for LeaderBoard
<img width="1182" alt="Screen Shot 2022-04-20 at 11 46 21 PM" src="https://user-images.githubusercontent.com/17181838/164368018-3244315b-9daf-436d-b9fc-bdade8686205.png">
<img width="1181" alt="Screen Shot 2022-04-20 at 11 46 34 PM" src="https://user-images.githubusercontent.com/17181838/164368054-feccb250-7e07-4a1f-9205-751f6ec0433d.png">

8. Worked on writing Unit test cases for Bookmarked questions
<img width="1176" alt="Screen Shot 2022-04-20 at 11 47 30 PM" src="https://user-images.githubusercontent.com/17181838/164368165-bb7fc50a-e3e9-4787-869d-0c3ed3c5c359.png">

9. Worked on writing Unit test cases for topic wise questions 
<img width="1185" alt="Screen Shot 2022-04-20 at 11 47 53 PM" src="https://user-images.githubusercontent.com/17181838/164368207-add64fa3-1b96-4d85-9bd6-1d0aa984b02c.png">

10. Worked on writing Unit test cases for Yearwise questions
<img width="1176" alt="Screen Shot 2022-04-20 at 11 48 39 PM" src="https://user-images.githubusercontent.com/17181838/164368287-46657fe8-8630-4b8e-8742-70bcc09ec77e.png">

11. Unit test cases for the User updates
<img width="1166" alt="Screen Shot 2022-04-20 at 11 49 49 PM" src="https://user-images.githubusercontent.com/17181838/164368390-14d666c1-2447-4527-95eb-73a06be806f1.png">

12. Unit test cases fot the User signup
<img width="1195" alt="Screen Shot 2022-04-20 at 11 50 17 PM" src="https://user-images.githubusercontent.com/17181838/164368443-1d4ec5b3-df9b-4ab4-8668-ff18d6fc536f.png">

13. Updated the data in the RDS Database instance in the AWS Servers. Contains the data of the past 25 years. 
<img width="1270" alt="Screen Shot 2022-04-21 at 12 15 02 AM" src="https://user-images.githubusercontent.com/17181838/164370877-50827f02-88e3-44f4-8d69-5b3c9256cddc.png">

14. Fixed the ASCII conversion error while inserting the data into the MySQL database of RDS : UTF8 errors

15. Wrote the python script for the data upload into the server database. 
<img width="1123" alt="Screen Shot 2022-04-21 at 12 19 19 AM" src="https://user-images.githubusercontent.com/17181838/164371302-3fa780ac-4ee9-4561-9f4c-8f0f4bac91cd.png">

16. Collected and modified the data of the past 25 years questions for the UPSC prelims questions. 
<img width="1696" alt="Screen Shot 2022-04-21 at 12 20 49 AM" src="https://user-images.githubusercontent.com/17181838/164371493-a6233a09-37c3-478e-8d8d-31bc7477d1b6.png">

17. Wrote the API for report question error
18. Updated the leaderboard API 
19. Fixed the bugs for the sql time error - Convert uint32 to int error
20. Fixed the testing bugs for the unit test cases error
21. Fixed the sql query errors in the topicwise, yearwise and bookmarked questions


https://user-images.githubusercontent.com/17181838/164367047-d8b4ed6e-df44-4934-b6de-c924a9d1e252.mp4

