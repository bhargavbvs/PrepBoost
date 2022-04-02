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
</br>Backend :  

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
