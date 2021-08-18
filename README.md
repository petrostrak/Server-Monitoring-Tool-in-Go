# Server Monitoring Tool in Go
Interactive and responsive application using WebSockets and Go

`WebSockets` are a technology that allows developers to build highly interactive, extremely fast web applications. Since WebSockets are supported by all modern browsers, and have been for years now, there is no reason why they cannot be implemented on any modern web application. This application shows how to work with WebSockets with Go.

`監視するために (Kanshi suru tame ni)` is a web application which `monitors` remote servers and hosts, and notifies us in real-time when a service goes down or comes back up again. It will send notifications in real time using WebSockets (which will update the appropriate content on the pages being viewed by all connected clients), by email, and, as a bonus, it'll also send notifications using `text messages` (SMS) with  `Twilio`.
