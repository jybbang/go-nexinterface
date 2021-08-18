package main

import (
	"time"

	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/go-core-architecture/infrastructure/mocks"
	"github.com/jybbang/go-core-architecture/middlewares"
	"github.com/jybbang/nexinterface/src/books/commands"
	"github.com/jybbang/nexinterface/src/books/events"
	"github.com/jybbang/nexinterface/src/books/queries"
	"github.com/jybbang/nexinterface/src/entities"
)

func applicationSetup() {
	mediatorBuilder := core.NewMediatorBuilder()

	// books
	mediatorBuilder.
		AddHandler(new(commands.CreateBookCommand), commands.CreateBookCommandHandler).
		AddHandler(new(commands.DeleteBookCommand), commands.DeleteBookCommandHandler).
		AddHandler(new(queries.GetBooksQuery), queries.GetBooksQueryHandler).
		AddHandler(new(queries.GetBookQuery), queries.GetBookQueryHandler).
		AddNotificationHandler(new(events.BookCreatedEvent), events.BookCreatedEventHandler).
		AddNotificationHandler(new(events.BookDeletedEvent), events.BookDeletedEventHandler)

	mediatorBuilder.
		Build().
		AddMiddleware(middlewares.NewLogMiddleware(logger)).
		AddNext(middlewares.NewPerformanceMiddleware(logger, time.Duration(500*time.Millisecond))).
		AddNext(middlewares.NewValidationMiddleware()).
		AddNext(middlewares.NewPublishDomainEventsMiddleware())

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
