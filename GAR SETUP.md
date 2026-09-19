<div align="center">

# Google Artifacts Registry Console Setup


</div>

## Prerequisites
- Create your GCP Project
- Take Note of your `Project Number` and `Project ID`

## Install gcloud CLI
- Install gcloud cli here: `https://docs.cloud.google.com/sdk/docs/install-sdk#deb`

## gcloud setup
- Make sure that you have login or just execute the command `gcloud init --console-only`
- Copy & Paste the generated link to authenticate via browser
- Copy & Paste the passcode from the browser and paste in Console
- Authenticate by executing the command: `gcloud auth login --no-browser`

## Enable Required Google APIs
```
gcloud services enable \
  iamcredentials.googleapis.com \
  iam.googleapis.com \
  sts.googleapis.com \
  artifactregistry.googleapis.com
```

## Create a Workload Identity Protocol
- execute below (github-actions can be replace with whater idp name you want):
```bash
gcloud iam workload-identity-pools create github-actions \
  --location="global" \
  --display-name="GitHub Actions"
```
- verify if you have successfully created WIP:
```bash
gcloud iam workload-identity-pools describe github-actions \
  --location="global"
```
- you should see something like this:
```
displayName: GitHub Actions
name: projects/PROJECT_NUMBER/locations/global/workloadIdentityPools/github-actions
state: ACTIVE
```

## Create the GitHub OIDC provider
- get your github repository (omit the https://github.com):
```
my-github-user/my-project
```
- Create the Provider (replace my-github-user/my-project):
```bash
gcloud iam workload-identity-pools providers create-oidc github \
  --location="global" \
  --workload-identity-pool="github-actions" \
  --issuer-uri="https://token.actions.githubusercontent.com/" \
  --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner" \
  --attribute-condition="assertion.repository == 'my-github-user/my-go-project'"
```
## Create a Google service account
- Create a dedicated service account for GitHub Actions:
```bash
gcloud iam service-accounts create github-actions \
  --display-name="GitHub Actions"
```
- You'll get something like this:
```
github-actions@my-project.iam.gserviceaccount.com
```

## Give the service account access to Artifact Registry
- replace the service account below and my-project
```bash
gcloud projects add-iam-policy-binding my-project \
  --member="serviceAccount:github-actions@my-project.iam.gserviceaccount.com" \
  --role="roles/artifactregistry.writer"
```

## Allow GitHub's identity to impersonate the service account
- replace the service account below, my-project, my-github-user/my-project and PROJECT_NUMBER
```bash
gcloud iam service-accounts add-iam-policy-binding \
  github-actions@my-project.iam.gserviceaccount.com \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/projects/PROJECT_NUMBER/locations/global/workloadIdentityPools/github-actions/attribute.repository/my-github-user/my-project"
```

## Get your Workload Identity Provider name
```bash
gcloud iam workload-identity-pools providers describe github \
  --location="global" \
  --workload-identity-pool="github-actions"
```




