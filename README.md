# AI-Powered Chat Server

This repository contains the source code for an AI-powered chat server built using the Gin framework in Go. The server handles multiple API endpoints with a focus on authentication, user management, and especially the `/chat` endpoint which utilizes large language models (LLMs) like OpenAI, Mistral, Anthropic, and Groq API with open models to analyze user messages and return appropriate responses.

This server has been created to work as the backend of this electron application : https://github.com/Chadi00/milio_desktop

## Key Endpoints
**Auth and User Management**:
POST /login and POST /register for user authentication.
DELETE /user/delete for user account deletion.

**Chat and Response Generation**:
POST /chat for processing user requests and generating responses or action codes.
POST /chat/stream for streamed text completions using Anthropic's Claude 3 API.

**Email Integration**:
POST /email/get-email to retrieve a user's email address.
POST /email/send to send emails on behalf of the user using Gmail OAuth.

## Middleware
**Authentication Middleware**: Ensures that requests to protected routes are accompanied by a valid JWT (JSON Web Token).
**Rate Limiting Middleware**: Limits the rate of requests to 2 requests per second per client based on IP address.

## Error Handling and Logging
Any issues encountered during the handling of requests are logged into an SQLite database hosted on Turso. This includes a table for error logging and a table for storing user information.

## Tech Stack
- Gin framework for the performance and efficiency in building scalable web servers with Go.
- LLM APIs to understand the user request (OpenAI, Anthropic, Mistral, Groq APIs). Using accurate prompt to transform user request into an action code to make the user request more deterministic.
- SQLite Database hosted on Turso to get a lightweight and reliable storage solution for user data and error logs. Using the db to essantially store errors to improve prompts in the future.
- Gmail OAuth to connect users to their gmail account to send emails on their behalf.

## Workflow with Client Application
**For instance, a user wishes to open the Discord application via the desktop app**:
- User types "open discord" in the app.
- The app sends this command to /chat with the user's JWT and the message in JSON format: {"Message": "open discord"}.
- The server validates the JWT and rate limits, then processes the message through sequential LLM API calls to:
  - Determine the user's intent to interact with software.
  - Identify the intent to launch an application.
  - Confirm the specific application to be opened.
- The server returns the action code 001discord.
- Upon receiving 001discord, the desktop application executes the command to launch Discord.

The strategy of making sequential LLM calls for each phase of request interpretation enhances accuracy and reliability. By narrowing down the scope incrementally, the server can generate precise action codes, significantly reducing the likelihood of errors compared to making a single, broad LLM query.
This multi-step processing ensures that user commands are executed accurately and efficiently, enhancing user experience and system responsiveness.
