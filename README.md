# StanCodingChallenge

Implementate a simple JSON-based web service written in GO and deployed using heroku. [Coding Challenge document.](https://challengeaccepted.streamco.com.au/)

# Run application
These instructions assume you have GO installed.

```
$> go run main.go
```

OR

```
$> go build
$> ./StanCodingChallenge
```

# Run tests

```
$> cd test
$> go test
```

# Heroku endpoint
[App deployed using Heroku](https://stan-coding-challenge-1.herokuapp.com/)

## Heroku
Heroku is a platform as a service (PaaS) that enables developers to build, run, and operate applications entirely in the cloud.

###How to deploy to Heroku:

```
$> heroku create stan-coding-challenge-1
$> git add .
$> git commit -m "message"
$> git push
$> git push heroku main
```

###How to change a github remote on Heroku?

```
$> git remote rm heroku
$> git remote add heroku https://git.heroku.com/stan-coding-challenge-1.git
$> git push heroku main
```
