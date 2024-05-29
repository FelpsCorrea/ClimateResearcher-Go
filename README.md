# ClimateResearcher-Go

This repository is a web API developed in Golang as a partial assessment for the completion of the Postgraduate Degree in Golang.

In this project, concepts such as the following were utilized:
- Clean Architecture: To decouple components as much as possible.
- Docker: Necessary for setting up the environment for deployment on Google Cloud's Cloud Run.
- Environment Variables with Viper: For configuration management.
- Google Cloud: The environment used to host the [application](https://climate-researcher-go-6oveq7mywa-uc.a.run.app/weather/83325625)
- External API Calls: Utilized the [ViaCep](https://viacep.com.br) API and [Weather](https://www.weatherapi.com) API
- Fast Json (Go 1.22)

Individual tests for the external APIs are present in `internal/infra/external`.

Additionally, if you wish to run this project locally, configure the environment variables present in `cmd/weathersystem`. Create a `.env` file following the model in the `.env.example` file.

To test the local http request, use the `weather.htt`p file present in `internal/infra/web/test`, which contains tests for scenarios like: invalid Zip code, valid Zip code, and Zip code not found.

Some other concepts that could have been covered in the project include:
- Panic Recover: To allow the program to recover after a failure.
- Dependency Injection using Google Wire: So that dependency injection doesn't need to be handled in the `main.go` file.

