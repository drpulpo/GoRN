# GoRN Docker Composed Fullstack Environment

1 Docker Container for Postgres
1 Docker Container for pgAdmin (Web)
1 Docker Container for Go written Backend API
1 Docker Container for Typescript written Frontend React Native Web Expo Server

\*Extra: Development tools for Expo Go React Native for Android and iOS using QR code

# Go React Native Expo CRUD with Postgres

CRUD App made with Go + React Native Expo. Uses Postgres container and pgAdmin Web container

# docker compose commands

1. docker-compose up
2. docker-compose down
3. docker-compose up --build

# GoRN Frontend

React Native (Expo) based frontend

# Go React Native CRUD with Postgres

CRUD App made with Go + React Native (Expo). Uses Postgres container

# Commands

After clonning the Repo

cd gorn-frontend

npm install

npx expo

Download Expo Go application for Android or iOS and Scan QR code

Press W for Web Client

Web is waiting on http://localhost:8081

---

For building frontend

npm run build:web

Then restart the Docker Compose at the root

cd ..

docker-compose up --build

Postgres runs at :5432

PGAdmin runs at :5050
admin@domain.com
password

Backend runs at :8080

Frontend runs at :3000
