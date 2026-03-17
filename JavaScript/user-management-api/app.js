const express = require('express');
const app = express();
const userRoutes = require('./routes/userRoutes')
require('dotenv').config()

app.use(express.json());

app.use('/api/users', userRoutes);

app.listen(process.env.APP_PORT, () => {
    console.log("Server run on port ", process.env.APP_PORT)
})