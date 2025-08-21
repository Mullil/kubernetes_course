#!/bin/bash
set -e

DATE=$(date +%Y-%m-%d_%H-%M-%S)
FILE="backup-${DATE}.sql"

PGPASSWORD=$DB_PASSWORD pg_dump -h $DB_HOST -U $DB_USER $DB_NAME > $FILE
gsutil cp $FILE gs://$GCS_BUCKET/backups/
