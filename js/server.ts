
import { Hono } from 'hono';

export const app = new Hono();

//构造几组数据
const todosmock = [{
  id: '1',
  title: 'Learn TypeScript',
  priority: 'HIGH',
  completed: false,
  content: 'Learn TypeScript and build awesome applications',
  createdAt: '2023-01-01T00:00:00.000Z'
},
{
  id: '2',
  title: 'Build a REST API',
  priority: 'MEDIUM',
  completed: false,
  content: 'Learn TypeScript and build awesome applications',
  createdAt: '2023-01-01T00:00:00.000Z'
},
{
  id: '3',
  title: 'Create a React App',
  priority: 'LOW',
  completed: false,
  content: 'Learn TypeScript and build awesome applications',
  createdAt: '2023-01-01T00:00:00.000Z'
}
]


app.get('/', (c) => {
  return c.json({
    ok: true,
    message: JSON.stringify(todosmock)
  })
})

export default app;
