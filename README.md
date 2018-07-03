## Build the app
```
go build main.go
```
## Run the app
```
./main
```

or 

```
go run main.go
```
https://github.com/GoogleCloudPlatform/container-builder-workshop

```
 gcloud container clusters create 
```

```
export PROJECT=
export CLUSTER=
export ZONE=
gcloud config set compute/zone $ZONE
```

```
 gcloud container clusters get-credentials ${CLUSTER} --project=${PROJECT} --zone=${ZONE}
```
```
 PROJECT_NUMBER="$(gcloud projects describe $(gcloud config get-value core/project -q) --format='get(projectNumber)')"
```
```
gcloud projects add-iam-policy-binding ${PROJECT} \
--member=serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com \
--role=roles/container.developer
```

```
gcloud container clusters delete
```