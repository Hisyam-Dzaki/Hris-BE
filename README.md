# Daun HR Backend (hris-BE)

Backend REST API untuk sistem HRIS Daun HR. Dibangun dengan Go + Gin + GORM + PostgreSQL.

## Tech Stack

| Komponen | Teknologi |
|----------|-----------|
| Bahasa | Go 1.25 |
| Framework | Gin v1.12 |
| ORM | GORM v1.31 |
| Database | PostgreSQL 16 |
| Auth | JWT (HS256) |
| Password | bcrypt |
| Container | Docker + Docker Compose |

## Struktur Project

```
hris-BE/
├── cmd/server/main.go                 # Entry point
├── internal/
│   ├── config/config.go               # App configuration (env vars)
│   ├── database/database.go           # DB connection + auto-migrate
│   ├── dto/dto.go                     # Request/Response structs
│   ├── middleware/
│   │   ├── auth.go                    # JWT auth + role middleware
│   │   └── cors.go                    # CORS configuration
│   ├── models/                        # GORM models (6 tabel)
│   │   ├── company.go
│   │   ├── user.go
│   │   ├── role.go
│   │   ├── employee.go
│   │   ├── department.go
│   │   └── location.go
│   ├── repositories/                  # Data access layer
│   │   ├── user_repo.go
│   │   ├── company_repo.go
│   │   ├── role_repo.go
│   │   ├── employee_repo.go
│   │   ├── department_repo.go
│   │   └── location_repo.go
│   ├── services/                      # Business logic layer
│   │   ├── auth_service.go
│   │   ├── company_service.go
│   │   ├── role_service.go
│   │   ├── employee_service.go
│   │   ├── department_service.go
│   │   └── location_service.go
│   ├── handlers/                      # HTTP handlers
│   │   ├── auth_handler.go
│   │   ├── company_handler.go
│   │   ├── role_handler.go
│   │   ├── employee_handler.go
│   │   ├── department_handler.go
│   │   ├── location_handler.go
│   │   └── helpers.go
│   └── router/router.go              # Route registration
├── pkg/
│   ├── jwt/jwt.go                     # JWT generate/verify
│   ├── bcrypt/bcrypt.go               # Password hash/check
│   └── response/response.go           # Standard JSON response
├── docker-compose.yml                 # PostgreSQL container
├── Dockerfile                         # Multi-stage build
├── .env                               # Environment variables
├── .gitignore
└── go.mod
```

## Arsitektur Clean Architecture

```
HTTP Request
    ↓
[Handler] ← menerima request, validasi input
    ↓
[Service]   ← business logic, validasi rules
    ↓
[Repository] ← akses database via GORM
    ↓
[Model]     ← representasi tabel database
```

**Alur:**
1. Client kirim request ke handler
2. Handler parse + validasi input (DTO)
3. Handler panggil service sesuai business logic
4. Service panggil repository untuk akses DB
5. Repository query database via GORM
6. Response dikembalikan ke client dalam format JSON

## Multi-Tenant

Sistem menggunakan **Shared Database + company_id**:

- Setiap tabel memiliki kolom `company_id`
- User hanya bisa akses data perusahaan sendiri
- JWT token menyimpan `company_id` pengguna

## Environment Variables

| Variable | Default | Keterangan |
|----------|---------|------------|
| `APP_PORT` | `8080` | Port server |
| `APP_ENV` | `development` | Environment |
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `hris_admin` | Database user |
| `DB_PASSWORD` | `hris_secret_2026` | Database password |
| `DB_NAME` | `hris_db` | Database name |
| `DB_SSLMODE` | `disable` | SSL mode |
| `JWT_SECRET` | - | Secret key JWT |
| `JWT_EXPIRY_HOURS` | `24` | Token expiry (jam) |

## Cara Jalankan

### 1. Start PostgreSQL

```bash
docker-compose up -d
```

### 2. Jalankan Server

```bash
go run ./cmd/server
```

Server berjalan di `http://localhost:8080`

### 3. Build Binary

```bash
go build -o hris-server ./cmd/server
./hris-server
```

### 4. Docker Build

```bash
docker build -t hris-be .
docker run -p 8080:8080 --env-file .env hris-be
```

## API Endpoints

### Health Check

```
GET /health
Response: { "status": "ok", "service": "hris-BE" }
```

### Authentication

| Method | Endpoint | Keterangan | Auth |
|--------|----------|------------|------|
| `POST` | `/api/auth/register` | Daftar akun baru | - |
| `POST` | `/api/auth/login` | Login | - |
| `GET` | `/api/auth/me` | Data user dari token | Ya |

### Company (Multi-Tenant)

| Method | Endpoint | Keterangan | Auth |
|--------|----------|------------|------|
| `POST` | `/api/companies` | Buat perusahaan | Ya |
| `GET` | `/api/companies` | List semua perusahaan | Ya |
| `GET` | `/api/companies/:id` | Detail perusahaan | Ya |
| `PUT` | `/api/companies/:id` | Update perusahaan | Ya |
| `DELETE` | `/api/companies/:id` | Hapus perusahaan | Ya |

### Roles

| Method | Endpoint | Keterangan | Auth |
|--------|----------|------------|------|
| `POST` | `/api/roles` | Buat role | Ya |
| `GET` | `/api/roles` | List role perusahaan | Ya |
| `GET` | `/api/roles/:id` | Detail role | Ya |
| `PUT` | `/api/roles/:id` | Update role | Ya |
| `DELETE` | `/api/roles/:id` | Hapus role | Ya |

### Employees

