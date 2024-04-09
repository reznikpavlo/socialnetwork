# Social Netowork project on golang + Vuejs

---

On this project used go 1.22 and VueJs 3+


---
### Additional commands for testing

#### GET all
fetch('/message/').then(response => response.json().then(console.log))

#### GET one
fetch('/message/2').then(response => response.json().then(console.log))

#### POST add new one
fetch(
'/message',
{
method: 'POST',
headers: { 'Content-Type': 'application/json' },
body: JSON.stringify({ text: 'Fourth message (4)', id: 10 })
}
).then(result => result.json().then(console.log))

#### PUT save existing
fetch(
'/message/4',
{
method: 'PUT',
headers: { 'Content-Type': 'application/json' },
body: JSON.stringify({ text: 'Fourth message', id: 10 })
}
).then(result => result.json().then(console.log));

#### DELETE existing
fetch('/message/4', { method: 'DELETE' }).then(result => console.log(result))
