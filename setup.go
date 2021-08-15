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
	core.GetEventBus().SetMessaingAdapter(mocks.NewMockAdapter())
	core.GetStateService().SetStateAdapter(mocks.NewMockAdapter())

	bookRepository := core.GetRepositoryService(new(entities.Book))
	bookRepository.SetQueryRepositoryAdapter(mocks.NewMockAdapter())
	bookRepository.SetCommandRepositoryAdapter(mocks.NewMockAdapter())

	Log.Info("infrastructure initialized")
}
