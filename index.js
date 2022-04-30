import express from 'express'

const PORT = 5000;

const app = express()

app.use(express.json())

let todo = [
    {
        id: "1",
        name: "test",
        done: false,
    },
    {
        id: "2",
        name: "test2",
        done: false,
    },
    {
        id: "3",
        name: "tes3",
        done: false,
    }
]

app.get('/', (req, res) => {
    res.status(200).json(todo)
})

app.get('/:id', (req, res) => {
    const todo_item = todo.filter(item => item.id === req.params.id)

    res.status(200).json(todo_item[0])
})

app.post('/', (req, res) => {
    todo.push(req.body)

    res.status(200).json(todo)
})

app.put('/:id', (req, res) => {
    const id = req.params.id

    let todo_item = todo.find((item) => item.id === id);

    if (!todo_item) {
        res.status(404)
    }
    else {
        todo_item.name = req.body.name
        todo_item.done = req.body.done

        res.status(200).json(todo_item)
    }
})

app.delete('/:id', (req, res) => {
    const id = req.params.id

    let todo_item = todo.find((item) => item.id === id);

    if (!todo_item) {
        res.status(404)
    }
    else {
        todo = todo.filter(item => item.id !== id)

        res.status(200)
    }
})

app.listen(PORT, () => console.log('Server listened on port ' + PORT))