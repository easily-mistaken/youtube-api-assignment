# YouTube Video API Dashboard

A modern web application that fetches and displays YouTube videos with real-time updates, built with Go and React TypeScript.

## 🚀 Features

- **Real-time Video Updates**: Automatically syncs with YouTube data
- **Pagination**: Efficient browsing through video lists
- **Channel Filtering**: Filter videos by specific channels
- **Responsive Design**: Works seamlessly on all devices
- **REST API**: Clean and well-structured Go backend
- **Database Storage**: Persistent PostgreSQL storage

## 🛠️ Setup and Installation

- git clone https://github.com/easily-mistaken/youtube-api-assignment
- Set up environment variables:
- Configure your `.env` file using `.env.sample` file:
- Database Configuration
  - DB_HOST=localhost
  - DB_PORT=5432
  - DB_USER=postgres
  - DB_PASSWORD=your_password
  - DB_NAME=youtube_db
  - YouTube API Configuration
  - YOUTUBE_API_KEY=your_youtube_api_key
  - Server Configuration
  - PORT=8080

### Prerequisites

- Go 1.19 or higher
- Node.js 16+ and npm
- PostgreSQL
- Docker (optional)

## 📝 API Documentation

### GET /api/videos

Fetch videos with pagination and filtering

#### Query Parameters

- `page`: Page number (default: 1)
- `limit`: Items per page (default: 10)
- `channel`: Filter by channel name (optional)

#### Response Example

````json
{
    "limit": 2,
    "page": 1,
    "total": 4751,
    "videos": [
        {
            "id": 1,
            "title": "Jersey Controversy 😱 Pak Name Likna Padega Kya 😎 Team India's Playing 11 🔥 Skyball Vs Bazzball",
            "description": "Jersey Controversy Pak Name Likna Padega Kya Team India's Playing 11 Skyball Vs Bazzball ...",
            "publishedAt": "2025-01-22T06:07:57+05:30",
            "thumbnailUrl": "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
            "channel": "AB Cricinfo"
        },
        {
            "id": 2,
            "title": "Jersey Controversy 😱 Pak Name Likna Padega Kya 😎 Team India's Playing 11 🔥 Skyball Vs Bazzball",
            "description": "Jersey Controversy Pak Name Likna Padega Kya Team India's Playing 11 Skyball Vs Bazzball ...",
            "publishedAt": "2025-01-22T06:07:57+05:30",
            "thumbnailUrl": "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
            "channel": "AB Cricinfo"
        }
    ]
}```

## 💻 Tech Stack

### Backend

- **Go**: Primary programming language
- **Gin**: Web framework for handling HTTP requests
- **GORM**: ORM for database operations
- **PostgreSQL**: Database for storing video data
- **Docker**: Containerization

### Frontend

- **React**: UI library
- **TypeScript**: Type-safe JavaScript
- **Axios**: HTTP client
- **Tailwind CSS**: Utility-first CSS framework

## 🔧 Development

### Backend Development

The backend is structured using a clean architecture:

- `handler.go`: Contains API endpoint logic
- `routes.go`: Defines API routes and middleware
- `database.go`: Manages database connections
- `video.go`: Defines the video data model

### Frontend Development

The frontend follows a component-based architecture:

- `axios.ts`: Configures API requests
- `video.tsx`: Main video listing page
- `App.tsx`: Application routing and layout
- `main.tsx`: Application initialization

## 🚀 Deployment

### Using Docker

## 📄 License

This project is licensed under the MIT License.

## 👥 Authors

- Prajjwal Choubey

## 🙏 Acknowledgments

- YouTube Data API
- Go community
- React community
````
