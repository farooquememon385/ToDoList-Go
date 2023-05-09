# Simple To Do List project in Go using file handling
## Author : M Farooq

### to run the project run the follwoing command
`go run project.go`

### expected output

<hr>
******** Welcome to To-Do List ********<br>

To-do list loaded from file.<br>
Current To-Do List:<br>
<br>
--------
0. Project
--------
<br>
Enter a command (add/remove/show/quit): add
<br>
Enter the to-do item: assignment
<br>
Added 'assignment' to the to-do list.

--------
0. Project
1. assignment
--------
<br>
Enter a command (add/remove/show/quit): remove
<br>
Enter the index of the item to remove: 1
<br>
Removed 'assignment' from the to-do list.
<br>
--------
0. Project
--------
<br>
Enter a command (add/remove/show/quit): quit
<br>
To-do list saved to file.<br>
To-do list saved. Exiting...