# PRD — akubisa

## Ringkasan Produk
akubisa adalah aplikasi pembelajaran membaca, menulis, dan berhitung (calistung) berbasis AI dan gamification untuk anak usia 4–8 tahun.

---

# Tujuan Produk

## Tujuan Utama
- Membantu anak belajar calistung secara menyenangkan
- Mendukung adaptive learning
- Mempermudah monitoring orang tua dan guru

## Tujuan Bisnis
- Subscription-based edtech platform
- Early childhood AI learning ecosystem

---

# Target Pengguna

## Primary Users
- Anak usia 4–8 tahun

## Secondary Users
- Orang tua
- Guru
- Tutor

---

# Fitur Utama

## Modul Membaca
- Belajar alfabet
- Fonik
- Suku kata
- Voice recognition
- Pronunciation scoring

## Modul Menulis
- Tracing huruf
- Tracing angka
- Penilaian tulisan

## Modul Berhitung
- Penjumlahan
- Pengurangan
- Counting objects
- Quiz interaktif

## Gamification
- XP
- Badge
- Reward
- Daily streak

## Adaptive Learning
- Personalized lesson
- Difficulty adjustment
- Weakness analysis

## Dashboard Orang Tua
- Progress monitoring
- Weekly report
- Learning insight

---

# Teknologi

| Layer | Technology |
|---|---|
| Mobile | Flutter |
| Backend | Golang Fiber |
| Database | PostgreSQL |
| Cache | Redis |
| Queue | RabbitMQ |
| AI | Whisper + OpenAI |
| Storage | S3 / MinIO |
| Infra | Kubernetes |

---

# Database Design

## Main Tables
- users
- children
- lessons
- quizzes
- progress
- achievements
- rewards
- analytics_events
- ai_pronunciation_results

---

# System Design

## Architecture
- Clean Architecture
- Modular Monolith
- Event-Driven Architecture
- Future-ready Microservices

## Services
- Authentication Service
- Learning Service
- Gamification Service
- AI Service
- Analytics Service
- Notification Service

---

# Backend Implementation

## Backend Stack
- Golang
- Fiber
- PostgreSQL
- Redis
- JWT Authentication

---

# API Development

## Authentication APIs
- POST /api/v1/auth/register
- POST /api/v1/auth/login

## Lesson APIs
- GET /api/v1/lessons
- GET /api/v1/lessons/:id

## Quiz APIs
- POST /api/v1/quizzes/:id/submit

## AI APIs
- POST /api/v1/ai/pronunciation/analyze

---

# Security

- JWT Authentication
- HTTPS only
- Rate limiting
- SQL injection prevention
- Child-safe architecture

---

# Deployment

## Infrastructure
- Docker
- Kubernetes
- CI/CD GitHub Actions

## Monitoring
- Grafana
- Prometheus
- Loki

---

# Roadmap

## Phase 1
- MVP
- Authentication
- Basic lessons

## Phase 2
- AI pronunciation
- Adaptive learning

## Phase 3
- Analytics
- Subscription
- Teacher dashboard

---

# Kesimpulan

akubisa dirancang sebagai platform edtech modern yang menggabungkan:
- AI
- adaptive learning
- gamification
- scalable architecture

untuk menciptakan pengalaman belajar calistung yang menyenangkan dan efektif.
