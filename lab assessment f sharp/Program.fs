open System
open System.IO
open System.Text.Json

type Department =
    | ComputerScience
    | Cybersecurity
    | DataScience

type Student = {
    Id: int
    Name: string
    Department: Department
    Gpa: float
}

type StudentDoc = {
    DocId: int
    DocName: string
    DocDept: string
    DocGpa: float
}

module Database =
    let private dbFile = "students_db.json"
    let private options = JsonSerializerOptions(WriteIndented = true)

    let private toDoc (s: Student) : StudentDoc =
        let dept =
            match s.Department with
            | ComputerScience -> "ComputerScience"
            | Cybersecurity -> "Cybersecurity"
            | DataScience -> "DataScience"
        { DocId = s.Id; DocName = s.Name; DocDept = dept; DocGpa = s.Gpa }

    let private fromDoc (d: StudentDoc) : Student =
        let dept =
            match d.DocDept with
            | "ComputerScience" -> ComputerScience
            | "Cybersecurity" -> Cybersecurity
            | _ -> DataScience
        { Id = d.DocId; Name = d.DocName; Department = dept; Gpa = d.DocGpa }

    let private readDocsAsync () : Async<StudentDoc list> =
        async {
            if not (File.Exists(dbFile)) then
                return []
            else
                let! json = File.ReadAllTextAsync(dbFile) |> Async.AwaitTask
                if String.IsNullOrWhiteSpace(json) then
                    return []
                else
                    return JsonSerializer.Deserialize<StudentDoc list>(json)
        }

    let private writeDocsAsync (docs: StudentDoc list) : Async<unit> =
        async {
            let json = JsonSerializer.Serialize(docs, options)
            do! File.WriteAllTextAsync(dbFile, json) |> Async.AwaitTask
        }

    let initAsync () : Async<unit> =
        async {
            if File.Exists(dbFile) then
                File.Delete(dbFile)
            do! writeDocsAsync []
        }

    let createAsync (student: Student) : Async<Result<Student, string>> =
        async {
            let! docs = readDocsAsync ()
            if docs |> List.exists (fun d -> d.DocId = student.Id) then
                return Error (sprintf "Student with ID %d already exists." student.Id)
            else
                let updated = docs @ [ toDoc student ]
                do! writeDocsAsync updated
                return Ok student
        }

    let getAllAsync () : Async<Student list> =
        async {
            let! docs = readDocsAsync ()
            return docs |> List.map fromDoc
        }

    let getByIdAsync (id: int) : Async<Student option> =
        async {
            let! students = getAllAsync ()
            return students |> List.tryFind (fun s -> s.Id = id)
        }

    let updateAsync (student: Student) : Async<Result<Student, string>> =
        async {
            let! docs = readDocsAsync ()
            if not (docs |> List.exists (fun d -> d.DocId = student.Id)) then
                return Error (sprintf "Student with ID %d not found." student.Id)
            else
                let updated = docs |> List.map (fun d -> if d.DocId = student.Id then toDoc student else d)
                do! writeDocsAsync updated
                return Ok student
        }

    let deleteAsync (id: int) : Async<Result<int, string>> =
        async {
            let! docs = readDocsAsync ()
            if not (docs |> List.exists (fun d -> d.DocId = id)) then
                return Error (sprintf "Student with ID %d not found." id)
            else
                let remaining = docs |> List.filter (fun d -> d.DocId <> id)
                do! writeDocsAsync remaining
                return Ok id
        }

module Operations =
    let evaluateStanding (s: Student) : string =
        match s.Gpa with
        | gpa when gpa >= 3.8 -> "Summa Cum Laude"
        | gpa when gpa >= 3.5 -> "Dean's List"
        | gpa when gpa >= 2.5 -> "Good Standing"
        | _ -> "Academic Warning"

    let format (s: Student) : string =
        let dept =
            match s.Department with
            | ComputerScience -> "Computer Science"
            | Cybersecurity -> "Cybersecurity"
            | DataScience -> "Data Science"

        let standing = evaluateStanding s
        sprintf "ID: %-2d | %-14s | %-18s | GPA: %0.2f | %s" s.Id s.Name dept s.Gpa standing

    let filterByDept (dept: Department) (students: Student list) : Student list =
        students |> List.filter (fun s -> s.Department = dept)

    let filterHonors (students: Student list) : Student list =
        students |> List.filter (fun s -> s.Gpa >= 3.5)

    let report (students: Student list) : string list =
        students |> List.map format

    let averageGpa (students: Student list) : float =
        if List.isEmpty students then 0.0
        else students |> List.averageBy (fun s -> s.Gpa)

[<EntryPoint>]
let main argv =
    async {
        printfn "=== 1. Initializing Database ==="
        do! Database.initAsync ()

        printfn "\n=== 2. Creating Students (Create) ==="
        let students = [
            { Id = 1; Name = "Alice";   Department = Cybersecurity;   Gpa = 3.9 }
            { Id = 2; Name = "Bob";     Department = ComputerScience; Gpa = 3.2 }
            { Id = 3; Name = "Charlie"; Department = Cybersecurity;   Gpa = 2.4 }
            { Id = 4; Name = "Diana";   Department = DataScience;     Gpa = 3.8 }
        ]

        for s in students do
            let! res = Database.createAsync s
            match res with
            | Ok created -> printfn "Added: %s (ID: %d)" created.Name created.Id
            | Error err -> printfn "Error: %s" err

        printfn "\n=== 3. Current Records (Read All) ==="
        let! current = Database.getAllAsync ()
        Operations.report current |> List.iter (printfn "%s")

        printfn "\n=== 4. Functional Queries (Filter & Map) ==="
        printfn "\n-- Cybersecurity Students:"
        current |> Operations.filterByDept Cybersecurity |> Operations.report |> List.iter (printfn "  %s")

        printfn "\n-- Honors Students (GPA >= 3.5):"
        current |> Operations.filterHonors |> Operations.report |> List.iter (printfn "  %s")

        printfn "\n-- Average GPA: %0.2f" (Operations.averageGpa current)

        printfn "\n=== 5. Updating Student 2 (Update) ==="
        let updatedBob = { Id = 2; Name = "Bob"; Department = ComputerScience; Gpa = 3.7 }
        let! updateRes = Database.updateAsync updatedBob
        match updateRes with
        | Ok u -> printfn "Updated successfully: %s new GPA is %0.2f" u.Name u.Gpa
        | Error err -> printfn "Error: %s" err

        let! bobCheck = Database.getByIdAsync 2
        match bobCheck with
        | Some b -> printfn "Verified: %s" (Operations.format b)
        | None -> ()

        printfn "\n=== 6. Deleting Student 3 (Delete) ==="
        let! deleteRes = Database.deleteAsync 3
        match deleteRes with
        | Ok id -> printfn "Deleted student ID: %d" id
        | Error err -> printfn "Error: %s" err

        printfn "\n=== 7. Final Records in Database ==="
        let! final = Database.getAllAsync ()
        Operations.report final |> List.iter (printfn "%s")

        return 0
    }
    |> Async.RunSynchronously
