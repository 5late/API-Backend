# Endpoints

## ``/people``

- Get JSON object of all people
- Returns this JSON object structure:
```

{
    "id": 1,
    "firstname": "John",
    "lastname": "Doe",
    "birthdate": "Jan 1, 1970",
    "age": 51
}

```

## ``/person``

- POST a new person 
- Must follow this JSON object structure:
```

{
    "id": 1,
    "firstname": "John",
    "lastname": "Doe",
    "birthdate": "Jan 1, 1970",
    "age": 51
}

```

## ``/createApp``

- POST a new appointment
- Must follow this JSON object structure:
```

{
    "id": 1,
    "firstname": "John",
    "lastname": "Doe",
    "date": "13/2/23",
    "time": "9:20",
    "reason": "Shoulder pain.",
    "discordid": ""
}

```

## ``/getApp/{id}``

- GET an appointment for a person based on their ID
- Will return JSON object that follows:
```

{
    "id": 1,
    "firstname": "John",
    "lastname": "Doe",
    "date": "13/2/23",
    "time": "9:20",
    "reason": "Shoulder pain.",
    "discordid": ""
}

```