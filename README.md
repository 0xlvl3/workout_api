# Workout API


## What is this? 

*This is a side project to learn how to build a working API*

The goal is to create an API where a user can authenticate, have there own workouts 
which consists of exercises that are under a catergory such as 'chest' for chest exercises 
or 'back' for back exercises.

User -> Workout -> []exercises

The workouts created by one user will be added to that users Workouts, which they will 
be able to pull up with a unique ID. 

------
### Calls

Users:

Workout:

Excerises:


## Dependencies 
I will include a script that if you `./dependencies.sh` you will automatically install dependencies.
Check `scripts` for that script.
------

***MongoDB***

`go get go.mongodb.org/mongo-driver/mongo`

***Go Fiber***

`go get github.com/gofiber/fiber/v2`

***Bcrypt***

`go get golang.org/x/crypto/bcrypt`

***JWT***

`go get github.com/golang-jwt/jwt/v5`
