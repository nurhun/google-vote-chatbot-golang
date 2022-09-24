#!/bin/bash

# Enable required services:
gcloud services enable \
    chat.googleapis.com \
    cloudbuild.googleapis.com \
    appengineflex.googleapis.com \

# Deploy with service account
gcloud beta app deploy --service-account=xxxxx@<PROJECT_ID>.iam.gserviceaccount.com
