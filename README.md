# Birthday Notify Service

[![Go](https://img.shields.io/badge/go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

To get started, clone this repository and follow these steps to run the Go application in your environment.

### Environment Variables

Create a `.env` file in the root directory with the following environment variables:

```bash
# API Configuration
BIRTHDAYS_API_URL=https://your-api-url.com/api/birthdays

# Email Configuration
SENDGRID_API_KEY=your-sendgrid-api-key
SENDGRID_FROM_NAME=Your Name
SENDGRID_FROM_EMAIL=your-email@example.com
```

**Required Environment Variables:**

- `BIRTHDAYS_API_URL` - The API endpoint URL to fetch birthday data from
- `SENDGRID_API_KEY` - Your SendGrid API key for sending emails
- `SENDGRID_FROM_NAME` - The display name for the email sender
- `SENDGRID_FROM_EMAIL` - The email address to send emails from (must be verified in SendGrid)

**Note:** The `.env` file is excluded from version control (see `.gitignore`). Make sure to create your own `.env` file with the appropriate values for your environment.

### Running the Application

Start the application:

```bash
make run
```

Build the application:

```bash
make build
```
