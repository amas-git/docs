package main

type App interface {
	OnStart() bool
	OnStop() bool
	HealthCheck() bool
}
