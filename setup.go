package main

import (
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/go-core-architecture/infrastructure/mocks"
	"github.com/jybbang/go-core-architecture/middlewares"
	"github.com/jybbang/nexinterface/src/books/commands"
	"github.com/jybbang/nexinterface/src/books/events"
	"github.com/jybbang/nexinterface/src/entities"
)

func applicationSetup() {
	mediator := core.NewMediatorBuilder().
		AddPerformanceMeasure(logger).
		AddHandler(new(commands.CreateBookCommand), commands.CreateBookCommandHandler).
		AddNotificationHandler(new(events.BookCreatedEvent), events.BookCreatedEventHandler).
		Build()

	mediator.
		AddMiddleware(middlewares.NewLogMiddleware(logger)).
		AddMiddleware(middlewares.NewValidationMiddleware())

	log.Info("application initialized")
}

func infrastructureSetup() {
	mock := mocks.NewMockAdapter()

	core.NewEventbusBuilder().
		MessaingAdapter(mock).
		Build()

	core.NewStateServiceBuilder().
		StateAdapter(mock).
		Build()

	core.NewRepositoryServiceBuilder(new(entities.Book), "T_BOOK").
		QueryRepositoryAdapter(mock).
		CommandRepositoryAdapter(mock).
		Build()

	log.Info("infrastructure initialized")
}
