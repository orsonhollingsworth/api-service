const { createClient } = require('redis');
const { Client } = require('pg');

const parser = {
  async parseQuery(sql) {
    const client = createClient();
    client.on('error', (err) => console.error('Redis Client Error', err));
    const query = await client.query(`SELECT * FROM pg_stat_statements WHERE query = $1`, [sql]);
    const stats = query.rows[0];
    return {
      query,
      duration: stats.total_time,
      calls: stats.calls,
      plan: stats.query_plan,
    };
  },

  async parseSql(sql) {
    const client = new Client();
    client.on('error', (err) => console.error('PostgreSQL Client Error', err));
    client.connect();
    return client.query(sql);
  },
};

module.exports = parser;