# word-frequency

Task: Calculate word frequency in large text file.
Also, I'm using this task to learn golang.

Implementation objectives:

- ~~Basic implementation~~
- ~~Add basic diagnostics to capture execution time~~
- More fine grain control on splitting words
- Create custom data structure (max-heap)
- Create test harness to 
    - verify implementation(s)
    - Measure implementation speeds 

Personal objectives: 

- Learn golang syntax and standard library
- Learn golang unit testing framework

--------

## Approaches

### 1. Bufio.Scanner with bufio.ScanWords

 - In-built package from golang and easy to use.
 - Splits words by splitting text on empty space delimiter.
 - Excludes empty spaces.
 - Not the best at splitting words correctly. Includes punctuation with the word.
 
 Notes about golang:
 
 - No ternary operator.
 - Formatting is non-negotiable.
 - Documentation on standard lib is decent.
 - Defer adds functions to execute into a stack, and hence they are run LIFO once current function has finished executing.
 
### 2. Bufio.Scanner with custom split function 

### 3. Bufio.Scanner with custom split function + Max heap

### 4. Bufio.Scanner with custom split function + Max heap + Channels
