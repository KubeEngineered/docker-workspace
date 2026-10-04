const express = require('express');
const { Pool } = require('pg');

const app = express();
const PORT = process.env.PORT || 3000;

// Database connection pool using environment variables
const pool = new Pool({
  host: process.env.DB_HOST || 'localhost',
  port: process.env.DB_PORT || 5432,
  user: process.env.DB_USER || 'postgres',
  password: process.env.DB_PASSWORD || 'postgres',
  database: process.env.DB_NAME || 'devops_db'
});

app.get('/', async (req, res) => {
  try {
    const dbResult = await pool.query('SELECT NOW()');
    res.json({
      status: 'running',
      message: 'Node.js connected to PostgreSQL databasesuccessfully!',
      dbTime: dbResult.rows[0].now
    });
  } catch (err) {
    res.status(500).json({
      status: 'error',
      message: 'Failed to establish connection to PostgreSQL database',
      error: err.message
    });
  }
});

app.listen(PORT, () => {
  console.log(`App Server listening on port ${PORT}`);
});
