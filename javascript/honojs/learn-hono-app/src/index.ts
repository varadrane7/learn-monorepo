import { serve } from '@hono/node-server'
import { Hono } from 'hono'
import { randomUUID } from 'crypto'
import { streamText } from 'hono/streaming'

let videos :any = []

const app = new Hono()

app.get('/', (c) => {
  return c.html('<h1>Welcome to hono crash course</h1>')
})

app.post('/videos', async (c) => {
  const {videoName, channelName, duration} = await c.req.json()
  const newVideo = {
    id: randomUUID(),
    videoName,
    channelName,
    duration
  }

  videos.push(newVideo)

  return c.json(newVideo)
})

// Read All objects (using stream api)
app.get('/videos/all', async (c) => {
  return streamText(c, async(stream) => {
    await stream.write(JSON.stringify(videos))
  })
})

// read by id
app.get('/videos/:id', async(c) => {
  const { id } = c.req.param()
  const video = videos.find((video: any) => video.id === id)
  if (!video)
    return c.json({message: "Video not found"}, 404)
  return c.json(video)
})

//upate by id
app.put('/video/:id', async (c) => {
  const { id } = c.req.param()
  const index = videos.findIndex((video: any) => video.id === id)
  if (index === -1)
    return c.json({message: "Video not found"}, 404)

  const {videoName, channelName, duration} = await c.req.json()
  //update the element
  videos[index] = {...videos[index], videoName, channelName, duration}
  return c.json(videos[index])
})

// delete by id
app.delete('/videos/:id', async (c) => {
  const {id} = c.req.param()
  videos = videos.filter((video: any) => video.id !== id )
  return c.json({message: "Video Deleted"})
})

const port = 3000
console.log(`Server is running on port ${port}`)

serve({
  fetch: app.fetch,
  port
})
