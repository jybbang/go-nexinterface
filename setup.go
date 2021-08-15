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
		AddHandler(new(commands.CreateBookCommand), commands.CreateBookCommandHandler).
		AddNotificationHandler(new(events.BookCreatedEvent), events.BookCreatedEventHandler).
		Build()

	mediator.
		AddMiddleware(middlewares.NewLogMiddleware())

	Log.Info("application initialized")
}

func infrastructureSetup() {
	core.NewEventbusBuilder().
		MessaingAdapter(mocks.NewMockAdapter()).
		Build()

	core.NewStateServiceBuilder().
		StateAdapter(mocks.NewMockAdapter()).
		Build()

	core.NewRepositoryServiceBuilder(new(entities.Book)).
		QueryRepositoryAdapter(mocks.NewMockAdapter()).
		CommandRepositoryAdapter(mocks.NewMockAdapter()).
		Build()

	Log.Info("infrastructure initialized")
}
