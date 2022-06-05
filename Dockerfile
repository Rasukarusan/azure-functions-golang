FROM golang:1.18-alpine
COPY . /go/app
WORKDIR /go/app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM mcr.microsoft.com/azure-functions/dotnet:3.0-appservice
ENV AzureWebJobsScriptRoot=/home/site/wwwroot \
    AzureFunctionsJobHost__Logging__Console__IsEnabled=true
ENV TZ=Asia/Tokyo
COPY --from=0 /go/app/ /home/site/wwwroot