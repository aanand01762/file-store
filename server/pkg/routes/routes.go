package routes

import (
	"github.com/gorilla/mux"

	"github.com/aanand01762/file-store/server/pkg/controllers"
)

var RegisterFileStoreRoutes = func(router *mux.Router) {

	//Define the routes for api calls
	router.HandleFunc("/store/add", controllers.AddFile).Methods("POST")
	router.HandleFunc("/store/update", controllers.UpdateFile).Methods("PUT")
	router.HandleFunc("/store/delete", controllers.DeleteFile).Methods("DELETE")
	router.HandleFunc("/store/files", controllers.GetFiles)
	router.HandleFunc("/store/word-count", controllers.GetWordCounts)

	/*


		//should update contents of file.txt in
		//server with the local file.txt or create a new file.txt in server if it is
		//absent.
		router.HandleFunc("/store/word-count", controllers.GetWordCounts)

		//store freq-words [--limit|-n 10] [--order=dsc|asc]
		// store freq-words should return the 10 most frequent words in
		// all the files combined. This should work the same as running the following
		// shell command:
		router.HandleFunc("/store/frequency", controllers.GetFrequency).Methods("POST")
	*/

}
