# Project - PrepBoost

# Frontend server
To run the Froentend Server run the given command below :
>`cd /frontend/prepboost/`

>`npm start`

# Backend server
> `cd ../backend`

> `go run main.go `

First the localhost:4200 will start upon ng serve and then a new tab for cypress test will also open upon running the cypress related commands.

## Sprint 2

Link to the video presentation for Sprint 2

Frontend : https://github.com/bhargavbvs/PrepBoost/blob/main/video1069825702.mp4
</br>Backend :  Added the screenshots below.

FRONTEND DEVELOPMENT
Features completed in Sprint1:
1. Integration with SQLite and GORM along with the Mux router for handlers.
2. React-router-dom to create user interface framework.
3. Home API - localhost:9000/home
4. Login API - localhost:9000/login

Deliverables :: 

1. Users should be able to get the questions along with the answers they marked of that specific year and topic.
2. Users can also bookmark questions and get them so they can reattempt those questions multiple times.
3. JWT authentication has been and the user session has been set for upto 24 hours.
4. 10 years Questions data is collected and are categorized according to the specific topic and subtopic. 


## 4. Unit Test Cases for the Backend APIs

Unit test cases are written in the specific controller folders.
Go to the specific folder and run the test cases using the command below.

> `go test -v`

Unit test case results for the backend APIs :

1. Unit test case for Fetch user endpoint : PASSED
<img width="1275" alt="Screen Shot 2022-04-01 at 8 51 08 PM" src="https://user-images.githubusercontent.com/17181838/161362775-34fa1032-4b1a-4f09-9a73-a63b40d83ec2.png">

2. Unit test case for Login of the User : PASSED
<img width="1262" alt="Screen Shot 2022-04-01 at 8 48 09 PM" src="https://user-images.githubusercontent.com/17181838/161362818-ac7e9e3d-d61d-4d3f-9f0e-9dd0a949977e.png">


3. Unit test case for updating the 

4. Added the Bookmarks API so that user can get the questions bookmarked. 
<img width="1408" alt="Screen Shot 2022-04-01 at 11 46 45 PM" src="https://user-images.githubusercontent.com/17181838/161365060-2421d831-85e6-411a-b7a3-5d6e8c7c1004.png">

5. Added the API to get questions yearwise.
<img width="1417" alt="Screen Shot 2022-04-01 at 11 51 24 PM" src="https://user-images.githubusercontent.com/17181838/161365135-7273aa2e-d970-4419-bd58-13df7f01aa64.png">

6. Added the API to get topicwise questions. 
<img width="1420" alt="Screen Shot 2022-04-01 at 11 54 29 PM" src="https://user-images.githubusercontent.com/17181838/161365240-91bcc4b4-09c9-4e25-a324-55ff0bd263ae.png">

7. Database migration from SQLite to MySql for multiple user access and scalability
<img width="1531" alt="Screen Shot 2022-04-01 at 11 58 18 PM" src="https://user-images.githubusercontent.com/17181838/161365365-8451ba66-c214-4e3e-bf79-144ea2684413.png">


# What we accomplished 
1. Functions
2. Created the questions API to get the year wise specific questions
3. Created the bookmark API to get bookmark questions of specific student
4. Database migration from SQLite to MySql for multiple user access and scalability
5. Create User session management
6. Create bookmarks table for the user to store bookmarked questions
7. Changed navigation bar from horizontal to vertical.
8. Implemented unit testing using jest.
9. Added login and signup pages for user to login and signup.
10. Installed axios to integrate front end with backend.
11. Integrated frontend with backend for login page component to fetch user details from backend.
12. Added bookmarks page where user can see the questions he/she bookmarked to visit later.
13. Implemented automated testing using cypress for various pages.
