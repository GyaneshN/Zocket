# Zocket

AI-Powered Task Management System
Overview
This is a full-stack AI-powered task management system built using Golang (Gin/Fiber) for the backend and Next.js (TypeScript + Tailwind CSS) for the frontend. The system enables real-time task tracking, JWT-based authentication, and AI-powered task recommendations. It is deployed on Render/Fly.io for the backend and Vercel for the frontend.



Note: The frontend depends on the backend API, so the backend is set up and running first. i am facing errors in the backend, they must be resolved before proceeding with frontend development
Features
User Authentication: Secure JWT-based authentication.

Task Management: Create, assign, update, and track tasks.

Real-Time Updates: WebSocket integration for instant task updates.

AI-Powered Suggestions: OpenAI/Gemini API integration for smart task breakdowns.

Cloud Deployment: Fully hosted on Render/Fly.io (backend) and Vercel (frontend).

High Performance: Golang's Gin/Fiber for speed and efficiency.

Scalability: PostgreSQL/MongoDB for data storage, Docker for containerization.

Project Structure
├── backend/ (Golang - Gin/Fiber)
│   ├── main.go
│   ├── routes/
│   ├── controllers/
│   ├── models/
│   ├── middleware/
│   ├── config/
│   ├── database/
│   └── .env
├── frontend/ (Next.js + Tailwind CSS)
│   ├── pages/
│   ├── components/
│   ├── styles/
│   ├── utils/
│   ├── public/
│   ├── .env.local
│   └── next.config.js
└── README.md
Setup & Installation
Backend (Golang)
Clone Repository:

git clone https://github.com/your-repo/ai-task-manager.git
cd ai-task-manager/backend
Install Dependencies:

go mod tidy
Setup Environment Variables (.env file):

PORT=4000
DB_URL=your_database_url
JWT_SECRET=your_jwt_secret
OPENAI_API_KEY=your_openai_key
Run Server:

go run main.go
Frontend (Next.js + TypeScript)
Note: The frontend depends on the backend API, so ensure the backend is set up and running first. If there are errors in the backend, they must be resolved before proceeding with frontend development.

Navigate to frontend:

cd ../frontend
Install Dependencies:

npm install
Setup Environment Variables (.env.local):

NEXT_PUBLIC_BACKEND_URL=http://localhost:4000
NEXT_PUBLIC_OPENAI_API_KEY=your_openai_key
Run Frontend:

npm run dev
Deployment
Backend Deployment (Render/Fly.io)
Create a Render/Fly.io account

Deploy Golang Backend:

git push render main
Frontend Deployment (Vercel)
Install Vercel CLI:

npm install -g vercel
Deploy Frontend:

vercel
AI Utilization
How AI Tools Helped


OpenAI API: Provided smart task suggestions based on user history.

AutoGPT: Used for generating efficient database schemas and API structures
