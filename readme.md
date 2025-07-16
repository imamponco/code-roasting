# Code Roasting

This project provides an AI-powered code review tool that integrates with GitLab. It analyzes the diff of a merge request and posts AI-generated feedback as a comment.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- [Docker](https://docs.docker.com/get-docker/) (for running with Docker)

## Local Development

### Configuration

1. **Create an environment file:**
Copy the example environment file to a new `.env` file:

```bash
cp .example.env .env
```

1. **Edit the `.env` file:**
Open the `.env` file and fill in the required values:

```env
# GitLab Configuration
GITLAB_BASE_URL="https://gitlab.com"
GITLAB_PROJECT_ID="<your_project_id>"
GITLAB_MR_ID="<your_mr_id>"
GITLAB_ACCESS_TOKEN="<your_gitlab_personal_access_token>"

# AI Configuration
AI_TYPE="1" # Or your preferred AI type
AI_API_KEY="<your_ai_api_key>"
AI_CUSTOM_PROMPT="" # Optional: Add a custom prompt

# Debug Mode
DEBUG=false
```

**Note:** The application can also be configured using command-line flags, will override the values in the `.env` file.

### Running with Go

To run the application directly using Go, execute the following command:

```bash
go run . --gitlab-base-url $GITLAB_BASE_URL --gitlab-project-id $GITLAB_PROJECT_ID --gitlab-mr-id $GITLAB_MR_ID --gitlab-access-token $GITLAB_ACCESS_TOKEN --ai-api-key $AI_API_KEY
```

You can also pass the other flags as needed:

```bash
go run . --ai-type <type> --ai-custom-prompt "<prompt>" --debug
```

### Running with Docker

1. **Build the Docker image:**

```bash
docker build -t code-roasting .
```

1. **Run the Docker container:**
    You can run the application in a container, passing the configuration as environment variables or command-line arguments.

**Using an environment file:**

```bash
docker run --rm --env-file .env code-roasting
```

**Using command-line arguments:**

```bash
docker run --rm code-roasting \
  --gitlab-base-url "https://gitlab.com" \
  --gitlab-project-id "<your_project_id>" \
  --gitlab-mr-id "<your_mr_id>" \
  --gitlab-access-token "<your_gitlab_personal_access_token>" \
  --ai-api-key "<your_ai_api_key>"
```

## Command-Line Flags

| Flag                      | Environment Variable      | Description                       |
| ---------------------     | ------------------------- | --------------------------------- |
| `--gitlab-base-url`       | `GITLAB_BASE_URL`         | GitLab base URL                   |
| `--gitlab-project-id`     | `GITLAB_PROJECT_ID`       | GitLab project ID                 |
| `--gitlab-mr-id`          | `GITLAB_MR_ID`            | GitLab merge request ID           |
| `--gitlab-access-token`   | `GITLAB_ACCESS_TOKEN`     | GitLab access token               |
| `--ai-type`               | `AI_TYPE`                 | AI type (e.g., "1")               |
| `--ai-api-key`            | `AI_API_KEY`              | AI API key                        |
| `--ai-custom-prompt`      | `AI_CUSTOM_PROMPT`        | Custom prompt for the AI          |
| `--debug`                 | `DEBUG`                   | Enable debug mode                 |
