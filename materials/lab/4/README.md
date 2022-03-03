# Lab 4

## Required [10 points]
- Update wyoassign.go by specifically updating: `func UpdateAssignment`
- Update main.go to enable this new route
- Use postman to exercise/run through ALL of the available endpoints
- Save the responses from each run and place them in one text file


## Option 1
- Create a new set of endpoints for "classes" 
  - create new data structure(s) and 
  - create new endpoints (minimum create, get, delete)
  - update main.go 

## Option 2
- Modify/Harden the existing wyoassign.go endpoints
- Current endpoints lack real testing
- The POST / Create Assignment endpoint is terrible

## Option 3
- Create a wyoassign_test.go file which tests the functionality of the code

# Submission

CreateAssignment function has been updated to correctly return the status on
creation and log the call. UpdateAssignment has been updated and enabled. The
assignment with the id in the url is updated with the form values of the request.

Option 1 was chosen.

New endpoints are:
- GET /classes
- DEL /class/{id}
- POST /class
-- form values are strings with keys id, title, instr, desc, room.