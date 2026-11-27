// Lab 14 - Experiment 29: Collections and Functional Transformations

data class Student(
    val id: Int,
    val name: String,
    val department: String,
    val percentage: Double
)

fun main() {
    println("=================================================================")
    println(" Experiment 29: Collections and Functional Transformations")
    println("=================================================================\n")

    // -------------------------------------------------------------------------
    // Part 1: Read-Only and Mutable Variants of List, Set, and Map
    // -------------------------------------------------------------------------
    println("--- 1. Demonstrating Read-Only and Mutable Collections ---")

    // A. List
    println("\n[List Demonstration]")
    val readOnlyList: List<String> = listOf("Kotlin", "Java", "Python")
    println("Read-only List: $readOnlyList")
    println("Get element at index 1: ${readOnlyList.get(1)} (or ${readOnlyList[1]})")
    println("Contains 'Kotlin': ${readOnlyList.contains("Kotlin")}")

    val mutableList: MutableList<String> = mutableListOf("Kotlin", "Java")
    println("\nInitial Mutable List: $mutableList")
    mutableList.add("C++")
    println("After add('C++'): $mutableList")
    mutableList.remove("Java")
    println("After remove('Java'): $mutableList")
    println("Contains 'Java': ${mutableList.contains("Java")}")
    println("Get element at index 0: ${mutableList[0]}")

    // B. Set
    println("\n[Set Demonstration]")
    val readOnlySet: Set<Int> = setOf(10, 20, 30, 20)
    println("Read-only Set (duplicates discarded): $readOnlySet")
    println("Contains 20: ${readOnlySet.contains(20)}")

    val mutableSet: MutableSet<Int> = mutableSetOf(10, 20, 30)
    println("\nInitial Mutable Set: $mutableSet")
    mutableSet.add(40)
    println("After add(40): $mutableSet")
    mutableSet.remove(20)
    println("After remove(20): $mutableSet")
    println("Contains 20: ${mutableSet.contains(20)}")

    // C. Map
    println("\n[Map Demonstration]")
    val readOnlyMap: Map<String, Int> = mapOf("Alice" to 85, "Bob" to 92)
    println("Read-only Map: $readOnlyMap")
    println("Get 'Alice': ${readOnlyMap.get("Alice")}")
    println("Contains key 'Bob': ${readOnlyMap.containsKey("Bob")}")

    val mutableMap: MutableMap<String, Int> = mutableMapOf("Alice" to 85, "Bob" to 92)
    println("\nInitial Mutable Map: $mutableMap")
    mutableMap.put("Charlie", 78)
    println("After add/put 'Charlie' -> 78: $mutableMap")
    mutableMap.remove("Alice")
    println("After remove('Alice'): $mutableMap")
    println("Get 'Charlie': ${mutableMap["Charlie"]}")
    println("Contains key 'Alice': ${mutableMap.containsKey("Alice")}")

    // -------------------------------------------------------------------------
    // Part 2: Functional Transformations (map, filter, reduce)
    // -------------------------------------------------------------------------
    println("\n-----------------------------------------------------------------")
    println("--- 2. Functional Transformations on Integer List ---")
    val numbers = listOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    println("Original list: $numbers")

    // Filter: pick even numbers
    val evens = numbers.filter { it % 2 == 0 }
    println("After filter (even numbers): $evens")

    // Map: square each number
    val squared = evens.map { it * it }
    println("After map (square each even): $squared")

    // Reduce: calculate total sum
    val totalSum = squared.reduce { acc, num -> acc + num }
    println("After reduce (sum of squares): $totalSum")

    // -------------------------------------------------------------------------
    // Part 3: Object Sorting and Grouping
    // -------------------------------------------------------------------------
    println("\n-----------------------------------------------------------------")
    println("--- 3. Sorting Objects and Grouping by Property ---")
    val students = listOf(
        Student(101, "Alice", "Cybersecurity", 88.5),
        Student(102, "Bob", "Computer Science", 74.0),
        Student(103, "Charlie", "Cybersecurity", 95.0),
        Student(104, "Diana", "Data Science", 91.5),
        Student(105, "Evan", "Computer Science", 82.0)
    )

    println("All Students:")
    students.forEach { println("  $it") }

    // Sort by percentage descending
    val sortedStudents = students.sortedByDescending { it.percentage }
    println("\nStudents sorted by Percentage (Descending):")
    sortedStudents.forEach { println("  ${it.name} - ${it.percentage}% (${it.department})") }

    // Group by department
    val groupedByDept = students.groupBy { it.department }
    println("\nStudents grouped by Department (groupBy):")
    groupedByDept.forEach { (dept, deptStudents) ->
        println("  Department: $dept")
        deptStudents.forEach { s ->
            println("    - ${s.name} (${s.percentage}%)")
        }
    }
}