| Method | Endpoint | Keterangan | Auth |
|--------|----------|------------|------|
| `POST` | `/api/employees` | Tambah karyawan | Ya |
| `GET` | `/api/employees` | List karyawan (paginated) | Ya |
| `GET` | `/api/employees/:id` | Detail karyawan | Ya |
| `PUT` | `/api/employees/:id` | Update karyawan | Ya |
| `DELETE` | `/api/employees/:id` | Hapus karyawan | Ya |

**Query Parameters untuk GET /api/employees:**
- `page` (int) - Halaman, default 1
- `limit` (int) - Item per halaman, default 20, max 100
- `search` (string) - Search by name, NIK, email
- `department_id` (int) - Filter by departemen
- `location_id` (int) - Filter by lokasi
- `status` (string) - Filter by status (active/inactive/terminated)

### Departments

| Method | Endpoint | Keterangan | Auth |
|--------|----------|------------|------|
| `POST` | `/api/departments` | Buat departemen | Ya |
| `GET` | `/api/departments` | List departemen | Ya |
| `GET` | `/api/departments/:id` | Detail departemen | Ya |
| `PUT` | `/api/departments/:id` | Update departemen | Ya |
| `DELETE` | `/api/departments/:id` | Hapus departemen | Ya |

### Locations

| Method | Endpoint | Keterangan | Auth |
|--------|----------|------------|------|
| `POST` | `/api/locations` | Buat lokasi | Ya |
| `GET` | `/api/locations` | List lokasi | Ya |
| `GET` | `/api/locations/:id` | Detail lokasi | Ya |
| `PUT` | `/api/locations/:id` | Update lokasi | Ya |
| `DELETE` | `/api/locations/:id` | Hapus lokasi | Ya |

## Format Response

### Success

```json
{
  "status": "success",
  "message": "Resource created",
  "data": { ... }
}
```

### Success with Pagination

```json
{
  "status": "success",
  "data": [ ... ],
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 150,
    "total_pages": 8
  }
}
```

### Error

```json
{
  "status": "error",
  "message": "Invalid email or password"
}
```

## Request/Response Contoh

### Register

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@company.com",
    "password": "secret123",
    "name": "Admin",
    "company_id": 1
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@company.com",
    "password": "secret123"
  }'
```

### Create Employee (dengan JWT)

```bash
curl -X POST http://localhost:8080/api/employees \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "nik": "EMP001",
    "first_name": "Budi",
    "last_name": "Santoso",
    "email": "budi@company.com",
    "phone": "08123456789",
    "department_id": 1,
    "location_id": 1,
    "position": "Software Engineer",
    "employment": "PKWTT",
    "join_date": "2026-01-15"
  }'
```

### Get Employees with Filter

```bash
curl "http://localhost:8080/api/employees?page=1&limit=10&search=budi&department_id=1" \
  -H "Authorization: Bearer <token>"
```

## Database Schema

### companies
| Kolom | Tipe | Keterangan |
|-------|------|------------|
| id | uint | Primary key |
| name | string | Nama perusahaan |
| slug | string | Slug (unique) |
| address | text | Alamat |
| phone | string | Telepon |
| email | string | Email |
| website | string | Website |
| logo | string | Logo URL |

### users
| Kolom | Tipe | Keterangan |
|-------|------|------------|
| id | uint | Primary key |
| company_id | uint | Foreign key → companies |
| email | string | Email (unique) |
| password | string | Hashed password |
| name | string | Nama lengkap |
| role_id | uint | Foreign key → roles |
| role_name | string | Nama role |
| is_active | bool | Status aktif |

### roles
| Kolom | Tipe | Keterangan |
|-------|------|------------|
| id | uint | Primary key |
| company_id | uint | Foreign key → companies |
| name | string | Nama role |
| description | string | Deskripsi |
| permissions | string | JSON permissions |

### employees
| Kolom | Tipe | Keterangan |
|-------|------|------------|
| id | uint | Primary key |
| company_id | uint | Foreign key → companies |
| user_id | uint | Foreign key → users (nullable) |
| nik | string | Nomor induk karyawan |
| first_name | string | Nama depan |
| last_name | string | Nama belakang |
| email | string | Email |
| phone | string | Telepon |
| department_id | uint | Foreign key → departments |
| location_id | uint | Foreign key → locations |
| position | string | Jabatan |
| employment | string | PKWT/PKWTT/Magang |
| status | string | active/inactive/terminated |
| join_date | date | Tanggal masuk |

### departments
| Kolom | Tipe | Keterangan |
|-------|------|------------|
| id | uint | Primary key |
| company_id | uint | Foreign key → companies |
| name | string | Nama departemen |
| code | string | Kode departemen |

### locations
| Kolom | Tipe | Keterangan |
|-------|------|------------|
| id | uint | Primary key |
| company_id | uint | Foreign key → companies |
| name | string | Nama lokasi |
| address | text | Alamat |
| city | string | Kota |
| timezone | string | Timezone |

## Phased Development

### Phase 1 (Completed)
- Project setup + Docker
- Auth system (register, login, JWT)
- User & Role management
- Company (multi-tenant)
- Employee management (CRUD + search + pagination)
- Department & Location

### Phase 2 (Next)
- Attendance (clock-in/out)
- Leave management
- Shift scheduling
- Contract management
- Onboarding/Offboarding

### Phase 3
- Payroll processing
- Salary components & grades
- Compensation & benefits
- Reimbursement
- Compliance (PPh21, BPJS)

### Phase 4
- Performance reviews
- Recruitment pipeline
- Learning/LMS
- Certifications

### Phase 5
- Asset management
- Helpdesk tickets
- Surveys & engagement
- Announcements
- Audit logs

### Phase 6
- Settings (all configurations)
- Roles & permissions matrix
- Integrations & API keys
- Reports & analytics
- Multi-company entity switching

## License

Private - Daun HR
