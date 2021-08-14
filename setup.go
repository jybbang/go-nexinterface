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
	mediator := core.GetMediator()

	mediator.AddHandler(new(commands.CreateBookCommand), commands.CreateBookCommandHandler)

	mediator.AddNotificationHandler(new(events.BookCreatedEvent), events.BookCreatedEventHandler)

	mediator.AddMiddleware(middlewares.NewLogMiddleware())

	Log.Info("application initialized")
}

func infrastructureSetup() {
	mock := mocks.NewMockAdapter()

	core.GetEventBus().SetMessaingAdapter(mock)
	core.GetStateService().SetStateAdapter(mock)

	bookRepository := core.GetRepositoryService(new(entities.Book))
	bookRepository.SetQueryRepositoryAdapter(mock)
	bookRepository.SetCommandRepositoryAdapter(mock)

	Log.Info("infrastructure initialized")
}
