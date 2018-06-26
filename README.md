# todo-api

My implementation of the todo list api as found at https://app.swaggerhub.com/apis/aweiker/ToDo/1.0.0

A couple of notes, I used 2 third party libraries for this implementation. The first is the gorilla mux router for parsing paths and routing them to their respective methods. I chose to use this 3rd party library as it is a light weight router that has a high level of adoption within many of the golang open source repositories on github.com. I chose to use this library rather than implement my own because I felt that this assignment was to showcase how I work on a day to day basis, which for me includes making use of relevant open source libraries where applicable. Under the same logic I used the Google uuid generator library to provide the ids for lists and tasks. If desired I can go back and implement my own version of these 2 libraries.

I implemented the in memory list storage as a lightweight example of how these simple todo lists could be implemented. However with the list interface I created it would also be possible to swap out this implementation for other storage methods, such as a database or file storage for persistence of the lists. 