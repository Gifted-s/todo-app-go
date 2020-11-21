package router

import (
	"github.com/gorilla/mux"
    "../middlewares"
)


func Router() *mux.Router{
 router := mux.NewRouter()
 router.HandleFunc("/", middlewares.HandleHome).Methods("GET", "OPTIONS")
 router.HandleFunc("/lists", middlewares.HandleCreateList).Methods("POST", "OPTIONS")
 router.HandleFunc("/lists", middlewares.HandleGetLists).Methods("GET", "OPTIONS")
 router.HandleFunc("/lists/{id}", middlewares.HandleEditList).Methods("PATCH", "OPTIONS")
 router.HandleFunc("/lists/{id}", middlewares.HandleDelete).Methods("DELETE", "OPTIONS")
 router.HandleFunc("/lists/{id}", middlewares.HandleGetList).Methods("GET", "OPTIONS")
 router.HandleFunc("/lists/add-task/{id}", middlewares.HandleAddTask).Methods("POST", "OPTIONS")
 router.HandleFunc("/lists/edit-task/{id}/{task_id}", middlewares.HandleEditTask).Methods("PATCH", "OPTIONS")
 router.HandleFunc("/lists/delete-task/{id}/{task_id}", middlewares.HandleDeleteTask).Methods("DELETE", "OPTIONS")
 router.HandleFunc("/lists/complete-task/{id}/{task_id}", middlewares.HandleCompleteTask).Methods("PATCH", "OPTIONS")
 router.HandleFunc("/lists/undo-task/{id}/{task_id}", middlewares.HandleUndoTask).Methods("PATCH", "OPTIONS")
 return router
}