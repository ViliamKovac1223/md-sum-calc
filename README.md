# MD Sum Calc
This is a very small personal project that automates calculating sums in a
specially structured markdown file. I use this to calculate how much money I
have spent on what, and I have all my spendings in human readable "database"
with comments. Could I use spreadsheet? Yes, but I find this more enjoyable
than using spreadsheets.

I also used this project as an opportunity to learn go programming language.

# Structure Of MD File For This Project
```md
# First Category
<!-- hello this is comment -->
<!-- - date; sum; comment to the purchase -->
- 09.11.2024; 10.0; Hello this is comment
- 10.11.2024; 1.0;  Hello this is comment
- sum: 11.00

# Second Category
- 09.11.2024; 12.0; Hello this is comment
- 10.11.2024; -10.0; Hello this is comment
- sum: 2.00

<!-- Total sum of all categories -->
- total sum: 13.00
```

# How To Run This Project
Running this project from the code base folder.
```
go run . -f /path/to/file.md
```

You can also build the project and simply add the binary to the system path and
use this utility from anywhere.
```
go build
cp mdsumcalc /path/that/is/in/syste/path/
mdsumcalc -f /path/to/file.md
```
