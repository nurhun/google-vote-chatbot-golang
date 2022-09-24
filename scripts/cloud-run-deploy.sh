#!/bin/bash

# Enable required services:
gcloud services enable \
    chat.googleapis.com \
    cloudbuild.googleapis.com \
    run.googleapis.com \
    artifactregistry.googleapis.com


# Deploy with service account
gcloud beta run deploy --service-account=xxxxx@<PROJECT_ID>.iam.gserviceaccount.com
