# Notable Take Home Test

To access the files from my computer, go to Ubuntu terminal and type "explorer.exe ." and take the files from there.

## Welcome to my take home assignment
To run this code, you need to have go set up and working in your environment.

Use the following command to start it:
`go run main.go`

Now, open a new terminal and follow along the following steps to test the functionality.

## First, add a doctor:
curl -X POST 'http://localhost:8080/doctor/add' -H 'Content-Type: application/json' -d '{"first_name": "Robert", "last_name": "Malmstein"}'

## Then try getting doctors:
curl -X GET http://localhost:8080/doctors

## Add another doctor:
curl -X POST 'http://localhost:8080/doctor/add' -H 'Content-Type: application/json' -d '{"first_name": "Tim", "last_name": "Lin"}'

## Get a specific doctor:
curl -X GET http://localhost:8080/doctor?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8

## Then try getting all doctors again:
curl -X GET http://localhost:8080/doctors

## Then, try getting all appointments for the doctor (use the UUID from above to query this):
curl -X GET 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8&date=2023-04-05'

## Then, add an appointment to the new doctor:
curl -X POST 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8' -H 'Content-Type: application/json' -d '{"start_time": "2023-04-05 09:15:00", "end_time": "2023-04-05 09:30:00", "patient": {"first_name": "John", "last_name": "Doe"}, "kind": "New Patient"}'

## Try adding an appointment that isn't in a 15 minute interval:
curl -X POST 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8' -H 'Content-Type: application/json' -d '{"start_time": "2023-04-05 09:20:00", "end_time": "2023-04-05 09:30:00", "patient": {"first_name": "John", "last_name": "Doe"}, "kind": "New Patient"}'
<Invalid appointment time>

## Try adding more than 3 to the same time slot
curl -X POST 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8' -H 'Content-Type: application/json' -d '{"start_time": "2023-04-05 09:15:00", "end_time": "2023-04-05 09:30:00", "patient": {"first_name": "Jane", "last_name": "Doe"}, "kind": "New Patient"}'

curl -X POST 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8' -H 'Content-Type: application/json' -d '{"start_time": "2023-04-05 09:15:00", "end_time": "2023-04-05 09:30:00", "patient": {"first_name": "Susy", "last_name": "Doe"}, "kind": "New Patient"}'

curl -X POST 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8' -H 'Content-Type: application/json' -d '{"start_time": "2023-04-05 09:15:00", "end_time": "2023-04-05 09:30:00", "patient": {"first_name": "BamBam", "last_name": "Doe"}, "kind": "New Patient"}'
<Maximum number of appointments reached for this doctor>

## Then add a new appointment to 9:30:
curl -X POST 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8' -H 'Content-Type: application/json' -d '{"start_time": "2023-04-05 09:30:00", "end_time": "2023-04-05 09:30:00", "patient": {"first_name": "Wowzer", "last_name": "Bowser"}, "kind": "New Patient"}'

## Try deleting an appointment (use the get request earlier to get the id's of appointments):
curl -X DELETE 'http://localhost:8080/appointments?doctor_id=c1b6a4f5-dccb-458a-a506-2add802fe5b8&appointment_id=ebf4ceb3-4dd3-4ff1-9d7c-6da738ef4035'

## Get all doctors again to see the changes:
curl -X GET http://localhost:8080/doctors