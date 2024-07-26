import { Server } from "socket.io";

const io = new Server({
  cors: {
    "origin": "*"
  }
});

io.on("connection", socket => {
  socket.on("join", pollId => {
    socket.join(pollId)
  });

  socket.on("update", pollId => {
    io.to(pollId).emit("update")
  });
});

io.listen(4001);
