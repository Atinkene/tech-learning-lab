const mysql = require('mysql2/promise'); 
require('dotenv').config();

const pool = mysql.createPool({
  host: process.env.DB_HOST,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  port: process.env.DB_PORT
});

pool.on('error', (err, client) => {
    console.error("Erreur de connection a la base de donnees");
    process.exit(-1);
})


module.exports = pool;