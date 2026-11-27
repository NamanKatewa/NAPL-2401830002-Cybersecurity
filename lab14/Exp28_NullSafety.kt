// Lab 14 - Experiment 28: Null Safety (Nullable Types, Safe Calls, Elvis, and Force-Unwrap)

fun main() {
    println("=================================================================")
    println(" Experiment 28: Kotlin Null Safety Demonstrations")
    println("=================================================================\n")

    // -------------------------------------------------------------------------
    // Scenario 1: Nullable vs Non-Nullable Variable Types
    // -------------------------------------------------------------------------
    println("--- Scenario 1: Nullable vs Non-Nullable Types ---")

    val nonNullableText: String = "Hello Kotlin"
    println("Non-nullable variable: \"$nonNullableText\"")
    // nonNullableText = null  // COMPILE ERROR: Null can not be a value of a non-null type String

    val nullableText: String? = null
    println("Nullable variable initialized with null: $nullableText")

    val assignedNullable: String? = "Present text"
    println("Nullable variable initialized with value: \"$assignedNullable\"")
    println("Note: Assigning null to a non-nullable type (e.g., 'val s: String = null')")
    println("fails at compile-time with: 'Null can not be a value of a non-null type String'.\n")

    // -------------------------------------------------------------------------
    // Scenario 2: Safe Call Operator (?.)
    // -------------------------------------------------------------------------
    println("--- Scenario 2: Safe Call Operator (?.) ---")

    val wordWithContent: String? = "Cybersecurity"
    val wordAsNull: String? = null

    // Access property using safe call
    val len1 = wordWithContent?.length
    val len2 = wordAsNull?.length
    println("Length of \"$wordWithContent\" using (?.) : $len1")
    println("Length of null string using (?.)        : $len2 (safely evaluated to null)")

    // Method call using safe call
    val upper1 = wordWithContent?.uppercase()
    val upper2 = wordAsNull?.uppercase()
    println("Uppercase of \"$wordWithContent\" using (?.) : $upper1")
    println("Uppercase of null string using (?.)        : $upper2 (safely evaluated to null)\n")

    // -------------------------------------------------------------------------
    // Scenario 3: Elvis Operator (?:)
    // -------------------------------------------------------------------------
    println("--- Scenario 3: Elvis Operator (?:) ---")

    val username: String? = null
    val displayName = username ?: "Guest User"
    println("Input was null, default applied using (?:) : $displayName")

    val realUsername: String? = "admin_naman"
    val verifiedName = realUsername ?: "Guest User"
    println("Input was non-null, actual value preserved : $verifiedName")

    val textLength = wordAsNull?.length ?: 0
    println("Safe length calculation with Elvis fallback : $textLength\n")

    // -------------------------------------------------------------------------
    // Scenario 4: Not-Null Assertion Operator (!!) with try-catch
    // -------------------------------------------------------------------------
    println("--- Scenario 4: Force-Unwrap (!!) and NullPointerException Handling ---")

    val dangerousReference: String? = null

    try {
        println("Attempting to force-unwrap null reference with (!!)...")
        val forcedValue: String = dangerousReference!!
        println("Result: $forcedValue")
    } catch (e: NullPointerException) {
        println("Caught expected NullPointerException: ${e.javaClass.name}")
        println("Exception message: ${e.message}")
    }

    // Showing successful unwrap when value is present
    val safeReference: String? = "Valid Data"
    val unwrapped: String = safeReference!!
    println("Force-unwrap on valid non-null reference succeeds: \"$unwrapped\"")
}
