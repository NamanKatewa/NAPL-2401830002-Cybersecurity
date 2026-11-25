(ns student-management.core
  (:require [next.jdbc :as jdbc]
            [ring.adapter.jetty :as jetty])
  (:gen-class))

;; --- (3) Relational Database (H2 in-memory) with try/catch error handling ---
(def ds (jdbc/get-datasource {:dbtype "h2:mem" :dbname "simpledb"}))

(defn init-db []
  (try
    (jdbc/execute! ds ["CREATE TABLE students (id INT PRIMARY KEY, name VARCHAR(50), cgpa DOUBLE)"])
    (println "Database initialized.")
    (catch Exception e
      (println "Error initializing DB:" (.getMessage e)))))

(defn write-student [id name cgpa]
  (try
    (jdbc/execute! ds ["INSERT INTO students (id, name, cgpa) VALUES (?, ?, ?)" id name cgpa])
    (println (str "Inserted student: " name " (CGPA: " cgpa ")"))
    (catch Exception e
      (println (str "Error writing student " id ": " (.getMessage e))))))

(defn read-students []
  (try
    (mapv #(hash-map :id (:STUDENTS/ID %) :name (:STUDENTS/NAME %) :cgpa (:STUDENTS/CGPA %))
          (jdbc/execute! ds ["SELECT id, name, cgpa FROM students"]))
    (catch Exception e
      (println "Error reading from DB:" (.getMessage e))
      [])))

;; --- (1) Process dataset using sequence functions (filter, map, reduce) ---
(defn process-pipeline [students]
  (let [honors   (filter #(>= (:cgpa %) 8.5) students)                ;; filter
        labeled  (map #(assoc % :status "Distinction") honors)        ;; map
        avg-cgpa (if (seq labeled)                                    ;; reduce
                   (/ (reduce + (map :cgpa labeled)) (count labeled))
                   0.0)]
    {:distinction (vec labeled)
     :average-cgpa (double avg-cgpa)}))

;; --- (2) Ring HTTP Endpoint ---
(defn handler [request]
  (if (= (:uri request) "/results")
    {:status 200
     :headers {"Content-Type" "text/plain"}
     :body (str "Processed Results:\n" (pr-str (process-pipeline (read-students))))}
    {:status 200
     :headers {"Content-Type" "text/plain"}
     :body "Clojure Data Pipeline API. Go to /results to see processed data."}))

;; --- (4) Java Interop & Main Entry Point ---
(defn -main [& args]
  ;; (4) Java library interop
  (println "Timestamp (Java Interop):" (str (java.time.Instant/now)))

  ;; (3) Database Init & Writes with error handling
  (init-db)
  (write-student 1 "Alice" 9.2)
  (write-student 2 "Bob" 7.5)
  (write-student 3 "Charlie" 8.8)
  ;; Trigger try/catch on duplicate primary key
  (write-student 1 "Duplicate Alice" 9.5)

  ;; (1) Read from DB and process through pipeline
  (let [data (read-students)
        results (process-pipeline data)]
    (println "\nRaw DB Records:" data)
    (println "Processed Pipeline Results:" results))

  ;; (2) Ring HTTP Server
  (println "\nStarting Ring server on port 3000...")
  (let [server (jetty/run-jetty handler {:port 3000 :join? false})]
    ;; Test the endpoint directly
    (println "Testing /results endpoint:")
    (println (:body (handler {:uri "/results"})))
    (.stop server)
    (println "Done.")))
